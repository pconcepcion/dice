// Package rpg provides tools to develop rpg games
package rpg

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"sort"
)

// ExplodingMaxDices is the maximum number of explosions of a dice
const ExplodingMaxDices = 100

// DiceResults is an array of ints to hold the dice results
type DiceResults []int

// Implement the Sort interfaace

// Len gives the length of the Dice results
func (dr DiceResults) Len() int { return len(dr) }

// Swap swaps the positions of two elements from the Dice Results
func (dr DiceResults) Swap(i, j int) { dr[i], dr[j] = dr[j], dr[i] }

// Less returns true if the ith element is less than the jth
func (dr DiceResults) Less(i, j int) bool { return dr[i] < dr[j] }

// Sum sums the values in a DiceResults slice
func (dr DiceResults) Sum() int {
	sum := 0
	for i := 0; i < len(dr); i++ {
		sum += dr[i]
	}
	return sum
}

// DiceExpressionResult is an interfaace that represents the result of rolling some kind of dice expression
type DiceExpressionResult interface {
	fmt.Stringer
	GetResults() DiceResults
	GetTotal() int
}

type simpleDiceExpressionResult struct {
	diceExpression   SimpleDiceExpression
	diceResults      DiceResults
	extraDiceResults DiceResults
	total            int
	verbose          bool
}

// handleModifier does all the magic and applies the modifier corresponding to the received
// SimpleDiceExpression to the result
func (sder *simpleDiceExpressionResult) handleModifier(sde *SimpleDiceExpression) {
	switch sde.modifier {
	case keep:
		sder.diceResults = sder.diceResults[:sde.modifierValue]
		log.WithFields(logrus.Fields{"sder.diceResults": sder.diceResults}).Debug("Keep")
		sder.SumTotal()
	case keepLower:
		// TODO: solve this wihout so much sorting...
		sort.Sort(sder.diceResults)
		sder.diceResults = sder.diceResults[:sde.modifierValue]
		sort.Sort(sort.Reverse(sder.diceResults))
		log.WithFields(logrus.Fields{"sder.diceResults": sder.diceResults}).Debug("Keep Lower")
		sder.SumTotal()
	case success:
		sder.Success(sde.modifierValue)
	case explodingSuccess:
		sder.ExplodingSuccess(sde.modifierValue)
		log.WithFields(logrus.Fields{"sder.diceResults": sder.diceResults,
			"sder.extrDiceResults": sder.extraDiceResults}).Debug("Exploding Success")
	case explode:
		sder.Explode()
		log.WithFields(logrus.Fields{"sder.diceResults": sder.diceResults,
			"sder.extrDiceResults": sder.extraDiceResults}).Debug("Explode")
		sder.SumTotal()
	case open:
		sder.Open()
		log.WithFields(logrus.Fields{"sder.diceResults": sder.diceResults,
			"sder.extrDiceResults": sder.extraDiceResults}).Debug("Open")
		sort.Sort(sort.Reverse(sder.diceResults))
		sder.total += sder.diceResults[0]
	case reroll:
		sder.Reroll(sde.modifierValue)
		sort.Sort(sort.Reverse(sder.diceResults))
		log.WithFields(logrus.Fields{"sder.diceResults": sder.diceResults}).Debug("Reroll")
	case drop:
		sder.diceResults = sder.diceResults[:(sde.numDices - sde.modifierValue)]
		log.WithFields(logrus.Fields{"sder.diceResults": sder.diceResults}).Debug("Drop")
		sder.SumTotal()
	default:
		sder.SumTotal()
	}
}

// Success counts the number of results with a value mayor or equal to the target value and stores the result in the total
func (sder *simpleDiceExpressionResult) Success(targetValue int) {
	sder.total = 0
	for _, dr := range sder.diceResults {
		if dr >= targetValue {
			sder.total++
		}
	}
}

// SumTotal sums the values on the diceResults into the total field
func (sder *simpleDiceExpressionResult) SumTotal() {
	sder.total = sder.diceResults.Sum()
}

// TODO: Adde example to the documentation
// explodeDice explodes the result of one dice rolling and adding a dice to the results while the result is the same than the number of sides
func (sder *simpleDiceExpressionResult) explodeDice() DiceResults {
	numSides := sder.diceExpression.sides
	threshold := sder.diceExpression.modifierValue
	d := NewDice(numSides)
	if threshold <= 1 {
		log.Debugf("Threshold <= 1 appending  %d results\n", ExplodingMaxDices)
		results := make([]int, ExplodingMaxDices)
		newValue := d.Roll()
		results = append(results, newValue)
		for i := 0; i < ExplodingMaxDices; i++ {
			results[i] = d.Roll()
		}
		return results
	}
	results := make([]int, 0, 3)
	newValue := d.Roll()
	results = append(results, newValue)
	log.Debugln("append 1st new result: ", newValue)
	for newValue >= threshold {
		newValue = d.Roll()
		results = append(results, newValue)
		log.Debugln("append new result: ", newValue)
	}
	return results
}

// Open rolls all the dices and explodes them and sets the total as the maximum of the results
// the new roll is still equal to the nubmer of sides
func (sder *simpleDiceExpressionResult) Open() {
	numSides := sder.diceExpression.sides
	for i, res := range sder.diceResults {
		if res == numSides {
			results := sder.explodeDice()
			sder.extraDiceResults = append(sder.extraDiceResults, results...)
			sder.diceResults[i] = results.Sum() + res
		}
	}
}

// Explode rolls one extra dice for each reult that it's equal to the number of sides, and keeps doing it if the result of
// the new roll is still equal to the nubmer of sides
func (sder *simpleDiceExpressionResult) Explode() {
	for _, res := range sder.diceResults {
		if res >= sder.diceExpression.modifierValue {
			results := sder.explodeDice()
			sder.extraDiceResults = append(sder.extraDiceResults, results...)
		}
	}
	sder.diceResults = append(sder.diceResults, sder.extraDiceResults...)
}

// Success but with exploding Dices
func (sder *simpleDiceExpressionResult) ExplodingSuccess(targetValue int) {
	sder.total = 0
	numSides := sder.diceExpression.sides
	for i, res := range sder.diceResults {
		if res == numSides {
			results := sder.explodeDice()
			sder.extraDiceResults = append(sder.extraDiceResults, results...)
			sder.diceResults[i] = results.Sum() + res
		}
		if sder.diceResults[i] >= targetValue {
			sder.total++
		}
	}
}

// Reroll rerrols the dices with a result smaller than minValue until obtaining something biger or equal than minValue
func (sder *simpleDiceExpressionResult) Reroll(minValue int) {
	numSides := sder.diceExpression.sides
	d := NewDice(numSides)
	sder.total = 0
	// We can't reroll and get more dices than whe have thrown
	if minValue >= numSides {
		for i := range sder.diceResults {
			sder.diceResults[i] = minValue
			sder.total += minValue
		}
	} else {
		for i := range sder.diceResults {
			// while the value of the result is lower reroll the dice
			for sder.diceResults[i] < minValue {
				sder.diceResults[i] = d.Roll()
			}
			sder.total += sder.diceResults[i]
		}
	}

}

// String returns a string representing the simpleDiceExpressionResult, if verbose is true  it will print more info.
func (sder *simpleDiceExpressionResult) String() string {
	if sder.verbose {
		return fmt.Sprintf("%s : %v -> %d", sder.diceExpression.expressionText, sder.diceResults, sder.total)
	}
	return fmt.Sprintf("%d", sder.total)
}

// GetResults return the DiceResults
func (sder *simpleDiceExpressionResult) GetResults() DiceResults {
	return sder.diceResults
}

// GetResults return the total
func (sder *simpleDiceExpressionResult) GetTotal() int {
	return sder.total
}
