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
		{"1", []Token{{tokenNumber, "1"}, {tokenEOF, ""}}},
		{"10", []Token{{tokenNumber, "10"}, {tokenEOF, ""}}},
		{"100", []Token{{tokenNumber, "100"}, {tokenEOF, ""}}},
		{"1000", []Token{{tokenNumber, "1000"}, {tokenEOF, ""}}},
		{"10000", []Token{{tokenNumber, "10000"}, {tokenEOF, ""}}},
		{"4321", []Token{{tokenNumber, "4321"}, {tokenEOF, ""}}},
		// Basic dices
		{"d2", []Token{{tokenDice, "d"}, {tokenNumber, "2"}, {tokenEOF, ""}}},
		{"d3", []Token{{tokenDice, "d"}, {tokenNumber, "3"}, {tokenEOF, ""}}},
		{"d4", []Token{{tokenDice, "d"}, {tokenNumber, "4"}, {tokenEOF, ""}}},
		{"d6", []Token{{tokenDice, "d"}, {tokenNumber, "6"}, {tokenEOF, ""}}},
		{"d8", []Token{{tokenDice, "d"}, {tokenNumber, "8"}, {tokenEOF, ""}}},
		{"d10", []Token{{tokenDice, "d"}, {tokenNumber, "10"}, {tokenEOF, ""}}},
		{"d12", []Token{{tokenDice, "d"}, {tokenNumber, "12"}, {tokenEOF, ""}}},
		{"d20", []Token{{tokenDice, "d"}, {tokenNumber, "20"}, {tokenEOF, ""}}},
		{"d100", []Token{{tokenDice, "d"}, {tokenNumber, "100"}, {tokenEOF, ""}}},
		{"d200", []Token{{tokenDice, "d"}, {tokenNumber, "200"}, {tokenEOF, ""}}},
		{"d1000", []Token{{tokenDice, "d"}, {tokenNumber, "1000"}, {tokenEOF, ""}}},
		// More complex expressions
		{"3d3", []Token{{tokenNumber, "3"}, {tokenDice, "d"}, {tokenNumber, "3"}, {tokenEOF, ""}}},
		{"3d6", []Token{{tokenNumber, "3"}, {tokenDice, "d"}, {tokenNumber, "6"}, {tokenEOF, ""}}},
		{"1d2", []Token{{tokenNumber, "1"}, {tokenDice, "d"}, {tokenNumber, "2"}, {tokenEOF, ""}}},
		//{"2d2d1", []Token{Token{tokenNumber, "2"}, Token{tokenDice, "d"}, Token{tokenNumber, "2"}, Token{tokenModifier, "d",}, Token{tokenNumber, "1"},  Token{tokenEOF, ""}}},
		{"3d6k2", []Token{{tokenNumber, "3"}, {tokenDice, "d"}, {tokenNumber, "6"}, {tokenModifier, "k"}, {tokenNumber, "2"}, {tokenEOF, ""}}},
		{"3d6kl2", []Token{{tokenNumber, "3"}, {tokenDice, "d"}, {tokenNumber, "6"}, {tokenModifier, "kl"}, {tokenNumber, "2"}, {tokenEOF, ""}}},
		{"4d8r2", []Token{{tokenNumber, "4"}, {tokenDice, "d"}, {tokenNumber, "8"}, {tokenModifier, "r"}, {tokenNumber, "2"}, {tokenEOF, ""}}},
		{"5d6s4", []Token{{tokenNumber, "5"}, {tokenDice, "d"}, {tokenNumber, "6"}, {tokenModifier, "s"}, {tokenNumber, "4"}, {tokenEOF, ""}}},
		{"6d6e", []Token{{tokenNumber, "6"}, {tokenDice, "d"}, {tokenNumber, "6"}, {tokenModifier, "e"}, {tokenEOF, ""}}},
		{"6d6e4", []Token{{tokenNumber, "6"}, {tokenDice, "d"}, {tokenNumber, "6"}, {tokenModifier, "e"}, {tokenNumber, "4"}, {tokenEOF, ""}}},
		{"7d6es8", []Token{{tokenNumber, "7"}, {tokenDice, "d"}, {tokenNumber, "6"}, {tokenModifier, "es"}, {tokenNumber, "8"}, {tokenEOF, ""}}},
		{"8d6o", []Token{{tokenNumber, "8"}, {tokenDice, "d"}, {tokenNumber, "6"}, {tokenModifier, "o"}, {tokenEOF, ""}}},
		{"10d10o", []Token{{tokenNumber, "10"}, {tokenDice, "d"}, {tokenNumber, "10"}, {tokenModifier, "o"}, {tokenEOF, ""}}},
		// More complex expressions (omiting the number of dices -> 1 dice)
		{"d6o", []Token{{tokenDice, "d"}, {tokenNumber, "6"}, {tokenModifier, "o"}, {tokenEOF, ""}}},
		{"d6e", []Token{{tokenDice, "d"}, {tokenNumber, "6"}, {tokenModifier, "e"}, {tokenEOF, ""}}},
		{"d6e4", []Token{{tokenDice, "d"}, {tokenNumber, "6"}, {tokenModifier, "e"}, {tokenNumber, "4"}, {tokenEOF, ""}}},
		{"d6es4", []Token{{tokenDice, "d"}, {tokenNumber, "6"}, {tokenModifier, "es"}, {tokenNumber, "4"}, {tokenEOF, ""}}},
		{"d100es96", []Token{{tokenDice, "d"}, {tokenNumber, "100"}, {tokenModifier, "es"}, {tokenNumber, "96"}, {tokenEOF, ""}}},
		{"d100k1", []Token{{tokenDice, "d"}, {tokenNumber, "100"}, {tokenModifier, "k"}, {tokenNumber, "1"}, {tokenEOF, ""}}},
		{"100d100k90", []Token{{tokenNumber, "100"}, {tokenDice, "d"}, {tokenNumber, "100"}, {tokenModifier, "k"}, {tokenNumber, "90"}, {tokenEOF, ""}}},
		{"100d100e96", []Token{{tokenNumber, "100"}, {tokenDice, "d"}, {tokenNumber, "100"}, {tokenModifier, "e"}, {tokenNumber, "96"}, {tokenEOF, ""}}},

		// Some errors:
		{" 10000", []Token{{tokenError, "unexpected token 49, expected either 'd' or number"}}},
		{"1000 ", []Token{{tokenNumber, "1000"}, {tokenError, "unexpected token after num"}}},
		{"10 000", []Token{{tokenNumber, "10"}, {tokenError, "unexpected token after num"}}},
		{"01000", []Token{{tokenError, "expecting non zero digit, got '0'"}}},
		{"d6a", []Token{{tokenDice, "d"}, {tokenNumber, "6"}, {tokenError, "unexpected token after num"}}},
		{"0d6", []Token{{tokenError, "expecting non zero digit, got '0'"}}},
		{"5d6v4", []Token{{tokenNumber, "5"}, {tokenDice, "d"}, {tokenNumber, "6"}, {tokenError, "unexpected token after num"}}},
		{"5d6k4d", []Token{{tokenNumber, "5"}, {tokenDice, "d"}, {tokenNumber, "6"}, {tokenModifier, "k"}, {tokenNumber, "4"}, {tokenDice, "d"}, {tokenError, "expecting non zero digit, got '\\x00'"}}},
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
			fmt.Printf("Lexer Test %d OK: %s\n", i, resultTokens)
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
