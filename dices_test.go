// Package rpg provides tools to develop rpg games
package rpg

import (
	"fmt"
	"math"
	"testing"
)

/*func TestNewDice(sides int) *Dice {
	return &Dice{sides}
}
func Test(d *Dice) Roll() int {
	return rand.Intn(d.sides)
}
*/

const ACCEPTABLE_ERROR = 0.1
const num_iterations = 20 * 1000 * 1000

// Test that D2() returns a random integer between 1 and 2
func TestD2(t *testing.T) {
	var v int
	v = D2()
	if v < 1 || v > 2 {
		t.Error("Expected value between 1 and 2", v)
	}
}

// Test that D4() returns a random integer between 1 and 4
func TestD4(t *testing.T) {
	var v int
	v = D4()
	if v < 1 || v > 4 {
		t.Error("Expected value between 1 and 4", v)
	}
}

// Test that D6() returns a random integer between 1 and 6
func TestD6(t *testing.T) {
	var v int
	v = D6()
	if v < 1 || v > 6 {
		t.Error("Expected value between 1 and 6", v)
	}
}

// Test that D8() returns a random integer between 1 and 8
func TestD8(t *testing.T) {
	var v int
	v = D8()
	if v < 1 || v > 8 {
		t.Error("Expected value between 1 and 8", v)
	}
}

// Test that D10() returns a random integer between 1 and 10
func TestD10(t *testing.T) {
	var v int
	v = D10()
	if v < 1 || v > 10 {
		t.Error("Expected value between 1 and 10", v)
	}
}

// Test that D12() returns a random integer between 1 and 12
func TestD12(t *testing.T) {
	var v int
	v = D12()
	if v < 1 || v > 12 {
		t.Error("Expected value between 1 and 12", v)
	}
}

// Test that D20() returns a random integer between 1 and 20
func TestD20(t *testing.T) {
	var v int
	v = D20()
	if v < 1 || v > 20 {
		t.Error("Expected value between 1 and 20", v)
	}
}

// Test that D30() returns a random integer between 1 and 30
func TestD30(t *testing.T) {
	var v int
	v = D30()
	if v < 1 || v > 30 {
		t.Error("Expected value between 1 and 30", v)
	}
}

// Test that D100() returns a random integer between 1 and 100
func TestD100(t *testing.T) {
	var v int
	v = D100()
	if v < 1 || v > 100 {
		t.Error("Expected value between 1 and 100", v)
	}
}

// Test that D200() returns a random integer between 1 and 200
func TestD200(t *testing.T) {
	var v int
	v = D200()
	if v < 1 || v > 200 {
		t.Error("Expected value between 1 and 200", v)
	}
}

// Test that D1000() returns a random integer between 1 and 1000
func TestD1000(t *testing.T) {
	var v int
	v = D1000()
	if v < 1 || v > 1000 {
		t.Error("Expected value between 1 and 1000", v)
	}
}

// Test that average of num_iterations D2() is arround 1.5
func TestAverageD2(t *testing.T) {
	var sum int
	sum = 0
	for i := 0; i < num_iterations; i++ {
		sum += D2()
	}
	var average float64
	average = float64(sum) / num_iterations
	fmt.Println("Average D2: ", average)
	if math.Abs(average-1.5) > ACCEPTABLE_ERROR {
		t.Error("Expected value between 1.4 and 1.6", average)
	}
}

// Test that average of num_iterations D4() is arround 2.5
func TestAverageD4(t *testing.T) {
	var sum int
	sum = 0
	for i := 0; i < num_iterations; i++ {
		sum += D4()
	}
	var average float64
	average = float64(sum) / num_iterations
	fmt.Println("Average D4: ", average)
	if math.Abs(average-2.5) > ACCEPTABLE_ERROR {
		t.Error("Expected value between 2.4 and 2.6", average)
	}
}

// Test that average of num_iterations D6() is arround 3.5
func TestAverageD6(t *testing.T) {
	var sum int
	sum = 0
	for i := 0; i < num_iterations; i++ {
		sum += D6()
	}
	var average float64
	average = float64(sum) / num_iterations
	fmt.Println("Average D6: ", average)
	if math.Abs(average-3.5) > ACCEPTABLE_ERROR {
		t.Error("Expected value between 3.4 and 3.6", average)
	}
}

// Test that average of num_iterations D8() is arround 4.5
func TestAverageD8(t *testing.T) {
	var sum int
	sum = 0
	for i := 0; i < num_iterations; i++ {
		sum += D8()
	}
	var average float64
	average = float64(sum) / num_iterations
	fmt.Println("Average D8: ", average)
	if math.Abs(average-4.5) > ACCEPTABLE_ERROR {
		t.Error("Expected value between 4.4 and 4.6", average)
	}
}

// Test that average of num_iterations D10() is arround 5.5
func TestAverageD10(t *testing.T) {
	var sum int
	sum = 0
	for i := 0; i < num_iterations; i++ {
		sum += D10()
	}
	var average float64
	average = float64(sum) / num_iterations
	fmt.Println("Average D10: ", average)
	if math.Abs(average-5.5) > ACCEPTABLE_ERROR {
		t.Error("Expected value between 5.4 and 5.6", average)
	}
}

// Test that average of num_iterations D12() is arround 6.5
func TestAverageD12(t *testing.T) {
	var sum int
	sum = 0
	for i := 0; i < num_iterations; i++ {
		sum += D12()
	}
	var average float64
	average = float64(sum) / num_iterations
	fmt.Println("Average D12: ", average)
	if math.Abs(average-6.5) > ACCEPTABLE_ERROR {
		t.Error("Expected value between 6.4 and 6.6", average)
	}
}

// Test that average of num_iterations D20() is arround 10.5
func TestAverageD20(t *testing.T) {
	var sum int
	sum = 0
	for i := 0; i < num_iterations; i++ {
		sum += D20()
	}
	var average float64
	average = float64(sum) / num_iterations
	fmt.Println("Average D20: ", average)
	if math.Abs(average-10.5) > ACCEPTABLE_ERROR {
		t.Error("Expected value between 10.4 and 10.6", average)
	}
}

// Test that average of num_iterations D30() is arround 15.5
func TestAverageD30(t *testing.T) {
	var sum int
	sum = 0
	for i := 0; i < num_iterations; i++ {
		sum += D30()
	}
	var average float64
	average = float64(sum) / num_iterations
	fmt.Println("Average D30: ", average)
	if math.Abs(average-15.5) > ACCEPTABLE_ERROR {
		t.Error("Expected value between 15.4 and 15.6", average)
	}
}

// Test that average of num_iterations D100() is arround 50.5
func TestAverageD100(t *testing.T) {
	var sum int
	sum = 0
	for i := 0; i < num_iterations; i++ {
		sum += D100()
	}
	var average float64
	average = float64(sum) / num_iterations
	fmt.Println("Average D100: ", average)
	if math.Abs(average-50.5) > ACCEPTABLE_ERROR {
		t.Error("Expected value between 50.4 and 50.6", average)
	}
}

// Test that average of num_iterations for D200() is arround 100.5
func TestAverageD200(t *testing.T) {
	var sum int
	sum = 0
	for i := 0; i < num_iterations; i++ {
		sum += D200()
	}
	var average float64
	average = float64(sum) / num_iterations
	fmt.Println("Average D200: ", average)
	if math.Abs(average-100.5) > ACCEPTABLE_ERROR {
		t.Error("Expected value between 100.4 and 100.6", average)
	}
}

// Test that average of Average for num_iterations D1000() is arround 500.5
func TestAverageD1000(t *testing.T) {
	var sum int
	sum = 0
	for i := 0; i < num_iterations; i++ {
		sum += D1000()
	}
	var average float64
	average = float64(sum) / num_iterations
	fmt.Println("Average D1000: ", average)
	if math.Abs(average-500.5) > ACCEPTABLE_ERROR {
		t.Error("Expected value between 500.4 and 500.6", average)
	}
}

// Test a dice with random sideso
func TestRandomDice(t *testing.T) {
	// Create a random dice
	var sides int
	sides = D100()
	var expectedAverage float64
	expectedAverage = float64(sides+1) / 2.0
	d := NewDice(sides)

	// Test it
	var sum int
	for i := 0; i < num_iterations; i++ {
		sum += d.Roll()
	}
	var average float64
	average = float64(sum) / num_iterations
	fmt.Printf("Average D%d: %g\n", sides, average)
	if math.Abs(average-expectedAverage) > ACCEPTABLE_ERROR {
		t.Error("Expected value for D", sides, " was around ", expectedAverage, " with an ", ACCEPTABLE_ERROR, " error")

	}
}
