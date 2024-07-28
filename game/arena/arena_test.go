package arena

import (
	"game/player"
	"testing"
)

func TestNewArena(t *testing.T) {
	a := NewArena() // Initializing a new arena
	if a.Dice == nil {
		t.Errorf("Arena created without a dice") // always one value of dice will be available, hence nil value is not possible
	}
}

func TestDetermineOrder(t *testing.T) {
	a := NewArena()
	// Creating Mock players for testing
	p1 := player.NewPlayer("P1", 50, 10, 5)
	p2 := player.NewPlayer("P2", 100, 10, 5)

	attacker, defender := a.DetermineOrder(p1, p2) // as per the DetermineOrder function, p1,p2 should be initialized to attacker and defender
	if attacker != p1 || defender != p2 {
		t.Errorf("Order determination failed")
	}

	attacker, defender = a.DetermineOrder(p2, p1) // as per the DetermineOrder function, p1,p2 should be initialized to defender and attacker
	if attacker != p1 || defender != p2 {
		t.Errorf("Order determination failed")
	}
}

func TestFight(t *testing.T) {
	a := NewArena()
	// Creating Mock players for testing
	p1 := player.NewPlayer("P1", 50, 10, 5)
	p2 := player.NewPlayer("P2", 100, 10, 5)

	winner := a.Fight(p1, p2) // Calling the fighting function between two players
	if !winner.IsAlive() {    // winner should always be alive, there will be no case where winner is dead
		t.Errorf("Winner should be alive")
	}
	if p1.IsAlive() && p2.IsAlive() { // At the end, only one player is designed to win in the arena, hence two players cannot be alive
		t.Errorf("Both players cannot be alive after a fight")
	}
}
