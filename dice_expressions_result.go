// Package dice provides tools to develop rpg games
package dice

import (
	"fmt"
	"sort"

	log "github.com/sirupsen/logrus"
)

// ExplodingMaxDices is the maximum number of explosions of a dice
const ExplodingMaxDices = 100

// Results is an array of ints to hold the dice results
type Results []int

// Implement the Sort interfaace

// Len gives the length of the Dice results
func (dr Results) Len() int { return len(dr) }

// Swap swaps the positions of two elements from the Dice Results
func (dr Results) Swap(i, j int) { dr[i], dr[j] = dr[j], dr[i] }

// Less returns true if the ith element is less than the jth
func (dr Results) Less(i, j int) bool { return dr[i] < dr[j] }

// Sum sums the values in a Results slice
func (dr Results) Sum() int {
	sum := 0
	for i := 0; i < len(dr); i++ {
		sum += dr[i]
	}
	return sum
}

// ExpressionResult is an interfaace that represents the result of rolling some kind of dice expression
type ExpressionResult interface {
	fmt.Stringer
	GetResults() Results
	GetTotal() int
}

type simpleExpressionResult struct {
	diceExpression SimpleExpression
	Results        Results
	extraResults   Results
	total          int
	verbose        bool
}

// handleModifier does all the magic and applies the modifier corresponding to the received
// SimpleExpression to the result
func (sder *simpleExpressionResult) handleModifier(sde *SimpleExpression) {
	switch sde.modifier {
	case keep:
		// TODO: This may be extacted to a function
		if sde.modifierValue > len(sder.Results) {
			log.Warnln("Trying to keep more dices than rolled dices")
			sde.modifierValue = len(sder.Results)
		}
		sder.Results = sder.Results[:sde.modifierValue]
		log.WithFields(log.Fields{"sder.Results": sder.Results}).Debug("Keep")
		sder.SumTotal()
	case keepLower:
		if sde.modifierValue > len(sder.Results) {
			log.Warnln("Trying to keep more dices than rolled dices")
			sde.modifierValue = len(sder.Results)
		}
		// TODO: solve this wihout so much sorting...
		sort.Sort(sder.Results)
		sder.Results = sder.Results[:sde.modifierValue]
		sort.Sort(sort.Reverse(sder.Results))
		log.WithFields(log.Fields{"sder.Results": sder.Results}).Debug("Keep Lower")
		sder.SumTotal()
	case success:
		sder.Success(sde.modifierValue)
	case explodingSuccess:
		sder.ExplodingSuccess(sde.modifierValue)
		log.WithFields(log.Fields{"sder.Results": sder.Results,
			"sder.extrResults": sder.extraResults}).Debug("Exploding Success")
	case explode:
		sder.Explode()
		log.WithFields(log.Fields{"sder.Results": sder.Results,
			"sder.extrResults": sder.extraResults}).Debug("Explode")
		sder.SumTotal()
	case open:
		sder.Open()
		log.WithFields(log.Fields{"sder.Results": sder.Results,
			"sder.extrResults": sder.extraResults}).Debug("Open")
		sort.Sort(sort.Reverse(sder.Results))
		sder.total += sder.Results[0]
	case reroll:
		sder.Reroll(sde.modifierValue)
		sort.Sort(sort.Reverse(sder.Results))
		log.WithFields(log.Fields{"sder.Results": sder.Results}).Debug("Reroll")
	case drop:
		sder.Results = sder.Results[:(sde.numDices - sde.modifierValue)]
		log.WithFields(log.Fields{"sder.Results": sder.Results}).Debug("Drop")
		sder.SumTotal()
	default:
		sder.SumTotal()
	}
}

// Success counts the number of results with a value mayor or equal to the target value and stores the result in the total
func (sder *simpleExpressionResult) Success(targetValue int) {
	sder.total = 0
	for _, dr := range sder.Results {
		if dr >= targetValue {
			sder.total++
		}
	}
}

// SumTotal sums the values on the Results into the total field
func (sder *simpleExpressionResult) SumTotal() {
	sder.total = sder.Results.Sum()
}

// TODO: Adde example to the documentation
// explodeDice explodes the result of one dice rolling and adding a dice to the results while the result is the same than the number of sides
func (sder *simpleExpressionResult) explodeDice() Results {
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
func (sder *simpleExpressionResult) Open() {
	numSides := sder.diceExpression.sides
	for i, res := range sder.Results {
		if res == numSides {
			results := sder.explodeDice()
			sder.extraResults = append(sder.extraResults, results...)
			sder.Results[i] = results.Sum() + res
		}
	}
}

// Explode rolls one extra dice for each reult that it's equal to the number of sides, and keeps doing it if the result of
// the new roll is still equal to the nubmer of sides
func (sder *simpleExpressionResult) Explode() {
	for _, res := range sder.Results {
		if res >= sder.diceExpression.modifierValue {
			results := sder.explodeDice()
			sder.extraResults = append(sder.extraResults, results...)
		}
	}
	sder.Results = append(sder.Results, sder.extraResults...)
}

// Success but with exploding Dices
func (sder *simpleExpressionResult) ExplodingSuccess(targetValue int) {
	sder.total = 0
	numSides := sder.diceExpression.sides
	for i, res := range sder.Results {
		if res == numSides {
			results := sder.explodeDice()
			sder.extraResults = append(sder.extraResults, results...)
			sder.Results[i] = results.Sum() + res
		}
		if sder.Results[i] >= targetValue {
			sder.total++
		}
	}
}

// Reroll rerrols the dices with a result smaller than minValue until obtaining something biger or equal than minValue
func (sder *simpleExpressionResult) Reroll(minValue int) {
	numSides := sder.diceExpression.sides
	d := NewDice(numSides)
	sder.total = 0
	// We can't reroll and get more dices than whe have thrown
	if minValue >= numSides {
		for i := range sder.Results {
			sder.Results[i] = minValue
			sder.total += minValue
		}
	} else {
		for i := range sder.Results {
			// while the value of the result is lower reroll the dice
			for sder.Results[i] < minValue {
				sder.Results[i] = d.Roll()
			}
			sder.total += sder.Results[i]
		}
	}

}

// String returns a string representing the simpleExpressionResult, if verbose is true  it will print more info.
func (sder *simpleExpressionResult) String() string {
	if sder.verbose {
		return fmt.Sprintf("%s : %v -> %d", sder.diceExpression.expressionText, sder.Results, sder.total)
	}
	return fmt.Sprintf("%d", sder.total)
}

// GetResults return the Results
func (sder *simpleExpressionResult) GetResults() Results {
	return sder.Results
}

// GetResults return the total
func (sder *simpleExpressionResult) GetTotal() int {
	return sder.total
}
