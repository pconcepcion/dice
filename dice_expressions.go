// Package rpg provides tools to develop rpg games
package rpg

import (
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const (
	emptyExpression = iota
	wholeExpression
	numDicesPosition
	numSidesPosition
	modifierExpression
)

// SimpleDiceExpression represents a dice expression with just one type of dice
// dice expresions are based on the ones in RPtools ( http://lmwcs.com/rptools/wiki/Dice_Expressions )
type SimpleDiceExpression struct {
	expressionText string // Text that represents the dice expression
	numDices       int    // number of dices
	sides          int    // dice sides
	modifier       int    // modifier to the result
	drop           int    // number of lower results to drop
	keep           int    // number of lower results to keep
	explodeResult  int    // explode the result if greater than X
	target         int    // target to consider a success the value must be greater than X
	open           int    // count as success if greater than X
	reroll         int    // number of high results to reroll
	constant       int    // constant value
}

type Roller interface {
	Roll() DiceExpressionResult
}

var baseRegEx *regexp.Regexp
var constantRegexp *regexp.Regexp
var modifiersRegEx *regexp.Regexp

// Compiles the regular expressions
func init() {
	baseRegEx = regexp.MustCompile(`(?P<numDices>\d*)d(?P<numSides>\d+)(?P<modifiers>.*)`)
	constantRegexp = regexp.MustCompile(`^\d+$`)
	modifiersRegEx = regexp.MustCompile(`(?P<modifier>[dekors])(?P<value>\d*)(?P<addSubstract>[+-]*)(?P<constant>\d*)`)
}

// handleTokenMoffier handles the Modifier possible extra number
func (sde *SimpleDiceExpression) handleTokenModifier(tok, nextToken Token) {
	switch nextToken.typ {
	case tokenNumber:
		switch tok.val {
		case "k":
			sde.keep, _ = strconv.Atoi(nextToken.val)
		case "s":
			sde.target, _ = strconv.Atoi(nextToken.val)
		case "es":
			sde.explodeResult = sde.sides
			sde.target, _ = strconv.Atoi(nextToken.val)
		case "r":
			sde.reroll, _ = strconv.Atoi(nextToken.val)
		default:
			panic("Unexpected modifier")
		}
	case tokenEOF:
		switch tok.val {
		case "e":
			sde.explodeResult = sde.sides
		case "o":
			sde.open = sde.sides
		}
	}
}

// handlelTokenNumber handles the second or third tokenNumber
func (sde *SimpleDiceExpression) handleTokenNumber(tok, nextToken Token) {
	switch nextToken.typ {
	case tokenDice:
		panic("Unexpected diceToken")
	case tokenModifier:
		sde.sides, _ = strconv.Atoi(tok.val)
	case tokenEOF:
		if sde.sides == 0 {
			sde.sides, _ = strconv.Atoi(tok.val)
		}
		// if not the caller would know the modifier and assing to the propper place the value
	}
}

// handleInitialTokenNumber handles the first token when it's a number
func (sde *SimpleDiceExpression) handleInitialTokenNumber(tok, nextToken Token) {
	switch nextToken.typ {
	case tokenEOF:
		sde.constant, _ = strconv.Atoi(tok.val)
	case tokenDice:
		sde.numDices, _ = strconv.Atoi(tok.val)
	}
}

/**
 * Parse a simple dice expresion and save the relevant information on the struct
 */
func (sde *SimpleDiceExpression) parse() error {
	sde.expressionText = strings.TrimSpace(sde.expressionText)
	_, tokensChannel := lex(sde.expressionText)
	for tok := range tokensChannel {
		switch tok.typ {
		case tokenError:
			return errors.New(tok.val)
		case tokenNumber:
			nextToken := <-tokensChannel
			/// If it's the first number numDices must be 0
			if sde.numDices == 0 {
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
	}
	return nil
}

//Roll the expression and return the reslut or an error
func (sde *SimpleDiceExpression) Roll() (DiceExpressionResult, error) {
	if err := sde.parse(); err != nil {
		return nil, err
	}
	result := &simpleDiceExpressionResult{diceExpression: *sde, diceResults: make([]int, sde.numDices)}
	d := NewDice(sde.sides)
	for i := 0; i < sde.numDices; i++ {
		result.diceResults[i] = d.Roll()
	}
	fmt.Println("result.diceExpression: ", result.diceExpression)
	fmt.Println("result.diceResults: ", result.diceResults)
	sort.Sort(result.diceResults)
	fmt.Println("sorted result.diceResults: ", result.diceResults)
	if sde.keep > 0 {
		result.diceResults = result.diceResults[:sde.keep]
	}
	fmt.Println("kept result.diceResults: ", result.diceResults)
	result.total = sde.constant
	for j := 0; j < len(result.diceResults); j++ {
		result.total = result.diceResults[j]
	}
	fmt.Println("total: ", result.total)

	return result, nil
}
