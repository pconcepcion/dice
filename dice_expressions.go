// Package rpg provides tools to develop rpg games
package rpg

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	wholeExpression = iota
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
	target         int    // count as success if greater than X
	rerollL        int    // number of lower dices to reroll
	rerollH        int    // number of high results to reroll
	constant       int    // constant value
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

/**
 * Parse a simple dice expresion and save the relevant information on the struct
 */
func (sde *SimpleDiceExpression) parse() ([]string, error) {
	var err error
	sde.expressionText = strings.TrimSpace(sde.expressionText)
	// Check if it's a constant dice expression
	if constantRegexp.MatchString(sde.expressionText) == true {
		sde.constant, err = strconv.Atoi(sde.expressionText)
		if err != nil {
			return nil, err
		}
		return []string{sde.expressionText, "", ""}, nil
	}
	dices := baseRegEx.FindStringSubmatch(strings.TrimSpace(sde.expressionText))
	if dices == nil {
		return nil, errors.New("Invalid dice expression")
	}
	// If it's not specified assume it's one dice
	sde.numDices = 1
	if dices[numDicesPosition] != "" {
		sde.numDices, err = strconv.Atoi(dices[numDicesPosition])
		if err != nil {
			return nil, err
		}
	}
	sde.sides, err = strconv.Atoi(dices[numSidesPosition])
	if err != nil {
		return nil, err
	}
	if dices[modifierExpression] != "" {
		modifiers, err := sde.parseModifiers(dices[modifierExpression])
		if err != nil {
			return nil, err
		}
		dices = append(dices[:len(dices)-1], modifiers[1:]...)
	}
	return dices, nil
}

/**
 * parse the modifiers part of the exprssion and store the values on the expression structure
 */
func (sde *SimpleDiceExpression) parseModifiers(modifierString string) ([]string, error) {
	var err error
	modifiers := modifiersRegEx.FindStringSubmatch(modifierString)
	fmt.Printf("modifiers %s %#v\n", modifierString, modifiers)
	if modifiers == nil {
		return nil, errors.New("Invalid dice expression")
	}
	switch modifierString[0] {
	case 'k': // Keep
		if modifiers[2] == "" {
			return nil, errors.New("Invalid dice expression, missing number of dices to keep")
		}
		sde.keep, err = strconv.Atoi(modifiers[2])
		if err != nil {
			return nil, err
		}
	case 'r': // Rerroll
		if modifiers[2] == "" {
			return nil, errors.New("Invalid dice expression, missing number of dices to reroll")
		}
		sde.rerollL, err = strconv.Atoi(modifiers[2])
		if err != nil {
			return nil, err
		}
	case 's': // Keep
		if modifiers[2] == "" {
			return nil, errors.New("Invalid dice expression, missing success number")
		}
		sde.target, err = strconv.Atoi(modifiers[2])
		if err != nil {
			return nil, err
		}
	// case 'd': // Discard

	default:
		return nil, errors.New("Invalid dice modifier")
	}
	return modifiers, nil
}
