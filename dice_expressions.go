// Package rpg provides tools to develop rpg games
package rpg

import (
	"errors"
	"fmt"
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
	exlpodingSuccess
	explode
	open
	drop
)

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

// TODO:handle strconv.Atoi errors

// handleTokenMoffier handles the Modifier possible extra number
func (sde *SimpleDiceExpression) handleTokenModifier(tok, nextToken Token) {
	switch nextToken.typ {
	case tokenNumber:
		switch tok.val {
		case "k":
			sde.modifierValue, _ = strconv.Atoi(nextToken.val)
			sde.modifier = keep
		case "kl":
			sde.modifierValue, _ = strconv.Atoi(nextToken.val)
			sde.modifier = keepLower
		case "e":
			sde.modifierValue, _ = strconv.Atoi(nextToken.val)
			sde.modifier = explode
		case "s":
			sde.modifierValue, _ = strconv.Atoi(nextToken.val)
			sde.modifier = success
		case "es":
			sde.modifier = exlpodingSuccess
			sde.modifierValue, _ = strconv.Atoi(nextToken.val)
		case "r":
			sde.modifier = reroll
			sde.modifierValue, _ = strconv.Atoi(nextToken.val)
		default:
			panic("Unexpected modifier")
		}
	case tokenEOF:
		switch tok.val {
		case "e":
			sde.modifier = explode
			sde.modifierValue = sde.sides
		case "o":
			sde.modifier = open
			sde.modifierValue = sde.sides
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
		return nil, err
	}
	if sde.numDices == 0 || sde.sides == 0 {
		return &simpleDiceExpressionResult{diceExpression: *sde, total: 0}, nil
	}

	result := &simpleDiceExpressionResult{diceExpression: *sde, diceResults: make([]int, sde.numDices)}
	d := NewDice(sde.sides)
	for i := 0; i < sde.numDices; i++ {
		result.diceResults[i] = d.Roll()
	}
	fmt.Printf("result.diceExpression: %+v", result.diceExpression)
	fmt.Println("result.diceResults: ", result.diceResults)
	sort.Sort(sort.Reverse(result.diceResults))
	fmt.Println("sorted result.diceResults: ", result.diceResults)
	switch sde.modifier {
	case keep:
		result.diceResults = result.diceResults[:sde.modifierValue]
		fmt.Println("kept result.diceResults: ", result.diceResults)
		result.SumTotal()
	case keepLower:
		// TODO: solve this wihout so much sorting...
		sort.Sort(result.diceResults)
		result.diceResults = result.diceResults[:sde.modifierValue]
		sort.Sort(sort.Reverse(result.diceResults))
		fmt.Println("keptLower result.diceResults: ", result.diceResults)
		result.SumTotal()
	case success:
		result.Success(sde.modifierValue)
	case exlpodingSuccess:
		result.ExplodingSuccess(sde.modifierValue)
		fmt.Println("explode result.diceResults: ", result.diceResults)
		fmt.Println("explode result.extrDiceResults: ", result.extraDiceResults)
	case explode:
		result.Explode()
		fmt.Println("explode result.diceResults: ", result.diceResults)
		fmt.Println("explode result.extrDiceResults: ", result.extraDiceResults)
		result.SumTotal()
	case open:
		result.Open()
		fmt.Println("explode result.diceResults: ", result.diceResults)
		fmt.Println("explode result.extrDiceResults: ", result.extraDiceResults)
		sort.Sort(sort.Reverse(result.diceResults))
		result.total += result.diceResults[0]
	case reroll:
		result.Reroll(sde.modifierValue)
		sort.Sort(sort.Reverse(result.diceResults))
		fmt.Println("reroll result.diceResults: ", result.diceResults)
	case drop:
		result.diceResults = result.diceResults[:(sde.numDices - sde.modifierValue)]
		fmt.Println("drop result.diceResults: ", result.diceResults)
		result.SumTotal()
	default:
		result.SumTotal()
	}
	result.total += sde.constant
	fmt.Println("total: ", result.total)

	return result, nil
}
