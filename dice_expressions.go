// Package dice provides tools to develop rpg games
package dice

import (
	"sort"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type diceModifier int

const (
	emptyExpression = iota
	wholeExpression
	numDicesPosition
	numSidesPosition
	modifierExpression
)

// MaxExpressionLength represents the maximum length of the string representing a DiceExpression
const MaxExpressionLength = 64

//go:generate stringer -type=diceModifier
const (
	normal diceModifier = iota
	keep
	keepLower
	reroll
	success
	explodingSuccess
	explode
	open
	drop
)

var (
	// ErrExpressionTooLong is the error thrown when the string repressenting the expression is longer than
	// MaxExpressionLength
	ErrExpressionTooLong = errors.New("Expression Too Long")
)

// var log = logrus.New()

func init() {
	log.SetFormatter(&log.TextFormatter{DisableLevelTruncation: true, FullTimestamp: true, PadLevelText: true})
	log.SetLevel(log.DebugLevel)
	//log.SetLevel(log.WarnLevel)
}

// SimpleExpression represents a dice expression with just one type of dice
// dice expresions are based on the ones in RPtools ( http://lmwcs.com/rptools/wiki/Dice_Expressions )
type SimpleExpression struct {
	expressionText string       // Text that represents the dice expression
	numDices       int          // number of dices
	sides          int          // dice sides
	modifier       diceModifier // modifier to the result
	modifierValue  int          // value related to the modifier
	constant       int          // constant value
}

// Roller interface represents anthiing that can be "rolled" and generate a ExpressionResult
type Roller interface {
	Roll() ExpressionResult
}

// extractTokenValue extracts from the received Token the value and converts it to to an int
// it will panic on failure
func extractTokenValue(tok Token) int {
	intValue, err := strconv.Atoi(tok.val)
	if err != nil {
		log.Panicf("Unexpected token value, not an int, %v\n", tok)
		panic("Unexpected token value, not an int")
	}
	return intValue
}

// NewSimpleExpression creates a new SimpleExpression initialized expressionText received but
// doesn't parse the expression SimpleExpression.parse() should be called to parse the expression
func NewSimpleExpression(expression string) SimpleExpression {
	return SimpleExpression{expressionText: expression}
}

// NewParsedSimpleExpression creates a new SimpleExpression initialized expressionText received and
// parses the expression returning an error if the parse fails
func NewParsedSimpleExpression(expression string) (*SimpleExpression, error) {
	if len(expression) > MaxExpressionLength {
		return nil, ErrExpressionTooLong
	}
	sde := SimpleExpression{expressionText: expression}
	if err := sde.parse(); err != nil {
		return nil, errors.Wrap(err, "Parsing error")
	}
	return &sde, nil
}

// handleNextTokenNumber handles the state when the next token is a tokenNumber
func (sde *SimpleExpression) handleNextTokenNumber(tok, nextToken Token) {
	log.Debug("sde.modifier: ", sde.modifier, ", sde.modifierValue: ", sde.modifierValue,
		", tok: ", tok, ", nextToken: ", nextToken)
	switch tok.val {
	case "k":
		sde.handleExtraTokenModifier(nextToken, keep)
	case "kl":
		sde.handleExtraTokenModifier(nextToken, keepLower)
	case "e":
		sde.handleExtraTokenModifier(nextToken, explode)
	case "s":
		sde.handleExtraTokenModifier(nextToken, success)
	case "es":
		sde.handleExtraTokenModifier(nextToken, explodingSuccess)
	case "r":
		sde.handleExtraTokenModifier(nextToken, reroll)
	default:
		log.Panicln("Unexpected modifier")
		panic("Unexpected modifier")
	}
}

// handleExtraTokenModifier handles a modifier that requires a numeric extra token, stores the
// modifier on sde and the extra token value on sde.modifierValue
func (sde *SimpleExpression) handleExtraTokenModifier(nextToken Token, modifier diceModifier) {
	log.Debug("sde.modifier: ", sde.modifier, ", sde.modifierValue: ", sde.modifierValue,
		", newModifier: ", modifier, ", nextToken: ", nextToken)
	sde.modifierValue = extractTokenValue(nextToken)
	sde.modifier = modifier
}

// handleNextTokenEOF handles the state when the next token is a tokenEOF
func (sde *SimpleExpression) handleNextTokenEOF(tok, nextToken Token) {
	log.Debug("sde.modifier: ", sde.modifier, ", sde.modifierValue: ", sde.modifierValue,
		", tok: ", tok, ", nextToken: ", nextToken)
	switch tok.val {
	case "e":
		sde.modifier = explode
		sde.modifierValue = sde.sides
	case "o":
		sde.modifier = open
		sde.modifierValue = sde.sides
	}
}

// handleTokenMoffier handles the Modifier optional extra number
func (sde *SimpleExpression) handleTokenModifier(tok, nextToken Token) {
	log.Debug("sde.modifier: ", sde.modifier, ", sde.modifierValue: ", sde.modifierValue,
		", tok: ", tok, ", nextToken: ", nextToken)
	switch nextToken.typ {
	case tokenNumber:
		sde.handleNextTokenNumber(tok, nextToken)
	case tokenEOF:
		sde.handleNextTokenEOF(tok, nextToken)
	default:
		log.Panicln("Unexpected nextToken")
		panic("Unexpected nextToken")
	}
}

// handlelTokenNumber handles the second or third tokenNumber
func (sde *SimpleExpression) handleTokenNumber(tok, nextToken Token) {
	log.Debug("sde.modifier: ", sde.modifier, ", sde.modifierValue: ", sde.modifierValue,
		", tok: ", tok, ", nextToken: ", nextToken)
	switch nextToken.typ {
	case tokenDice:
		log.Panicln("Unexpected modifier")
		panic("Unexpected diceToken")
	case tokenModifier:
		sde.sides = extractTokenValue(tok)
	case tokenEOF:
		if sde.sides == 0 {
			sde.sides = extractTokenValue(tok)
		}
		// if not the caller would know the modifier and assing to the propper place the value
	}
}

// handleInitialTokenNumber handles the first token when it's a number
func (sde *SimpleExpression) handleInitialTokenNumber(tok, nextToken Token) {
	log.Debug("sde.modifier: ", sde.modifier, ", sde.modifierValue: ", sde.modifierValue,
		", tok: ", tok, ", nextToken: ", nextToken)
	switch nextToken.typ {
	case tokenEOF:
		sde.constant = extractTokenValue(tok)
	case tokenDice:
		sde.numDices = extractTokenValue(tok)
	}
}

// parse a simple dice expresion and save the relevant information on the struct
func (sde *SimpleExpression) parse() error {
	firstToken := true
	sde.expressionText = strings.TrimSpace(sde.expressionText)
	if len(sde.expressionText) > MaxExpressionLength {
		return ErrExpressionTooLong
	}
	_, tokensChannel := lex(sde.expressionText)
	for tok := range tokensChannel {
		switch tok.typ {
		case tokenError:
			return errors.New(tok.val)
		case tokenNumber:
			nextToken := <-tokensChannel
			/// If it's the first
			if firstToken {
				sde.handleInitialTokenNumber(tok, nextToken)
			} else {
				sde.handleTokenNumber(tok, nextToken)
				if nextToken.typ == tokenModifier {
					sde.handleTokenModifier(nextToken, <-tokensChannel)
				}
			}
		case tokenDice:
			// Only found when then number was ommited so it's one dice
			sde.numDices = 1
		}
		firstToken = false
	}
	return nil
}

//Roll the expression and return the result or an error
func (sde *SimpleExpression) Roll() (ExpressionResult, error) {
	if err := sde.parse(); err != nil {
		return nil, errors.Wrap(err, "Parsing error")
	}
	result := sde.RollPreParsed()
	return result, nil
}

//RollPreParsed rolls an already parsed expression and return the result
func (sde *SimpleExpression) RollPreParsed() ExpressionResult {
	if sde.numDices == 0 || sde.sides == 0 {
		return &simpleExpressionResult{diceExpression: *sde, total: 0}
	}

	result := &simpleExpressionResult{diceExpression: *sde, Results: make([]int, sde.numDices)}
	d := NewDice(sde.sides)
	for i := 0; i < sde.numDices; i++ {
		result.Results[i] = d.Roll()
	}
	log.WithFields(log.Fields{"result.diceExpresion": result.diceExpression}).Debug("Dice Expression")
	log.WithFields(log.Fields{"result.Results": result.Results}).Info("Dices rolled")
	sort.Sort(sort.Reverse(result.Results))
	log.WithFields(log.Fields{"result.Results": result.Results}).Debug("Sorted")
	result.handleModifier(sde)
	result.total += sde.constant
	log.Infoln("total: ", result.total)

	return result
}
