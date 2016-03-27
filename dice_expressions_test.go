// Package rpg provides tools to develop rpg games
package rpg

import (
	"fmt"
	"testing"
)

// Test that Parse splits the dices properly
func TestParse(t *testing.T) {
	var parseTestStrings = []struct {
		s     string
		out   []string
		dices int
		sides int
	}{
		// TODO: test with spaces and invalid dice expressions
		// Constants
		{"1", []string{"", ""}, 0, 0},
		{"10", []string{"", ""}, 0, 0},
		{"100", []string{"", ""}, 0, 0},
		{"1000", []string{"", ""}, 0, 0},
		{"4321", []string{"", ""}, 0, 0},
		{" 1000", []string{"", ""}, 0, 0},
		{"1000 ", []string{"", ""}, 0, 0},
		{"10 00", []string{"", ""}, 0, 0},
		{" 1000 ", []string{"", ""}, 0, 0},
		// Basic dices
		{"d2", []string{"", "2"}, 1, 2},
		{"d4", []string{"", "4"}, 1, 4},
		{"d6", []string{"", "6"}, 1, 6},
		{"d8", []string{"", "8"}, 1, 8},
		{"d10", []string{"", "10"}, 1, 10},
		{"d12", []string{"", "12"}, 1, 12},
		{"d20", []string{"", "20"}, 1, 20},
		{"d100", []string{"", "100"}, 1, 100},
		{"d200", []string{"", "200"}, 1, 200},
		{"d1000", []string{"", "1000"}, 1, 1000},
		// More complex expressions
		{"3d3", []string{"3", "3"}, 3, 3},
		//{"3d3d3", []string{"3", "3", "d3"}, 3, 3},
		//{"3d4d5", []string{"3", "4", "d5"}, 3, 4},
		//{"3d4d5+2", []string{"3", "4", "d5+2"}, 3, 4},
		{"1d2", []string{"1", "2"}, 1, 2}, // rolls one two sides die and calculates the sum
		//{"2d6d1", []string{"2", "6", "d1"}, 2, 6},   // rolls two six-sided dice, drops the lowest roll, and sums the total
		{"3d6k3", []string{"3", "6", "k3"}, 3, 6},   // rolls thee six-sided dice, keeps the highest 3 rolls, and presents the total
		{"4d8r2", []string{"4", "8", "r2"}, 4, 8},   // rolls four eight-sided dice, repeatedly rerolls any dice that are lower than 2 until all dice rolls are higher than or equal to 2, and then sums and presents the total
		{"5d6s4", []string{"5", "6", "s4"}, 5, 6},   // rolls five six-sided dice, and counts any individual roll that exceeds four, presenting the number of "successes"
		{"6d6e", []string{"6", "6", "e"}, 6, 6},     // rolls eight six-sided dice, and if either rolls a 6, it is rerolled and added to the total (this continues until neither die rolls a 6).
		{"7d6es8", []string{"7", "6", "es8"}, 7, 6}, // will roll seven six-sided dice, explode any that roll their maximum, and then total the rolls that exceed 8
		{"8d6o", []string{"8", "6", "o"}, 8, 6},     // rolls 5 six-sided dice, and explodes any that roll 6
	}

	var sde SimpleDiceExpression
	for i, pts := range parseTestStrings {
		sde.expressionText = pts.s
		res, err := sde.parse()
		//fmt.Printf("sde: %+v", sde)
		if err != nil {
			t.Errorf("%d) Failed to parse %s: %v", i, sde.expressionText, err)
		} else {
			for j, v := range pts.out {
				if v != res[j+1] {
					t.Errorf("%d) Failed to parse %s: %s != %s", i, sde.expressionText, v, res[j+1])
				}
			}
			if sde.sides != pts.sides {
				t.Errorf("%d) Failed to parse %s: sides doesn't match %d != %d", i, sde.expressionText, sde.sides, pts.sides)
			}
			//t.Logf("Parsed %s\n", sde.expressionText)
			fmt.Printf("%d) Parsed %s -> %#v\n", i, sde.expressionText, res)
		}
	}
}
