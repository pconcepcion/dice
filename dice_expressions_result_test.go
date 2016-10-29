// Package dice provides tools to develop rpg games
package dice

import (
	"testing"
)

// Helper Functions

// assertEcualResults compares two Results slices and returns true if both are have the same content
func assertEqualDiceReults(a, b Results) bool {
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

// TestResultsLen test
func TestResultsLen(t *testing.T) {
	var ResultsTests = []struct {
		dr  Results
		out int
	}{
		{Results{}, 0},
		{Results{0}, 1},
		{Results{1, 2}, 2},
		{Results{4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4}, 20},
	}
	for i, drt := range ResultsTests {
		res := drt.dr.Len()
		if res != drt.out {
			t.Errorf("%d) Len KO expected  %d: got %d", i, drt.dr, res)
		} else {
			t.Logf("%d) Len %v = %d -> OK", i, drt.dr, res)
		}
	}
}

// TestResultsSwap test
func TestResultsSwap(t *testing.T) {
	var ResultsTests = []struct {
		dr  Results
		i   int
		j   int
		out Results
	}{
		{Results{1, 2}, 0, 1, Results{2, 1}},
		{Results{1, 2, 3, 4, 5, 6, 7}, 0, 0, Results{1, 2, 3, 4, 5, 6, 7}},
		{Results{1, 2, 3, 4, 5, 6, 7}, 3, 4, Results{1, 2, 3, 5, 4, 6, 7}},
		{Results{1, 2, 3, 4, 5, 6, 7}, 2, 2, Results{1, 2, 3, 4, 5, 6, 7}},
		{Results{1, 2, 3, 4, 5, 6, 7}, 1, 5, Results{1, 6, 3, 4, 5, 2, 7}},
	}
	for i, drt := range ResultsTests {
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

// TestResultsLess test
func TestResultsLess(t *testing.T) {
	var ResultsTests = []struct {
		dr  Results
		i   int
		j   int
		out bool
	}{
		{Results{1, 2}, 0, 1, true},
		{Results{1, 2, 3, 4, 5, 6, 7}, 0, 0, false},
		{Results{1, 2, 3, 4, 5, 6, 7}, 0, 0, false},
		{Results{1, 2, 3, 4, 5, 6, 7}, 2, 3, true},
		{Results{1, 2, 3, 4, 5, 6, 7}, 6, 0, false},
		{Results{1, 2, 3, 4, 5, 6, 7}, 0, 6, true},
	}
	for i, drt := range ResultsTests {
		res := drt.dr.Less(drt.i, drt.j)
		if res != drt.out {
			t.Errorf("%d) comparing %d and %d expected  %t: got %t", i, drt.dr[drt.i], drt.dr[drt.j], drt.out, res)
		} else {
			t.Logf("%d) comparing g dr[%d] and dr[%d] ->  %d < %d -> %v -> OK", i, drt.i, drt.j, drt.dr[drt.i], drt.dr[drt.j], drt.out)
		}
	}
}

// TestResultsUSum test
func TestResultsSum(t *testing.T) {
	var ResultsTests = []struct {
		dr  Results
		out int
	}{
		{Results{1, 2}, 3},
		{Results{1, 2, 3, 4, 5, 6, 7}, 28},
		{Results{1, 2, 3, 4, 5, 6}, 21},
		{Results{}, 0},
		{Results{5}, 5},
	}
	for i, drt := range ResultsTests {
		res := drt.dr.Sum()
		if res != drt.out {
			t.Errorf("%d) sum of the elements of %d expected  %d: got %d", i, drt.dr, drt.out, res)
		} else {
			t.Logf("%d) sum of the elements of  %v -> %d -> OK", i, drt.dr, drt.out)
		}
	}
}

// TestSimpleExpressionResultSuccess test
func TestSimpleExpressionResultSuccess(t *testing.T) {
	var simpleExpressionResultTests = []struct {
		sder   simpleExpressionResult
		target int
		out    int
	}{
		{simpleExpressionResult{SimpleExpression{}, Results{7}, Results{}, 0, false}, 4, 1},
		{simpleExpressionResult{SimpleExpression{}, Results{}, Results{}, 0, false}, 4, 0},
		{simpleExpressionResult{SimpleExpression{}, Results{7}, Results{}, 0, false}, 10, 0},
		{simpleExpressionResult{SimpleExpression{}, Results{1, 2, 3, 4, 5, 6, 7}, Results{}, 0, false}, 4, 4},
		{simpleExpressionResult{SimpleExpression{}, Results{1, 2, 3, 4, 5, 6, 7}, Results{}, 0, false}, 6, 2},
		{simpleExpressionResult{SimpleExpression{}, Results{1, 2, 3, 4, 5, 6, 7}, Results{}, 0, false}, 6, 2},
		{simpleExpressionResult{SimpleExpression{}, Results{1, 2, 3, 4, 5, 6, 7}, Results{}, 0, false}, 2, 6},
		{simpleExpressionResult{SimpleExpression{}, Results{1, 2, 3, 4, 5, 6, 7}, Results{}, 0, false}, 9, 0},
		{simpleExpressionResult{SimpleExpression{}, Results{1, 2, 3, 4, 5, 6, 7}, Results{}, 0, false}, 0, 7},
	}
	for i, sdert := range simpleExpressionResultTests {
		sdert.sder.Success(sdert.target)
		if sdert.sder.total != sdert.out {
			t.Errorf("%d) expression: %v target: %d expected  %d: got %d", i, sdert.sder, sdert.target, sdert.out, sdert.sder.total)
		} else {
			t.Logf("%d) expression: %v target: %d  num success %d: OK", i, sdert.sder, sdert.target, sdert.sder.total)
		}
	}
}

// TestSimpleExpressionResultSumTotal test
func TestSimpleExpressionResultSumTotal(t *testing.T) {
	var simpleExpressionResultTests = []struct {
		sder simpleExpressionResult
		out  int
	}{
		{simpleExpressionResult{SimpleExpression{}, Results{}, Results{}, 0, false}, 0},
		{simpleExpressionResult{SimpleExpression{}, Results{7}, Results{}, 0, false}, 7},
		{simpleExpressionResult{SimpleExpression{}, Results{1, 2, 3, 4, 5, 6, 7}, Results{}, 0, false}, 28},
		{simpleExpressionResult{SimpleExpression{}, Results{1, 2, 3, 4, 5, 6, 7}, Results{1, 2, 3}, 0, false}, 28},
	}
	for i, sdert := range simpleExpressionResultTests {
		sdert.sder.SumTotal()
		if sdert.sder.total != sdert.out {
			t.Errorf("%d) expression: %v expected  %d: got %d", i, sdert.sder, sdert.out, sdert.sder.total)
		} else {
			t.Logf("%d) expression: %v Sum Total %d: OK", i, sdert.sder, sdert.sder.total)
		}
	}
}

// TestSimpleExpressionResultExplodeDice test
func TestSimpleExpressionResultExplodeDice(t *testing.T) {
	var simpleExpressionResultTests = []simpleExpressionResult{
		simpleExpressionResult{SimpleExpression{sides: 6, modifierValue: 1}, Results{1}, Results{}, 0, false},
		simpleExpressionResult{SimpleExpression{sides: 6, modifierValue: 4}, Results{4, 3}, Results{}, 0, false},
		simpleExpressionResult{SimpleExpression{sides: 4, modifierValue: 4}, Results{4, 3}, Results{}, 0, false},
		simpleExpressionResult{SimpleExpression{sides: 100, modifierValue: 96}, Results{97, 3}, Results{}, 0, false},
	}
	for i, sdert := range simpleExpressionResultTests {
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

// TestSimpleExpressionResultOpen test
func TestSimpleExpressionResultOpen(t *testing.T) {
	var simpleExpressionResultTests = []struct {
		sder simpleExpressionResult
		out  int
	}{
		{simpleExpressionResult{SimpleExpression{sides: 6, modifierValue: 6}, Results{}, Results{}, 0, false}, 0},
		{simpleExpressionResult{SimpleExpression{sides: 6, modifierValue: 6}, Results{1}, Results{}, 0, false}, 0},
		{simpleExpressionResult{SimpleExpression{sides: 6, modifierValue: 6}, Results{6, 1}, Results{}, 0, false}, 1},
		{simpleExpressionResult{SimpleExpression{sides: 6, modifierValue: 6}, Results{6, 6, 6, 6, 3, 3, 3, 3, 1}, Results{}, 0, false}, 4},
	}
	for i, sdert := range simpleExpressionResultTests {
		sdert.sder.Open()
		if len(sdert.sder.extraResults) < sdert.out {
			t.Errorf("%d) expression: %v expected length %d: got %d", i, sdert.sder, sdert.out, len(sdert.sder.extraResults))
			continue
		}
		ok := true
		for j, r := range sdert.sder.Results[:sdert.out] {
			/*if sdert.sder.Results[j] != (sdert.sder.diceExpression.sides + r) {
				t.Errorf("%d) diceResut[%d] expected to be %d: got %d -> expression: %v KO", i, j, sdert.sder.diceExpression.sides+r, sdert.sder.Results[j], sdert.sder)
				ok = false
				break
			}*/
			if sdert.sder.Results[j] <= sdert.sder.diceExpression.sides {
				t.Errorf("%d) diceResuts[%d] expected to be greater than %d got %d -> expression: %v KO", i, j, sdert.sder.diceExpression.sides, r, sdert.sder)
				ok = false
				break
			}
		}
		if ok {
			t.Logf("%d) expression: %v Open %v: OK", i, sdert.sder, sdert.sder.extraResults)
		}
	}
}

// TestSimpleExpressionResultExplode test
func TestSimpleExpressionResultExplode(t *testing.T) {
	var simpleExpressionResultTests = []struct {
		sder simpleExpressionResult
		out  int
	}{
		{simpleExpressionResult{SimpleExpression{sides: 6, modifierValue: 6}, Results{}, Results{}, 0, false}, 0},
		{simpleExpressionResult{SimpleExpression{sides: 6, modifierValue: 6}, Results{1}, Results{}, 0, false}, 0},
		{simpleExpressionResult{SimpleExpression{sides: 6, modifierValue: 6}, Results{6, 1}, Results{}, 0, false}, 1},
		{simpleExpressionResult{SimpleExpression{sides: 4, modifierValue: 3}, Results{4, 3, 2, 1}, Results{}, 0, false}, 2},
		{simpleExpressionResult{SimpleExpression{sides: 100, modifierValue: 96}, Results{100, 96, 35, 1}, Results{}, 0, false}, 2},
		{simpleExpressionResult{SimpleExpression{sides: 6, modifierValue: 6}, Results{6, 6, 6, 6, 3, 3, 3, 3, 1}, Results{}, 0, false}, 4},
	}
	for i, sdert := range simpleExpressionResultTests {
		sdert.sder.Explode()
		if len(sdert.sder.extraResults) < sdert.out {
			t.Errorf("%d) expression: %v expected length %d: got %d", i, sdert.sder, sdert.out, len(sdert.sder.extraResults))
		} else {
			t.Logf("%d) expression: %v Explode %v: OK", i, sdert.sder, sdert.sder.extraResults)
		}
	}
}

// TestSimpleExpressionResultExplodingSuccess test
func TestSimpleExpressionResultExplodingSuccess(t *testing.T) {
	var simpleExpressionResultTests = []struct {
		sder            simpleExpressionResult
		minTotal        int
		minExtraResults int
	}{
		{simpleExpressionResult{SimpleExpression{sides: 6, modifierValue: 6}, Results{}, Results{}, 0, false}, 0, 0},
		{simpleExpressionResult{SimpleExpression{sides: 6, modifierValue: 6}, Results{1}, Results{}, 0, false}, 0, 0},
		{simpleExpressionResult{SimpleExpression{sides: 6, modifierValue: 6}, Results{6, 1}, Results{}, 0, false}, 1, 1},
		{simpleExpressionResult{SimpleExpression{sides: 4, modifierValue: 3}, Results{4, 3, 2, 1}, Results{}, 0, false}, 2, 1},
		{simpleExpressionResult{SimpleExpression{sides: 100, modifierValue: 96}, Results{100, 96, 35, 1}, Results{}, 0, false}, 2, 1},
		{simpleExpressionResult{SimpleExpression{sides: 6, modifierValue: 6}, Results{6, 6, 6, 6, 3, 3, 3, 3, 1}, Results{}, 0, false}, 4, 4},
	}
	for i, sdert := range simpleExpressionResultTests {
		sdert.sder.ExplodingSuccess(sdert.sder.diceExpression.modifierValue)
		if len(sdert.sder.extraResults) < sdert.minExtraResults {
			t.Errorf("%d) expression: %v expected length %d: got %d", i, sdert.sder, sdert.minExtraResults, len(sdert.sder.extraResults))
			continue
		}
		if sdert.sder.total < sdert.minTotal {
			t.Errorf("%d) expression: %v expected min total of %d: got %d", i, sdert.sder, sdert.minTotal, sdert.sder.total)
		} else {
			t.Logf("%d) expression: %v ExplodingSuccess %v: OK", i, sdert.sder, sdert.sder.extraResults)
		}
	}
}

// TestSimpleExpressionResultReroll test
func TestSimpleExpressionResultReroll(t *testing.T) {
	var simpleExpressionResultTests = []simpleExpressionResult{
		simpleExpressionResult{SimpleExpression{sides: 6, modifierValue: 6}, Results{}, Results{}, 0, false},
		simpleExpressionResult{SimpleExpression{sides: 6, modifierValue: 6}, Results{1}, Results{}, 0, false},
		simpleExpressionResult{SimpleExpression{sides: 6, modifierValue: 6}, Results{6, 1}, Results{}, 0, false},
		simpleExpressionResult{SimpleExpression{sides: 4, modifierValue: 3}, Results{4, 3, 2, 1}, Results{}, 0, false},
		simpleExpressionResult{SimpleExpression{sides: 100, modifierValue: 96}, Results{100, 96, 35, 1}, Results{}, 0, false},
		simpleExpressionResult{SimpleExpression{sides: 6, modifierValue: 6}, Results{6, 6, 6, 6, 3, 3, 3, 3, 1}, Results{}, 0, false},
	}
	for i, sdert := range simpleExpressionResultTests {
		sdert.Reroll(sdert.diceExpression.modifierValue)
		for _, r := range sdert.Results {
			if r < sdert.diceExpression.modifierValue {
				t.Errorf("%d) expression: %v expected min result %d: got %d KO", i, sdert, sdert.diceExpression.modifierValue, r)
				break
			}
		}
		t.Logf("%d) expression: %v Reroll %v: OK", i, sdert, sdert.extraResults)
	}
}

// TestSimpleExpressionResultGetResults test
func TestSimpleExpressionResultGetResults(t *testing.T) {
	var simpleExpressionResultTests = []struct {
		sder simpleExpressionResult
		out  Results
	}{
		{simpleExpressionResult{SimpleExpression{}, Results{7}, Results{}, 0, false}, Results{7}},
		{simpleExpressionResult{SimpleExpression{}, Results{}, Results{}, 0, false}, Results{}},
		{simpleExpressionResult{SimpleExpression{}, Results{1, 2, 3, 4, 5, 6, 7}, Results{}, 0, false}, Results{1, 2, 3, 4, 5, 6, 7}},
	}
	for i, sdert := range simpleExpressionResultTests {
		res := sdert.sder.GetResults()
		if !assertEqualDiceReults(res, sdert.out) {
			t.Errorf("%d) expression: %v expected results  %d: got %d", i, sdert.sder, sdert.out, res)
		} else {
			t.Logf("%d) expression: %v GetResults OK", i, sdert.sder)
		}
	}
}

// TestSimpleExpressionResultGetTotal test
func TestSimpleExpressionResultGetTotal(t *testing.T) {
	var simpleExpressionResultTests = []struct {
		sder simpleExpressionResult
		out  int
	}{
		{simpleExpressionResult{SimpleExpression{}, Results{7}, Results{}, 0, false}, 0},
		{simpleExpressionResult{SimpleExpression{}, Results{}, Results{}, 0, false}, 0},
		{simpleExpressionResult{SimpleExpression{}, Results{1, 2, 3, 4, 5, 6, 7}, Results{}, 28, false}, 28},
	}
	for i, sdert := range simpleExpressionResultTests {
		res := sdert.sder.GetTotal()
		if res != sdert.out {
			t.Errorf("%d) expression: %v expected total %d: got %d", i, sdert.sder, sdert.out, res)
		} else {
			t.Logf("%d) expression: %v GetTotal OK", i, sdert.sder)
		}
	}
}
