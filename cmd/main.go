package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/pconcepcion/dice"
	log "github.com/sirupsen/logrus"
)

// TODO: do proper validation
var validCharacters = regexp.MustCompile(`^[a-z0-9]+$`)

func main() {
	// Output to stderr instead of stdout, could also be a file.
	log.SetOutput(os.Stderr)
	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)
	if len(os.Args) != 2 {
		// We need an extra argument
		fmt.Println("Usage: dice <expression>")
		os.Exit(-1)

	}
	expression := os.Args[1]
	if !validCharacters.MatchString(expression) {
		fmt.Printf("Error: Invalid expresssion: #%s#\n", expression)
		fmt.Println("Usage: dice <expression>")
		os.Exit(-2)
	}
	sde := dice.NewSimpleExpression(expression)
	res, err := sde.Roll()
	if err != nil {
		fmt.Printf("Error rolling the dices %v\n", err)
		os.Exit(-3)
	}
	fmt.Printf("result: %v\n", res)
}
