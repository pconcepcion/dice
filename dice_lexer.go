// Dice lexer based on the talk "Lexical Scanning in Go" by Rob Pike [1]
// Runs the lexer as a gorutine that emits tokns on a channel for the
// parser to get them
// [1]:  https://www.youtube.com/watch?v=HxaD_trXwRE

package rpg

import (
	"fmt"
	"unicode/utf8"
)

// Token type identifier
type tokenType int

// Definition of token types
const (
	tokenError tokenType = iota
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
	widht  int        // widht of the last rune redaded
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

/// next returns the next rune in the input or eof
func (l *lexer) next() (rune, int) {
	var r rune
	if l.pos >= len(l.input) {
		l.widht = 0
		return eof, 0
	}
	r, l.widht = utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += l.widht
	return r, l.widht
}

// Character classes

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

// ignore skips over the pending input before this point.
func (l *lexer) ignore() {
	l.start = l.pos
}

// peek returns but does not consume
// the next rune in the input.
func (l *lexer) peek() int {
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
		return numDicesState
	case r == eof:
		return l.errorf("empty action")
	case r == 'd':
		return diceState
	}
	return l.errorf("unexpected token %s, expected either 'd' or number", l.next())
}

// numDicesState gets the nuber of dices and emits the token the next state should be diceState
func numbmerState(l *lexer) stateFn {
	digits := "0123456789"
	// If starts with 0 must be a 0 and and the next can't be a number
	if l.accept("0") && strings.IndexRune(digits, l.peek()) {
		return l.errorf("invalid sequence 0%s", l.peek())
	}
	l.acceptRun(digits)
	l.emit(tokenNumber)

	switch r := l.next(); {
	case r == "d":
		return diceState
	case strings.IndexRune("keors", r):
		return modifierState
	case r == eof:
		l.emit(tokenEOF)
		return nil // finish the lexer
	}
	return l.errrof("unexpected token after num")
}

// modifierState extracts one of the valid modifiers:
// * k = Kepp
// * e = Explode
// * es = Exploding success
// * o = Open
// * s = Success
// * r = Reroll
func modifierSate(l *lexer) {
	if l.accept("e") {
		if l.accept("s") {
			// es
			l.emmit(modifierToken)
			return numberState
		}
		// e
		l.emmit(modifierToken)
		return startState
	}
	if l.accept("o") {
		l.emmit(modifierToken)
		return startState
	}
	// Keep, Rerroll and Success need a number afterwards
	if l.accept("krs") {
		l.emmit(modifierToken)
		return numberState
	}
	return l.errrof("unexpected modifier token")
}

// diceState accepts the "d" marking the dice token and emits the token, shuld be numberState
func diceState(l *lexer) stateFn {
	if l.accept("d") {
		l.emit(tokenDice)
		return numberStated
	}
	return l.errrof("expected dice token, got %s", l.peek())
}

// lexText scanns regular text until a different token is found
func lexText(l *lexer) stateFn {

	return nil
}
