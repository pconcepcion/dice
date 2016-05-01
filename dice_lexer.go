// Dice lexer based on the talk "Lexical Scanning in Go" by Rob Pike [1]
// Runs the lexer as a gorutine that emits tokns on a channel for the
// parser to get them
// [1]:  https://www.youtube.com/watch?v=HxaD_trXwRE

package rpg

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Token type identifier
type tokenType int

// Generate the stringer for the tockentype
//go:generate stringer -type=tokenType

// Definition of token types
const (
	tokenUknown tokenType = iota
	tokenError
	tokenEOF
	// tokenSpace
	tokenNumber
	tokenDice
	tokenModifier
	/*tokenDot
	tokenComa
	tokenLeftParen
	tokenRightParen
	tokenLeftBrarce
	tokenRightBrace
	*/
)

// eof represnets the end of file
var eof = rune(0)

// digits represent the valid digits
const digits = "0123456789"
const nonZeroDigits = "123456789"

// Token represents a token of the lexer, and has a type and a value (a string)
type Token struct {
	typ tokenType
	val string
}

// Implementation of the Stringer interface taken from Rob Pike video
func (t Token) String() string {
	switch t.typ {
	case tokenEOF:
		return "EOF"
	case tokenError:
		return t.val
	}
	if len(t.val) > 10 {
		return fmt.Sprintf("%.10q...", t.val)
	}
	return fmt.Sprintf("%q", t.val)
}

// stateFn represents the state of the scanner as a function that returns the
// next state
type stateFn func(*lexer) stateFn

// lexer holds the state of the scanner
type lexer struct {
	input  string     // the string to tokeninze
	start  int        // starting position of the current token
	pos    int        // current position in the input
	width  int        // width of the last rune redaded
	tokens chan Token // channel of  scanned tokens
}

// run lexes the input by executing state functions until the state is nil
func (l *lexer) run() {
	for state := startState; state != nil; {
		state = state(l)
	}
	close(l.tokens) // Close the channel, no more tokens
}

// lex nitializes the lexer with an input string and get the reference of the
// lexer and the channel to receive tokens as a result
func lex(input string) (*lexer, chan Token) {
	l := &lexer{
		input:  input,
		tokens: make(chan Token),
		pos:    0,
		width:  0,
		start:  0,
	}
	go l.run()
	return l, l.tokens
}

// emit passes a Token back to the client trough the channel and updates
// the lexer start position to be used with the next Token
func (l *lexer) emit(t tokenType) {
	l.tokens <- Token{t, l.input[l.start:l.pos]}
	l.start = l.pos
}

// error emits token and terminates the scan
// by returing a nil pointer that will be the next
// state, terminating l.run.
func (l *lexer) errorf(format string, args ...interface{}) stateFn {
	l.tokens <- Token{
		tokenError,
		fmt.Sprintf(format, args...),
	}
	return nil
}

/// next returns the next rune in the input or eof
func (l *lexer) next() rune {
	var r rune
	if l.pos >= len(l.input) {
		l.width = 0
		return eof
	}
	r, l.width = utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += l.width
	return r
}

// ignore skips over the pending input before this point.
func (l *lexer) ignore() {
	l.start = l.pos
}

// peek returns but does not consume
// the next rune in the input.
func (l *lexer) peek() rune {
	rune := l.next()
	l.backup()
	return rune
}

// backup steps back one rune.
// Can be called only once per call of next.
func (l *lexer) backup() {
	l.pos -= l.width
}

// accept consumes the next rune
// if it's from the valid set.
func (l *lexer) accept(valid string) bool {
	if strings.IndexRune(valid, l.next()) >= 0 {
		return true
	}
	l.backup()
	return false
}

// acceptRun consumes a run of runes from the valid set.
func (l *lexer) acceptRun(valid string) {
	for strings.IndexRune(valid, l.next()) >= 0 {
	}
	l.backup()
}

/// State functions

// startState is the initial stateFn function, that looks for the first token and
// emits it
func startState(l *lexer) stateFn {
	switch r := l.next(); {
	case unicode.IsNumber(r):
		l.backup()
		return numberState
	case r == eof:
		l.emit(tokenEOF)
		return nil // finish the lexer
	case r == 'd':
		l.emit(tokenDice)
		return numberState
	}
	return l.errorf("unexpected token %v, expected either 'd' or number", l.next())
}

// numberState gets the nuber of dices and emits the token the next state should be diceState
func numberState(l *lexer) stateFn {
	if !l.accept(nonZeroDigits) {
		return l.errorf("expecting non zero digit, got %q", l.next())
	}
	l.acceptRun(digits)
	l.emit(tokenNumber)

	switch r := l.next(); {
	case r == 'd':
		l.emit(tokenDice)
		return numberState
	case strings.IndexRune("keors", r) != -1:
		l.backup()
		return modifierState
	case r == eof:
		l.emit(tokenEOF)
		return nil // finish the lexer
	}
	return l.errorf("unexpected token after num")
}

// modifierExplodingState handles the e and es tokens
func modifierExplodingState(l *lexer) stateFn {
	if l.accept("s") {
		// es
		l.emit(tokenModifier)
		return numberState
	}
	// e
	l.emit(tokenModifier)
	return startState
}

// modifierState extracts one of the valid modifiers:
// * k = Keep
// * e = Explode
// * es = Exploding success
// * o = Open
// * s = Success
// * r = Reroll
func modifierState(l *lexer) stateFn {
	if l.accept("e") {
		return modifierExplodingState
	}
	if l.accept("o") {
		l.emit(tokenModifier)
		return startState
	}
	// Keep, Rerroll and Success need a number afterwards
	if l.accept("krs") {
		l.emit(tokenModifier)
		return numberState
	}
	return l.errorf("unexpected modifier token")
}

// Character classes

// isWhitespace returns true when theceived rune is a whitespace
func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}
