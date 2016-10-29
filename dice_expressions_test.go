// Package dice provides tools to develop rpg games
package dice

import (
	"testing"
)

// Test that Parse splits the dices properly
func TestParse(t *testing.T) {
	var parseTestStrings = []struct {
		s   string
		out SimpleExpression
	}{
		// TODO: test with spaces and invalid dice expressions
		// Constants
		{"1", SimpleExpression{expressionText: "1", constant: 1}},
		{"10", SimpleExpression{expressionText: "10", constant: 10}},
		{"100", SimpleExpression{expressionText: "100", constant: 100}},
		{"1000", SimpleExpression{expressionText: "1000", constant: 1000}},
		{"4321", SimpleExpression{expressionText: "4321", constant: 4321}},
		{" 1000", SimpleExpression{expressionText: "1000", constant: 1000}},
		{"1000 ", SimpleExpression{expressionText: "1000", constant: 1000}},
		// Basic dices
		{"d2", SimpleExpression{numDices: 1, expressionText: "d2", sides: 2}},
		{"d4", SimpleExpression{numDices: 1, expressionText: "d4", sides: 4}},
		{"d6", SimpleExpression{numDices: 1, expressionText: "d6", sides: 6}},
		{"d8", SimpleExpression{numDices: 1, expressionText: "d8", sides: 8}},
		{"d10", SimpleExpression{numDices: 1, expressionText: "d10", sides: 10}},
		{"d12", SimpleExpression{numDices: 1, expressionText: "d12", sides: 12}},
		{"d20", SimpleExpression{numDices: 1, expressionText: "d20", sides: 20}},
		{"d100", SimpleExpression{numDices: 1, expressionText: "d100", sides: 100}},
		{"d200", SimpleExpression{numDices: 1, expressionText: "d200", sides: 200}},
		{"d1000", SimpleExpression{numDices: 1, expressionText: "d1000", sides: 1000}},
		// More complex expressions
		{"3d3", SimpleExpression{numDices: 3, expressionText: "3d3", sides: 3}},
		{"3d6", SimpleExpression{numDices: 3, expressionText: "3d6", sides: 6}},
		{"1d2", SimpleExpression{numDices: 1, expressionText: "1d2", sides: 2}},
		{"3d6k2", SimpleExpression{numDices: 3, expressionText: "3d6k2", sides: 6, modifier: keep, modifierValue: 2}},
		{"4d8r2", SimpleExpression{numDices: 4, expressionText: "4d8r2", sides: 8, modifier: reroll, modifierValue: 2}},
		{"4d8s6", SimpleExpression{numDices: 4, expressionText: "4d8s6", sides: 8, modifier: success, modifierValue: 6}},
		{"6d6e", SimpleExpression{numDices: 6, expressionText: "6d6e", sides: 6, modifier: explode, modifierValue: 6}},
		{"7d6es8", SimpleExpression{numDices: 7, expressionText: "7d6es8", sides: 6, modifier: explodingSuccess, modifierValue: 8}},
		{"8d6o", SimpleExpression{numDices: 8, expressionText: "8d6o", sides: 6, modifier: open, modifierValue: 6}},
		{"10d10o", SimpleExpression{numDices: 10, expressionText: "10d10o", sides: 10, modifier: open, modifierValue: 10}},
		// More complex expressions (omiting the number of dices -> 1 dice)
		{"d6o", SimpleExpression{numDices: 1, expressionText: "d6o", sides: 6, modifier: open, modifierValue: 6}},
		{"d6e", SimpleExpression{numDices: 1, expressionText: "d6e", sides: 6, modifier: explode, modifierValue: 6}},
		{"d6es4", SimpleExpression{numDices: 1, expressionText: "d6es4", sides: 6, modifier: explodingSuccess, modifierValue: 4}},
		{"d100es96", SimpleExpression{numDices: 1, expressionText: "d100es96", sides: 100, modifier: explodingSuccess, modifierValue: 96}},
		{"d100k1", SimpleExpression{numDices: 1, expressionText: "d100k1", sides: 100, modifier: keep, modifierValue: 1}},
		// Expressions wiht 0
		{"0", SimpleExpression{expressionText: "0", constant: 0}},
		{"d0", SimpleExpression{numDices: 1, expressionText: "d0", sides: 0}},
		{"0d6o", SimpleExpression{numDices: 0, expressionText: "0d6o", sides: 6, modifier: open, modifierValue: 6}},
		{"d0e", SimpleExpression{numDices: 1, expressionText: "d0e", sides: 0, modifier: explode, modifierValue: 0}},
		{"d0e0", SimpleExpression{numDices: 1, expressionText: "d0e0", sides: 0, modifier: explode, modifierValue: 0}},
		{"d6es0", SimpleExpression{numDices: 1, expressionText: "d6es0", sides: 6, modifier: explodingSuccess, modifierValue: 0}},

		/*
			{"10 00", []string{"", ""}, 0, 0},
			// More complex expressions
			//{"3d3d3", []string{"3", "3", "d3"}, 3, 3},
			//{"3d4d5", []string{"3", "4", "d5"}, 3, 4},
			//{"3d4d5+2", []string{"3", "4", "d5+2"}, 3, 4},
			//{"2d6d1", []string{"2", "6", "d1"}, 2, 6},   // rolls two six-sided dice, drops the lowest roll, and sums the total
			{"5d6s4", []string{"5", "6", "s4"}, 5, 6},   // rolls five six-sided dice, and counts any individual roll that exceeds four, presenting the number of "targetes"
		*/
	}
	for i, pts := range parseTestStrings {
		sde := SimpleExpression{expressionText: pts.s}
		err := sde.parse()
		//fmt.Printf("sde: %+v", sde)
		if err != nil {
			t.Errorf("%d) Failed to parse %s: %v", i, sde.expressionText, err)
		}
		if sde != pts.out {
			t.Errorf("%d) Failed to parse %s: %#v != %#v", i, pts.s, &sde, &pts.out)
		}
	}
}

//Roll() (DiceExpressionResult)
// Test that Roll generates the correct results
func TestRoll(t *testing.T) {
	var parseTestStrings = []struct {
		s   string
		out simpleExpressionResult
	}{
		// TODO: test with spaces and invalid dice expressions
		// Constants
		{"1", simpleExpressionResult{diceExpression: SimpleExpression{expressionText: "1", constant: 1}}},
		{"10", simpleExpressionResult{diceExpression: SimpleExpression{expressionText: "10", constant: 10}}},
		{"100", simpleExpressionResult{diceExpression: SimpleExpression{expressionText: "100", constant: 100}}},
		{"1000", simpleExpressionResult{diceExpression: SimpleExpression{expressionText: "1000", constant: 1000}}},
		{"4321", simpleExpressionResult{diceExpression: SimpleExpression{expressionText: "4321", constant: 4321}}},
		{" 1000", simpleExpressionResult{diceExpression: SimpleExpression{expressionText: "1000", constant: 1000}}},
		{"1000 ", simpleExpressionResult{diceExpression: SimpleExpression{expressionText: "1000", constant: 1000}}},
		// Basic dices
		{"d2", simpleExpressionResult{diceExpression: SimpleExpression{numDices: 1, expressionText: "d2", sides: 2}}},
		{"d4", simpleExpressionResult{diceExpression: SimpleExpression{numDices: 1, expressionText: "d4", sides: 4}}},
		{"d6", simpleExpressionResult{diceExpression: SimpleExpression{numDices: 1, expressionText: "d6", sides: 6}}},
		{"d8", simpleExpressionResult{diceExpression: SimpleExpression{numDices: 1, expressionText: "d8", sides: 8}}},
		{"d10", simpleExpressionResult{diceExpression: SimpleExpression{numDices: 1, expressionText: "d10", sides: 10}}},
		{"d12", simpleExpressionResult{diceExpression: SimpleExpression{numDices: 1, expressionText: "d12", sides: 12}}},
		{"d20", simpleExpressionResult{diceExpression: SimpleExpression{numDices: 1, expressionText: "d20", sides: 20}}},
		{"d100", simpleExpressionResult{diceExpression: SimpleExpression{numDices: 1, expressionText: "d100", sides: 100}}},
		{"d200", simpleExpressionResult{diceExpression: SimpleExpression{numDices: 1, expressionText: "d200", sides: 200}}},
		{"d1000", simpleExpressionResult{diceExpression: SimpleExpression{numDices: 1, expressionText: "d1000", sides: 1000}}},
		// More complex expressions
		{"3d3", simpleExpressionResult{diceExpression: SimpleExpression{numDices: 3, expressionText: "3d3", sides: 3}}},
		{"3d6", simpleExpressionResult{diceExpression: SimpleExpression{numDices: 3, expressionText: "3d6", sides: 6}}},
		{"1d2", simpleExpressionResult{diceExpression: SimpleExpression{numDices: 1, expressionText: "1d2", sides: 2}}},
		{"3d6k2", simpleExpressionResult{diceExpression: SimpleExpression{numDices: 3, expressionText: "3d6k2", sides: 6, modifier: keep, modifierValue: 2}}},
		{"3d6kl2", simpleExpressionResult{diceExpression: SimpleExpression{numDices: 3, expressionText: "3d6kl2", sides: 6, modifier: keepLower, modifierValue: 2}}},
		{"4d8r3", simpleExpressionResult{diceExpression: SimpleExpression{numDices: 4, expressionText: "4d8r3", sides: 8, modifier: reroll, modifierValue: 3}}},
		{"4d8s6", simpleExpressionResult{diceExpression: SimpleExpression{numDices: 4, expressionText: "4d8s6", sides: 8, modifier: success, modifierValue: 6}}},
		{"5d4e", simpleExpressionResult{diceExpression: SimpleExpression{numDices: 5, expressionText: "5d4e", sides: 4, modifier: explode, modifierValue: 4}}},
		{"7d6es8", simpleExpressionResult{diceExpression: SimpleExpression{numDices: 7, expressionText: "7d6es8", sides: 6, modifier: explodingSuccess, modifierValue: 8}}},
		{"8d6o", simpleExpressionResult{diceExpression: SimpleExpression{numDices: 8, expressionText: "8d6o", sides: 6, modifier: open, modifierValue: 6}}},
		{"10d10o", simpleExpressionResult{diceExpression: SimpleExpression{numDices: 10, expressionText: "10d10o", sides: 10, modifier: open, modifierValue: 10}}},
		// More complex expressions (omiting the number of dices -> 1 dice)
		{"d6o", simpleExpressionResult{diceExpression: SimpleExpression{numDices: 1, expressionText: "d6o", sides: 6, modifier: open, modifierValue: 6}}},
		{"d6e", simpleExpressionResult{diceExpression: SimpleExpression{numDices: 1, expressionText: "d6e", sides: 6, modifier: explode, modifierValue: 6}}},
		{"d6e4", simpleExpressionResult{diceExpression: SimpleExpression{numDices: 1, expressionText: "d6e4", sides: 6, modifier: explode, modifierValue: 4}}},
		{"d6es4", simpleExpressionResult{diceExpression: SimpleExpression{numDices: 1, expressionText: "d6es4", sides: 6, modifier: explodingSuccess, modifierValue: 4}}},
		{"d100es96", simpleExpressionResult{diceExpression: SimpleExpression{numDices: 1, expressionText: "d100es96", sides: 100, modifier: explodingSuccess, modifierValue: 96}}},
		{"d100e96", simpleExpressionResult{diceExpression: SimpleExpression{numDices: 1, expressionText: "d100e96", sides: 100, modifier: explode, modifierValue: 96}}},
		{"d100k1", simpleExpressionResult{diceExpression: SimpleExpression{numDices: 1, expressionText: "d100k1", sides: 100, modifier: keep, modifierValue: 1}}},
		// Expressions wiht 0
		{"0", simpleExpressionResult{diceExpression: SimpleExpression{expressionText: "0", constant: 0}}},
		{"d0", simpleExpressionResult{diceExpression: SimpleExpression{numDices: 1, expressionText: "d0", sides: 0}}},
		{"0d6o", simpleExpressionResult{diceExpression: SimpleExpression{numDices: 0, expressionText: "0d6o", sides: 6, modifier: open, modifierValue: 6}}},
		{"d0e", simpleExpressionResult{diceExpression: SimpleExpression{numDices: 1, expressionText: "d0e", sides: 0, modifier: explode, modifierValue: 0}}},
		{"d0e0", simpleExpressionResult{diceExpression: SimpleExpression{numDices: 1, expressionText: "d0e0", sides: 0, modifier: explode, modifierValue: 0}}},
		{"d6es0", simpleExpressionResult{diceExpression: SimpleExpression{numDices: 1, expressionText: "d6es0", sides: 6, modifier: explodingSuccess, modifierValue: 0}}},

		/*
			{"10 00", []string{"", ""}, 0, 0},
			// More complex expressions
			//{"3d3d3", []string{"3", "3", "d3"}, 3, 3},
			//{"3d4d5", []string{"3", "4", "d5"}, 3, 4},
			//{"3d4d5+2", []string{"3", "4", "d5+2"}, 3, 4},
			//{"2d6d1", []string{"2", "6", "d1"}, 2, 6},   // rolls two six-sided dice, drops the lowest roll, and sums the total
			{"5d6s4", []string{"5", "6", "s4"}, 5, 6},   // rolls five six-sided dice, and counts any individual roll that exceeds four, presenting the number of "targetes"
		*/
	}
	for i, pts := range parseTestStrings {
		sde := SimpleExpression{expressionText: pts.s}
		_, err := sde.Roll()
		//fmt.Printf("sde: %+v", sde)
		if err != nil {
			t.Errorf("%d) Failed to parse %s: %v", i, sde.expressionText, err)
		}
		/*if sde != pts.out {
			t.Errorf("%d) Failed to parse %s: %#v != %#v", i, pts.s, &sde, &pts.out)
		}*/
	}
}
