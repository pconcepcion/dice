// Package rpg provides tools to develop rpg games
package rpg

import (
	"fmt"
)

// DiceResults represent the values obtained in the dices roled
type DiceResults []int

// Implement the Sort interfaace

// Len gives the length of the Dice results
func (a DiceResults) Len() int { return len(a) }

// Swap swaps the positions of two elements from the Dice Results
func (a DiceResults) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// Less returns true if the ith element is less than the jth
func (a DiceResults) Less(i, j int) bool { return a[i] < a[j] }

// DiceExpressionResult is an interfaace that represents the result of rolling some kind of dice expression
type DiceExpressionResult interface {
	fmt.Stringer
	GetResults() DiceResults
	GetTotal() int
}

type simpleDiceExpressionResult struct {
	diceExpression SimpleDiceExpression
	diceResults    DiceResults
	total          int
	verbose        bool
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
