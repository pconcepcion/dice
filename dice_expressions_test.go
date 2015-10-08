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
		{"d3", []string{"", "3"}, 1, 3},
		{"3d3", []string{"3", "3"}, 3, 3},
		{"3d3d3", []string{"3", "3", "d3"}, 3, 3},
		{"3d4d5", []string{"3", "4", "d5"}, 3, 4},
		{"3d4d5+2", []string{"3", "4", "d5+2"}, 3, 4},
	}
	var sde SimpleDiceExpression
	for i, pts := range parseTestStrings {
		sde.expressionText = pts.s
		res, err := sde.parse()
		if err != nil {
			t.Errorf("%d) Failed to parse %s", i, sde.expressionText)
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
			fmt.Printf("%d) Parsed %s -> %v\n", i, sde.expressionText, res)
		}
	}
}
