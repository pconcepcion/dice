// Package rpg provides tools to develop rpg games
package rpg

import (
	"testing"
)

// Test that Parse splits the dices properly
func TestParse(t *testing.T) {
	var parseTestStrings = []struct {
		s   string
		out SimpleDiceExpression
	}{
		// TODO: test with spaces and invalid dice expressions
		// Constants
		{"1", SimpleDiceExpression{expressionText: "1", constant: 1}},
		{"10", SimpleDiceExpression{expressionText: "10", constant: 10}},
		{"100", SimpleDiceExpression{expressionText: "100", constant: 100}},
		{"1000", SimpleDiceExpression{expressionText: "1000", constant: 1000}},
		{"4321", SimpleDiceExpression{expressionText: "4321", constant: 4321}},
		{" 1000", SimpleDiceExpression{expressionText: "1000", constant: 1000}},
		{"1000 ", SimpleDiceExpression{expressionText: "1000", constant: 1000}},
		// Basic dices
		{"d2", SimpleDiceExpression{numDices: 1, expressionText: "d2", sides: 2}},
		{"d4", SimpleDiceExpression{numDices: 1, expressionText: "d4", sides: 4}},
		{"d6", SimpleDiceExpression{numDices: 1, expressionText: "d6", sides: 6}},
		{"d8", SimpleDiceExpression{numDices: 1, expressionText: "d8", sides: 8}},
		{"d10", SimpleDiceExpression{numDices: 1, expressionText: "d10", sides: 10}},
		{"d12", SimpleDiceExpression{numDices: 1, expressionText: "d12", sides: 12}},
		{"d20", SimpleDiceExpression{numDices: 1, expressionText: "d20", sides: 20}},
		{"d100", SimpleDiceExpression{numDices: 1, expressionText: "d100", sides: 100}},
		{"d200", SimpleDiceExpression{numDices: 1, expressionText: "d200", sides: 200}},
		{"d1000", SimpleDiceExpression{numDices: 1, expressionText: "d1000", sides: 1000}},
		// More complex expressions
		{"3d3", SimpleDiceExpression{numDices: 3, expressionText: "3d3", sides: 3}},
		{"3d6", SimpleDiceExpression{numDices: 3, expressionText: "3d6", sides: 6}},
		{"1d2", SimpleDiceExpression{numDices: 1, expressionText: "1d2", sides: 2}},
		{"3d6k2", SimpleDiceExpression{numDices: 3, expressionText: "3d6k2", sides: 6, keep: 2}},
		{"4d8r2", SimpleDiceExpression{numDices: 4, expressionText: "4d8r2", sides: 8, rerollH: 2}},
		{"6d6e", SimpleDiceExpression{numDices: 6, expressionText: "6d6e", sides: 6, explodeResult: 6}},
		{"7d6es8", SimpleDiceExpression{numDices: 7, expressionText: "7d6es8", sides: 6, explodeResult: 6, target: 8}},
		{"8d6o", SimpleDiceExpression{numDices: 8, expressionText: "8d6o", sides: 6, explodeResult: 6}},
		{"10d10o", SimpleDiceExpression{numDices: 10, expressionText: "10d10o", sides: 10, explodeResult: 10}},
		// More complex expressions (omiting the number of dices -> 1 dice)
		{"d6o", SimpleDiceExpression{numDices: 1, expressionText: "d6o", sides: 6, explodeResult: 6}},
		{"d6e", SimpleDiceExpression{numDices: 1, expressionText: "d6e", sides: 6, explodeResult: 6}},
		{"d6es4", SimpleDiceExpression{numDices: 1, expressionText: "d6es4", sides: 6, explodeResult: 6, target: 4}},
		{"d100es96", SimpleDiceExpression{numDices: 1, expressionText: "d100es96", sides: 100, explodeResult: 100, target: 96}},
		{"d100k1", SimpleDiceExpression{numDices: 1, expressionText: "d100k1", sides: 100, keep: 1}},
		/*{"d6o", []Token{{tokenDice, "d"}, {tokenNumber, "6"}, {tokenModifier, "o"}, {tokenEOF, ""}}},
		{"d6e", []Token{{tokenDice, "d"}, {tokenNumber, "6"}, {tokenModifier, "e"}, {tokenEOF, ""}}},
		{"d6es4", []Token{{tokenDice, "d"}, {tokenNumber, "6"}, {tokenModifier, "es"}, {tokenNumber, "4"}, {tokenEOF, ""}}},
		{"d100es96", []Token{{tokenDice, "d"}, {tokenNumber, "100"}, {tokenModifier, "es"}, {tokenNumber, "96"}, {tokenEOF, ""}}},
		{"d100k1", []Token{{tokenDice, "d"}, {tokenNumber, "100"}, {tokenModifier, "k"}, {tokenNumber, "1"}, {tokenEOF, ""}}},

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
		sde := SimpleDiceExpression{expressionText: pts.s}
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
