// Package rpg provides tools to develop rpg games
package rpg

import (
	"errors"
	"regexp"
	"strconv"
)

// SimpleDiceExpression represents a dice expression with just one type of dice
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
}

func (sde SimpleDiceExpression) parse() ([]string, error) {
	re := regexp.MustCompile("(\\d*)d(\\d)(.*)")
	dices := re.FindStringSubmatch(sde.expressionText)
	if dices == nil {
		return nil, errors.New("Invalid dice expression")
	}
	numDices, err := strconv.Atoi(dices[1])
	if err != nil {
		return nil, err
	}
	sde.numDices = numDices
	sides, err := strconv.Atoi(dices[2])
	if err != nil {
		return nil, err
	}
	sde.sides = sides
	return dices, nil
}
