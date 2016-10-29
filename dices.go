// Package dice provides tools to develop rpg games
package dice

//package main

import (
	"time"
	//    "crypto/rand"
	"fmt"
	"math/rand"
)

// Basic initialization
func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// Dice is a genrator of random integer numbers within a range
type Dice struct {
	sides int
}

// NewDice creates a ne Dice object with the given number of sides
func NewDice(sides int) *Dice {
	return &Dice{sides}
}

// Roll the dice and return the result
func (d *Dice) Roll() int {
	return rand.Intn(d.sides) + 1
}

// D2 returns a random integer between 1 and 2
func D2() int {
	return rand.Intn(2) + 1
}

// D4 returns a random integer between 1 and 4
func D4() int {
	return rand.Intn(4) + 1
}

// D6 returns a random integer between 1 and 6
func D6() int {
	return rand.Intn(6) + 1
}

// D8 returns a random integer between 1 and 8
func D8() int {
	return rand.Intn(8) + 1
}

// D10 returns a random integer between 1 and 10
func D10() int {
	return rand.Intn(10) + 1
}

// D12 returns a random integer between 1 and 12
func D12() int {
	return rand.Intn(12) + 1
}

// D20 returns a random integer between 1 and 20
func D20() int {
	return rand.Intn(20) + 1
}

// D30 returns a random integer between 1 and 30
func D30() int {
	return rand.Intn(30) + 1
}

// D100 returns a random integer between 1 and 100
func D100() int {
	return rand.Intn(100) + 1
}

// D200 returns a random integer between 1 and 200
func D200() int {
	return rand.Intn(200) + 1
}

// D1000 returns a random integer between 1 and 1000
func D1000() int {
	return rand.Intn(1000) + 1
}

func main() {
	// Seed the randon number generator...
	//r := rand.New(time.Now().UnixNano())
	fmt.Println("Dice: ", rand.Intn(6))
	d6 := NewDice(10)
	fmt.Println("Dice: ", d6.Roll())

	fmt.Println("D6: ", D6())
}
