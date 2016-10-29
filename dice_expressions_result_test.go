// Package dice provides tools to develop rpg games
package dice

import (
	"testing"
)

// Helper Functions

// assertEcualDiceResults compares two DiceResults slices and returns true if both are have the same content
func assertEqualDiceReults(a, b DiceResults) bool {
	if len(a) != len(b) {
		return false
	}
	for i, dr := range a {
		if dr != b[i] {
			return false
		}
	}
	return true
}

// TestDiceResultsLen test
func TestDiceResultsLen(t *testing.T) {
	var diceResultsTests = []struct {
		dr  DiceResults
		out int
	}{
		{DiceResults{}, 0},
		{DiceResults{0}, 1},
		{DiceResults{1, 2}, 2},
		{DiceResults{4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4}, 20},
	}
	for i, drt := range diceResultsTests {
		res := drt.dr.Len()
		if res != drt.out {
			t.Errorf("%d) Len KO expected  %d: got %d", i, drt.dr, res)
		} else {
			t.Logf("%d) Len %v = %d -> OK", i, drt.dr, res)
		}
	}
}

// TestDiceResultsSwap test
func TestDiceResultsSwap(t *testing.T) {
	var diceResultsTests = []struct {
		dr  DiceResults
		i   int
		j   int
		out DiceResults
	}{
		{DiceResults{1, 2}, 0, 1, DiceResults{2, 1}},
		{DiceResults{1, 2, 3, 4, 5, 6, 7}, 0, 0, DiceResults{1, 2, 3, 4, 5, 6, 7}},
		{DiceResults{1, 2, 3, 4, 5, 6, 7}, 3, 4, DiceResults{1, 2, 3, 5, 4, 6, 7}},
		{DiceResults{1, 2, 3, 4, 5, 6, 7}, 2, 2, DiceResults{1, 2, 3, 4, 5, 6, 7}},
		{DiceResults{1, 2, 3, 4, 5, 6, 7}, 1, 5, DiceResults{1, 6, 3, 4, 5, 2, 7}},
	}
	for i, drt := range diceResultsTests {
		originaldr := drt.dr
		auxdr := &originaldr
		t.Log(auxdr)
		auxdr.Swap(drt.i, drt.j)
		if !assertEqualDiceReults(drt.dr, drt.out) {
			// TODO: fix the output to show the original value properly
			t.Errorf("%d) swapping %d and %d expected  %d: got %d", i, drt.i, drt.j, drt.out, auxdr)
		} else {
			t.Logf("%d) swapping %d and %d on %v = %v -> OK", i, drt.i, drt.j, originaldr, drt.out)
		}
	}
}

// TestDiceResultsLess test
func TestDiceResultsLess(t *testing.T) {
	var diceResultsTests = []struct {
		dr  DiceResults
		i   int
		j   int
		out bool
	}{
		{DiceResults{1, 2}, 0, 1, true},
		{DiceResults{1, 2, 3, 4, 5, 6, 7}, 0, 0, false},
		{DiceResults{1, 2, 3, 4, 5, 6, 7}, 0, 0, false},
		{DiceResults{1, 2, 3, 4, 5, 6, 7}, 2, 3, true},
		{DiceResults{1, 2, 3, 4, 5, 6, 7}, 6, 0, false},
		{DiceResults{1, 2, 3, 4, 5, 6, 7}, 0, 6, true},
	}
	for i, drt := range diceResultsTests {
		res := drt.dr.Less(drt.i, drt.j)
		if res != drt.out {
			t.Errorf("%d) comparing %d and %d expected  %t: got %t", i, drt.dr[drt.i], drt.dr[drt.j], drt.out, res)
		} else {
			t.Logf("%d) comparing g dr[%d] and dr[%d] ->  %d < %d -> %v -> OK", i, drt.i, drt.j, drt.dr[drt.i], drt.dr[drt.j], drt.out)
		}
	}
}

// TestDiceResultsUSum test
func TestDiceResultsSum(t *testing.T) {
	var diceResultsTests = []struct {
		dr  DiceResults
		out int
	}{
		{DiceResults{1, 2}, 3},
		{DiceResults{1, 2, 3, 4, 5, 6, 7}, 28},
		{DiceResults{1, 2, 3, 4, 5, 6}, 21},
		{DiceResults{}, 0},
		{DiceResults{5}, 5},
	}
	for i, drt := range diceResultsTests {
		res := drt.dr.Sum()
		if res != drt.out {
			t.Errorf("%d) sum of the elements of %d expected  %d: got %d", i, drt.dr, drt.out, res)
		} else {
			t.Logf("%d) sum of the elements of  %v -> %d -> OK", i, drt.dr, drt.out)
		}
	}
}

// TestSimpleDiceExpressionResultSuccess test
func TestSimpleDiceExpressionResultSuccess(t *testing.T) {
	var simpleDiceExpressionResultTests = []struct {
		sder   simpleDiceExpressionResult
		target int
		out    int
	}{
		{simpleDiceExpressionResult{SimpleDiceExpression{}, DiceResults{7}, DiceResults{}, 0, false}, 4, 1},
		{simpleDiceExpressionResult{SimpleDiceExpression{}, DiceResults{}, DiceResults{}, 0, false}, 4, 0},
		{simpleDiceExpressionResult{SimpleDiceExpression{}, DiceResults{7}, DiceResults{}, 0, false}, 10, 0},
		{simpleDiceExpressionResult{SimpleDiceExpression{}, DiceResults{1, 2, 3, 4, 5, 6, 7}, DiceResults{}, 0, false}, 4, 4},
		{simpleDiceExpressionResult{SimpleDiceExpression{}, DiceResults{1, 2, 3, 4, 5, 6, 7}, DiceResults{}, 0, false}, 6, 2},
		{simpleDiceExpressionResult{SimpleDiceExpression{}, DiceResults{1, 2, 3, 4, 5, 6, 7}, DiceResults{}, 0, false}, 6, 2},
		{simpleDiceExpressionResult{SimpleDiceExpression{}, DiceResults{1, 2, 3, 4, 5, 6, 7}, DiceResults{}, 0, false}, 2, 6},
		{simpleDiceExpressionResult{SimpleDiceExpression{}, DiceResults{1, 2, 3, 4, 5, 6, 7}, DiceResults{}, 0, false}, 9, 0},
		{simpleDiceExpressionResult{SimpleDiceExpression{}, DiceResults{1, 2, 3, 4, 5, 6, 7}, DiceResults{}, 0, false}, 0, 7},
	}
	for i, sdert := range simpleDiceExpressionResultTests {
		sdert.sder.Success(sdert.target)
		if sdert.sder.total != sdert.out {
			t.Errorf("%d) expression: %v target: %d expected  %d: got %d", i, sdert.sder, sdert.target, sdert.out, sdert.sder.total)
		} else {
			t.Logf("%d) expression: %v target: %d  num success %d: OK", i, sdert.sder, sdert.target, sdert.sder.total)
		}
	}
}

// TestSimpleDiceExpressionResultSumTotal test
func TestSimpleDiceExpressionResultSumTotal(t *testing.T) {
	var simpleDiceExpressionResultTests = []struct {
		sder simpleDiceExpressionResult
		out  int
	}{
		{simpleDiceExpressionResult{SimpleDiceExpression{}, DiceResults{}, DiceResults{}, 0, false}, 0},
		{simpleDiceExpressionResult{SimpleDiceExpression{}, DiceResults{7}, DiceResults{}, 0, false}, 7},
		{simpleDiceExpressionResult{SimpleDiceExpression{}, DiceResults{1, 2, 3, 4, 5, 6, 7}, DiceResults{}, 0, false}, 28},
		{simpleDiceExpressionResult{SimpleDiceExpression{}, DiceResults{1, 2, 3, 4, 5, 6, 7}, DiceResults{1, 2, 3}, 0, false}, 28},
	}
	for i, sdert := range simpleDiceExpressionResultTests {
		sdert.sder.SumTotal()
		if sdert.sder.total != sdert.out {
			t.Errorf("%d) expression: %v expected  %d: got %d", i, sdert.sder, sdert.out, sdert.sder.total)
		} else {
			t.Logf("%d) expression: %v Sum Total %d: OK", i, sdert.sder, sdert.sder.total)
		}
	}
}

// TestSimpleDiceExpressionResultExplodeDice test
func TestSimpleDiceExpressionResultExplodeDice(t *testing.T) {
	var simpleDiceExpressionResultTests = []simpleDiceExpressionResult{
		simpleDiceExpressionResult{SimpleDiceExpression{sides: 6, modifierValue: 1}, DiceResults{1}, DiceResults{}, 0, false},
		simpleDiceExpressionResult{SimpleDiceExpression{sides: 6, modifierValue: 4}, DiceResults{4, 3}, DiceResults{}, 0, false},
		simpleDiceExpressionResult{SimpleDiceExpression{sides: 4, modifierValue: 4}, DiceResults{4, 3}, DiceResults{}, 0, false},
		simpleDiceExpressionResult{SimpleDiceExpression{sides: 100, modifierValue: 96}, DiceResults{97, 3}, DiceResults{}, 0, false},
	}
	for i, sdert := range simpleDiceExpressionResultTests {
		dr := sdert.explodeDice()
		if len(dr) <= 0 {
			t.Errorf("%d) expression: len(dr) = %d should be greater than 0: ", i, len(dr))
			continue
		}
		threshold := sdert.diceExpression.modifierValue
		numResults := 1
		if threshold <= 1 {
			numResults = 101
		} else {
			for _, r := range dr {
				if r >= threshold {
					numResults++
				}
			}
		}
		if len(dr) != numResults {
			t.Errorf("%d) expression: %v, explossion Results: %v,  expected length  %d: got %d", i, sdert, dr, numResults, len(dr))
		} else {
			t.Logf("%d) expression: %v explossion Results %v: OK", i, sdert, dr)
		}
	}
}

// TestSimpleDiceExpressionResultOpen test
func TestSimpleDiceExpressionResultOpen(t *testing.T) {
	var simpleDiceExpressionResultTests = []struct {
		sder simpleDiceExpressionResult
		out  int
	}{
		{simpleDiceExpressionResult{SimpleDiceExpression{sides: 6, modifierValue: 6}, DiceResults{}, DiceResults{}, 0, false}, 0},
		{simpleDiceExpressionResult{SimpleDiceExpression{sides: 6, modifierValue: 6}, DiceResults{1}, DiceResults{}, 0, false}, 0},
		{simpleDiceExpressionResult{SimpleDiceExpression{sides: 6, modifierValue: 6}, DiceResults{6, 1}, DiceResults{}, 0, false}, 1},
		{simpleDiceExpressionResult{SimpleDiceExpression{sides: 6, modifierValue: 6}, DiceResults{6, 6, 6, 6, 3, 3, 3, 3, 1}, DiceResults{}, 0, false}, 4},
	}
	for i, sdert := range simpleDiceExpressionResultTests {
		sdert.sder.Open()
		if len(sdert.sder.extraDiceResults) < sdert.out {
			t.Errorf("%d) expression: %v expected length %d: got %d", i, sdert.sder, sdert.out, len(sdert.sder.extraDiceResults))
			continue
		}
		ok := true
		for j, r := range sdert.sder.diceResults[:sdert.out] {
			/*if sdert.sder.diceResults[j] != (sdert.sder.diceExpression.sides + r) {
				t.Errorf("%d) diceResut[%d] expected to be %d: got %d -> expression: %v KO", i, j, sdert.sder.diceExpression.sides+r, sdert.sder.diceResults[j], sdert.sder)
				ok = false
				break
			}*/
			if sdert.sder.diceResults[j] <= sdert.sder.diceExpression.sides {
				t.Errorf("%d) diceResuts[%d] expected to be greater than %d got %d -> expression: %v KO", i, j, sdert.sder.diceExpression.sides, r, sdert.sder)
				ok = false
				break
			}
		}
		if ok {
			t.Logf("%d) expression: %v Open %v: OK", i, sdert.sder, sdert.sder.extraDiceResults)
		}
	}
}

// TestSimpleDiceExpressionResultExplode test
func TestSimpleDiceExpressionResultExplode(t *testing.T) {
	var simpleDiceExpressionResultTests = []struct {
		sder simpleDiceExpressionResult
		out  int
	}{
		{simpleDiceExpressionResult{SimpleDiceExpression{sides: 6, modifierValue: 6}, DiceResults{}, DiceResults{}, 0, false}, 0},
		{simpleDiceExpressionResult{SimpleDiceExpression{sides: 6, modifierValue: 6}, DiceResults{1}, DiceResults{}, 0, false}, 0},
		{simpleDiceExpressionResult{SimpleDiceExpression{sides: 6, modifierValue: 6}, DiceResults{6, 1}, DiceResults{}, 0, false}, 1},
		{simpleDiceExpressionResult{SimpleDiceExpression{sides: 4, modifierValue: 3}, DiceResults{4, 3, 2, 1}, DiceResults{}, 0, false}, 2},
		{simpleDiceExpressionResult{SimpleDiceExpression{sides: 100, modifierValue: 96}, DiceResults{100, 96, 35, 1}, DiceResults{}, 0, false}, 2},
		{simpleDiceExpressionResult{SimpleDiceExpression{sides: 6, modifierValue: 6}, DiceResults{6, 6, 6, 6, 3, 3, 3, 3, 1}, DiceResults{}, 0, false}, 4},
	}
	for i, sdert := range simpleDiceExpressionResultTests {
		sdert.sder.Explode()
		if len(sdert.sder.extraDiceResults) < sdert.out {
			t.Errorf("%d) expression: %v expected length %d: got %d", i, sdert.sder, sdert.out, len(sdert.sder.extraDiceResults))
		} else {
			t.Logf("%d) expression: %v Explode %v: OK", i, sdert.sder, sdert.sder.extraDiceResults)
		}
	}
}

// TestSimpleDiceExpressionResultExplodingSuccess test
func TestSimpleDiceExpressionResultExplodingSuccess(t *testing.T) {
	var simpleDiceExpressionResultTests = []struct {
		sder            simpleDiceExpressionResult
		minTotal        int
		minExtraResults int
	}{
		{simpleDiceExpressionResult{SimpleDiceExpression{sides: 6, modifierValue: 6}, DiceResults{}, DiceResults{}, 0, false}, 0, 0},
		{simpleDiceExpressionResult{SimpleDiceExpression{sides: 6, modifierValue: 6}, DiceResults{1}, DiceResults{}, 0, false}, 0, 0},
		{simpleDiceExpressionResult{SimpleDiceExpression{sides: 6, modifierValue: 6}, DiceResults{6, 1}, DiceResults{}, 0, false}, 1, 1},
		{simpleDiceExpressionResult{SimpleDiceExpression{sides: 4, modifierValue: 3}, DiceResults{4, 3, 2, 1}, DiceResults{}, 0, false}, 2, 1},
		{simpleDiceExpressionResult{SimpleDiceExpression{sides: 100, modifierValue: 96}, DiceResults{100, 96, 35, 1}, DiceResults{}, 0, false}, 2, 1},
		{simpleDiceExpressionResult{SimpleDiceExpression{sides: 6, modifierValue: 6}, DiceResults{6, 6, 6, 6, 3, 3, 3, 3, 1}, DiceResults{}, 0, false}, 4, 4},
	}
	for i, sdert := range simpleDiceExpressionResultTests {
		sdert.sder.ExplodingSuccess(sdert.sder.diceExpression.modifierValue)
		if len(sdert.sder.extraDiceResults) < sdert.minExtraResults {
			t.Errorf("%d) expression: %v expected length %d: got %d", i, sdert.sder, sdert.minExtraResults, len(sdert.sder.extraDiceResults))
			continue
		}
		if sdert.sder.total < sdert.minTotal {
			t.Errorf("%d) expression: %v expected min total of %d: got %d", i, sdert.sder, sdert.minTotal, sdert.sder.total)
		} else {
			t.Logf("%d) expression: %v ExplodingSuccess %v: OK", i, sdert.sder, sdert.sder.extraDiceResults)
		}
	}
}

// TestSimpleDiceExpressionResultReroll test
func TestSimpleDiceExpressionResultReroll(t *testing.T) {
	var simpleDiceExpressionResultTests = []simpleDiceExpressionResult{
		simpleDiceExpressionResult{SimpleDiceExpression{sides: 6, modifierValue: 6}, DiceResults{}, DiceResults{}, 0, false},
		simpleDiceExpressionResult{SimpleDiceExpression{sides: 6, modifierValue: 6}, DiceResults{1}, DiceResults{}, 0, false},
		simpleDiceExpressionResult{SimpleDiceExpression{sides: 6, modifierValue: 6}, DiceResults{6, 1}, DiceResults{}, 0, false},
		simpleDiceExpressionResult{SimpleDiceExpression{sides: 4, modifierValue: 3}, DiceResults{4, 3, 2, 1}, DiceResults{}, 0, false},
		simpleDiceExpressionResult{SimpleDiceExpression{sides: 100, modifierValue: 96}, DiceResults{100, 96, 35, 1}, DiceResults{}, 0, false},
		simpleDiceExpressionResult{SimpleDiceExpression{sides: 6, modifierValue: 6}, DiceResults{6, 6, 6, 6, 3, 3, 3, 3, 1}, DiceResults{}, 0, false},
	}
	for i, sdert := range simpleDiceExpressionResultTests {
		sdert.Reroll(sdert.diceExpression.modifierValue)
		for _, r := range sdert.diceResults {
			if r < sdert.diceExpression.modifierValue {
				t.Errorf("%d) expression: %v expected min result %d: got %d KO", i, sdert, sdert.diceExpression.modifierValue, r)
				break
			}
		}
		t.Logf("%d) expression: %v Reroll %v: OK", i, sdert, sdert.extraDiceResults)
	}
}

// TestSimpleDiceExpressionResultGetResults test
func TestSimpleDiceExpressionResultGetResults(t *testing.T) {
	var simpleDiceExpressionResultTests = []struct {
		sder simpleDiceExpressionResult
		out  DiceResults
	}{
		{simpleDiceExpressionResult{SimpleDiceExpression{}, DiceResults{7}, DiceResults{}, 0, false}, DiceResults{7}},
		{simpleDiceExpressionResult{SimpleDiceExpression{}, DiceResults{}, DiceResults{}, 0, false}, DiceResults{}},
		{simpleDiceExpressionResult{SimpleDiceExpression{}, DiceResults{1, 2, 3, 4, 5, 6, 7}, DiceResults{}, 0, false}, DiceResults{1, 2, 3, 4, 5, 6, 7}},
	}
	for i, sdert := range simpleDiceExpressionResultTests {
		res := sdert.sder.GetResults()
		if !assertEqualDiceReults(res, sdert.out) {
			t.Errorf("%d) expression: %v expected results  %d: got %d", i, sdert.sder, sdert.out, res)
		} else {
			t.Logf("%d) expression: %v GetResults OK", i, sdert.sder)
		}
	}
}

// TestSimpleDiceExpressionResultGetTotal test
func TestSimpleDiceExpressionResultGetTotal(t *testing.T) {
	var simpleDiceExpressionResultTests = []struct {
		sder simpleDiceExpressionResult
		out  int
	}{
		{simpleDiceExpressionResult{SimpleDiceExpression{}, DiceResults{7}, DiceResults{}, 0, false}, 0},
		{simpleDiceExpressionResult{SimpleDiceExpression{}, DiceResults{}, DiceResults{}, 0, false}, 0},
		{simpleDiceExpressionResult{SimpleDiceExpression{}, DiceResults{1, 2, 3, 4, 5, 6, 7}, DiceResults{}, 28, false}, 28},
	}
	for i, sdert := range simpleDiceExpressionResultTests {
		res := sdert.sder.GetTotal()
		if res != sdert.out {
			t.Errorf("%d) expression: %v expected total %d: got %d", i, sdert.sder, sdert.out, res)
		} else {
			t.Logf("%d) expression: %v GetTotal OK", i, sdert.sder)
		}
	}
}
