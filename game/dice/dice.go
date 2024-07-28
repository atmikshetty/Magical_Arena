package dice // everything regarding the usage of dice in our game will be stored in this package

// Necessary imports for perfroming the various operations
import (
	"math/rand"
	"time"
)

//  Initializing a  dice, sides is an integer since the sides of a dice are always in numbers

type Dice struct {
	sides int
}

// Function to created a six-sided dice, this is of type Dice as it is a pointer to it

func NewDice() *Dice {
	return &Dice{sides: 6}
}

// function implementing the rolling of the dice

func (d *Dice) Rolling() int { // d is the pointer to the dice type
	rand.Seed(time.Now().UnixNano()) // we are getting the current time in NanoSeconds using this and Seed makes sure we get a different number everytime
	return rand.Intn(d.sides) + 1    // gives a number between 1 and 6
}
