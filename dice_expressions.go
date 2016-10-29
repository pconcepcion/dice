// Package dice provides tools to develop rpg games
package dice

import (
	"github.com/Sirupsen/logrus"
	"github.com/pkg/errors"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"sort"
	"strconv"
	"strings"
)

type diceModifier int

const (
	emptyExpression = iota
	wholeExpression
	numDicesPosition
	numSidesPosition
	modifierExpression
)

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

var log = logrus.New()

func init() {
	log.Formatter = new(prefixed.TextFormatter)
	log.Level = logrus.DebugLevel

}

// SimpleDiceExpression represents a dice expression with just one type of dice
// dice expresions are based on the ones in RPtools ( http://lmwcs.com/rptools/wiki/Dice_Expressions )
type SimpleDiceExpression struct {
	expressionText string       // Text that represents the dice expression
	numDices       int          // number of dices
	sides          int          // dice sides
	modifier       diceModifier // modifier to the result
	modifierValue  int          // value related to the modifier
	constant       int          // constant value
}

// Roller interface represents anthiing that can be "rolled" and generate a DiceExpressionResult
type Roller interface {
	Roll() DiceExpressionResult
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

// NewSimpleDiceExpression creates a new SimpleDiceExpresion initialized expressionText received but
// doesn't parse the expression SimpleDiceExpression.parse() should be called to parse the expression
func NewSimpleDiceExpression(expression string) SimpleDiceExpression {
	return SimpleDiceExpression{expressionText: expression}
}

// handleNextTokenNumber handles the state when the next token is a tokenNumber
func (sde *SimpleDiceExpression) handleNextTokenNumber(tok, nextToken Token) {
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
func (sde *SimpleDiceExpression) handleExtraTokenModifier(nextToken Token, modifier diceModifier) {
	sde.modifierValue = extractTokenValue(nextToken)
	sde.modifier = modifier
}

// handleNextTokenEOF handles the state when the next token is a tokenEOF
func (sde *SimpleDiceExpression) handleNextTokenEOF(tok, nextToken Token) {
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
func (sde *SimpleDiceExpression) handleTokenModifier(tok, nextToken Token) {
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
func (sde *SimpleDiceExpression) handleTokenNumber(tok, nextToken Token) {
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
func (sde *SimpleDiceExpression) handleInitialTokenNumber(tok, nextToken Token) {
	switch nextToken.typ {
	case tokenEOF:
		sde.constant = extractTokenValue(tok)
	case tokenDice:
		sde.numDices = extractTokenValue(tok)
	}
}

/**
 * Parse a simple dice expresion and save the relevant information on the struct
 */
func (sde *SimpleDiceExpression) parse() error {
	firstToken := true
	sde.expressionText = strings.TrimSpace(sde.expressionText)
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

//Roll the expression and return the reslut or an error
func (sde *SimpleDiceExpression) Roll() (DiceExpressionResult, error) {
	if err := sde.parse(); err != nil {
		return nil, errors.Wrap(err, "Parsing error")
	}
	if sde.numDices == 0 || sde.sides == 0 {
		return &simpleDiceExpressionResult{diceExpression: *sde, total: 0}, nil
	}

	result := &simpleDiceExpressionResult{diceExpression: *sde, diceResults: make([]int, sde.numDices)}
	d := NewDice(sde.sides)
	for i := 0; i < sde.numDices; i++ {
		result.diceResults[i] = d.Roll()
	}
	log.WithFields(logrus.Fields{"result.diceExpresion": result.diceExpression}).Debug("Dice Expression")
	log.WithFields(logrus.Fields{"result.diceResults": result.diceResults}).Info("Dices rolled")
	sort.Sort(sort.Reverse(result.diceResults))
	log.WithFields(logrus.Fields{"result.diceResults": result.diceResults}).Debug("Sorted")
	result.handleModifier(sde)
	result.total += sde.constant
	log.Infoln("total: ", result.total)

	return result, nil
}
