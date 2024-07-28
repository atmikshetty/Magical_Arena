package dice

import (
	"testing"
)

func TestNewDice(t *testing.T) { // the parameter is provided by GO to see the logs and errors during the testing
	d := NewDice() // initializes a new Dice Instance

	if d.sides != 6 {
		t.Errorf("Expected 6 sides, got %d", d.sides) // if sides are not eqaul to 6, error will be displayed
	}
}

func TestRolling(t *testing.T) {
	d := NewDice() // initializes a new Dice Instance

	for i := 0; i < 1000; i++ { // performs multiple dice rolls
		roll := d.Rolling() // gets a dice value by calling the rolling function
		if roll < 1 || roll > 6 {
			t.Errorf("Dice Rolled is %d out of range 1-6", roll)
		}
	}
}
