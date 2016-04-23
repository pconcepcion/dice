// Package rpg provides tools to develop rpg games package rpg
package rpg

import (
	"fmt"
	"testing"
)

// Helper Functions

// equalTokenSlices compares two token slices and returns true if both are have the same content
func assertEqualTokenSlices(a, b []Token) bool {
	if len(a) != len(b) {
		return false
	}
	// TODO: extracth this to a compare token arrays function
	for i, tok := range a {
		if tok != b[i] {
			return false
		}
	}
	return true
}

// Test that the lexer sends the propper tokens
func TestLexer(t *testing.T) {
	var lexerTestStrings = []struct {
		s   string
		out []Token
	}{
		// Constants
		{"1", []Token{Token{tokenNumber, "1"}, Token{tokenEOF, ""}}},
		{"10", []Token{Token{tokenNumber, "10"}, Token{tokenEOF, ""}}},
		{"100", []Token{Token{tokenNumber, "100"}, Token{tokenEOF, ""}}},
		{"1000", []Token{Token{tokenNumber, "1000"}, Token{tokenEOF, ""}}},
		{"10000", []Token{Token{tokenNumber, "10000"}, Token{tokenEOF, ""}}},
		{"4321", []Token{Token{tokenNumber, "4321"}, Token{tokenEOF, ""}}},
		// Basic dices
		{"d2", []Token{Token{tokenDice, "d"}, Token{tokenNumber, "2"}, Token{tokenEOF, ""}}},
		{"d3", []Token{Token{tokenDice, "d"}, Token{tokenNumber, "3"}, Token{tokenEOF, ""}}},
		{"d4", []Token{Token{tokenDice, "d"}, Token{tokenNumber, "4"}, Token{tokenEOF, ""}}},
		{"d6", []Token{Token{tokenDice, "d"}, Token{tokenNumber, "6"}, Token{tokenEOF, ""}}},
		{"d8", []Token{Token{tokenDice, "d"}, Token{tokenNumber, "8"}, Token{tokenEOF, ""}}},
		{"d10", []Token{Token{tokenDice, "d"}, Token{tokenNumber, "10"}, Token{tokenEOF, ""}}},
		{"d12", []Token{Token{tokenDice, "d"}, Token{tokenNumber, "12"}, Token{tokenEOF, ""}}},
		{"d20", []Token{Token{tokenDice, "d"}, Token{tokenNumber, "20"}, Token{tokenEOF, ""}}},
		{"d100", []Token{Token{tokenDice, "d"}, Token{tokenNumber, "100"}, Token{tokenEOF, ""}}},
		{"d200", []Token{Token{tokenDice, "d"}, Token{tokenNumber, "200"}, Token{tokenEOF, ""}}},
		{"d1000", []Token{Token{tokenDice, "d"}, Token{tokenNumber, "1000"}, Token{tokenEOF, ""}}},
		// More complex expressions
		{"3d3", []Token{Token{tokenNumber, "3"}, Token{tokenDice, "d"}, Token{tokenNumber, "3"}, Token{tokenEOF, ""}}},
		{"3d6", []Token{Token{tokenNumber, "3"}, Token{tokenDice, "d"}, Token{tokenNumber, "6"}, Token{tokenEOF, ""}}},
		{"1d2", []Token{Token{tokenNumber, "1"}, Token{tokenDice, "d"}, Token{tokenNumber, "2"}, Token{tokenEOF, ""}}},
		//{"2d2d1", []Token{Token{tokenNumber, "2"}, Token{tokenDice, "d"}, Token{tokenNumber, "2"}, Token{tokenModifier, "d",}, Token{tokenNumber, "1"},  Token{tokenEOF, ""}}},
		{"3d6k2", []Token{Token{tokenNumber, "3"}, Token{tokenDice, "d"}, Token{tokenNumber, "6"}, Token{tokenModifier, "k"}, Token{tokenNumber, "2"}, Token{tokenEOF, ""}}},
		{"4d8r2", []Token{Token{tokenNumber, "4"}, Token{tokenDice, "d"}, Token{tokenNumber, "8"}, Token{tokenModifier, "r"}, Token{tokenNumber, "2"}, Token{tokenEOF, ""}}},
		{"5d6s4", []Token{Token{tokenNumber, "5"}, Token{tokenDice, "d"}, Token{tokenNumber, "6"}, Token{tokenModifier, "s"}, Token{tokenNumber, "4"}, Token{tokenEOF, ""}}},
		{"6d6e", []Token{Token{tokenNumber, "6"}, Token{tokenDice, "d"}, Token{tokenNumber, "6"}, Token{tokenModifier, "e"}, Token{tokenEOF, ""}}},
		{"7d6es8", []Token{Token{tokenNumber, "7"}, Token{tokenDice, "d"}, Token{tokenNumber, "6"}, Token{tokenModifier, "es"}, Token{tokenNumber, "8"}, Token{tokenEOF, ""}}},
		{"8d6o", []Token{Token{tokenNumber, "8"}, Token{tokenDice, "d"}, Token{tokenNumber, "6"}, Token{tokenModifier, "o"}, Token{tokenEOF, ""}}},
		{"10d10o", []Token{Token{tokenNumber, "10"}, Token{tokenDice, "d"}, Token{tokenNumber, "10"}, Token{tokenModifier, "o"}, Token{tokenEOF, ""}}},
		// More complex expressions (omiting the number of dices -> 1 dice)
		{"d6o", []Token{Token{tokenDice, "d"}, Token{tokenNumber, "6"}, Token{tokenModifier, "o"}, Token{tokenEOF, ""}}},
		{"d6e", []Token{Token{tokenDice, "d"}, Token{tokenNumber, "6"}, Token{tokenModifier, "e"}, Token{tokenEOF, ""}}},
		{"d6es4", []Token{Token{tokenDice, "d"}, Token{tokenNumber, "6"}, Token{tokenModifier, "es"}, Token{tokenNumber, "4"}, Token{tokenEOF, ""}}},
		{"d100es96", []Token{Token{tokenDice, "d"}, Token{tokenNumber, "100"}, Token{tokenModifier, "es"}, Token{tokenNumber, "96"}, Token{tokenEOF, ""}}},
		{"d100k1", []Token{Token{tokenDice, "d"}, Token{tokenNumber, "100"}, Token{tokenModifier, "k"}, Token{tokenNumber, "1"}, Token{tokenEOF, ""}}},

		// Some errors:
		{" 10000", []Token{Token{tokenError, " 1"}}},
		{"1000 ", []Token{Token{tokenNumber, "1000"}, Token{tokenError, " "}}},
		{"10 000", []Token{Token{tokenNumber, "10"}, Token{tokenError, " "}}},
		{"d6a", []Token{Token{tokenDice, "d"}, Token{tokenNumber, "6"}, Token{tokenError, "a"}}},
		{"5d6v4", []Token{Token{tokenNumber, "5"}, Token{tokenDice, "d"}, Token{tokenNumber, "6"}, Token{tokenError, "v"}}},
		// TODO: more test with spaces and invalid dice expressions
		/*
			// More complex expressions
			//{"3d4d5+2", []string{"3", "4", "d5+2"}, 3, 4},
			//{"2d6d1", []string{"2", "6", "d1"}, 2, 6},   // rolls two six-sided dice, drops the lowest roll, and sums the total
		*/
	}
	var resultTokens = []Token{}
	for i, lts := range lexerTestStrings {
		_, c := lex(lts.s)
		for tok := range c {
			resultTokens = append(resultTokens, tok)
		}

		if !assertEqualTokenSlices(lts.out, resultTokens) {
			t.Error("Expected value: ", lts.out, " obtaned: ", resultTokens)
			fmt.Printf("Lexer Test %d KO. Expected value: %v obtained  %v\n", i, lts.out, resultTokens)
		} else {
			fmt.Printf("Lexer Test %d OK: %v\n", i, resultTokens)
		}
		resultTokens = []Token{}
	}
}

// Test for the isWhitespace function
func TestIsWhitespace(t *testing.T) {
	var testStrings = []struct {
		s rune
		b bool
	}{
		{' ', true},
		{'\t', true},
		{'\n', true},
		{'a', false},
		{'ä', false},
		{'0', false},
		{'9', false},
		{'本', false},
		{'\000', false},
		{'\007', false},
		{'\377', false},
		{'\x07', false},
		{'\xff', false},
		{'\u12e4', false},
		{'\U00101234', false},
	}
	var testResult bool

	for _, lts := range testStrings {
		testResult = isWhitespace(lts.s)
		if testResult != lts.b {
			t.Error("Expected value: ", lts.b, " obtaned: ", testResult)
		}
	}
}
