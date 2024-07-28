package player

import (
	"testing"
)

func TestNewPlayer(t *testing.T) { // the parameter is provided by GO to see the logs and errors during the testing
	p := NewPlayer("Test", 100, 10, 5)                                            // providing mock data for testing
	if p.Name != "Test" || p.Health != 100 || p.Strength != 10 || p.Attack != 5 { // checking if the data is correcty taken by the instance of the new Player
		t.Errorf("Player not created correctly")
	}
}

func TestIsAlive(t *testing.T) {
	p := NewPlayer("Test", 100, 10, 5)
	if !p.IsAlive() { // if player is alive and the function still shows dead, there is an error
		t.Errorf("Player should be alive")
	}

	p.Health = 0
	if p.IsAlive() { // if the health is zero, the player is dead
		t.Errorf("Player should be dead")
	}
}

func TestDamageCheck(t *testing.T) {
	p := NewPlayer("Test", 100, 10, 5) // initialize a player for testing
	p.DamageCheck(30)                  // damage is given as 30
	if p.Health != 70 {                // health is 100 and since damage is subtracted the value should be 70
		t.Errorf("Expected health 70, got %d", p.Health) // if the return value is not 70 there is an error
	}

	p.DamageCheck(80) // damage is given as 80, since our current health is 70, this should directly set the health to 0
	if p.Health != 0 {
		t.Errorf("Health should not go below 0") // Error if the health is not zero
	}
}
