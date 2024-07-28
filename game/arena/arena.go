package arena

import (
	"fmt"
	"game/dice"
	"game/player"
)

// type structure for Arena where the players will play each other

type Arena struct {
	Dice *dice.Dice // pointer to the dice struct from the dice.go
}

// Creating a new arena

func NewArena() *Arena {
	return &Arena{
		Dice: dice.NewDice(), // creating a six sided dice by calling it from the dice.go file
	}
}

// Function for battle/fighting between the two players

func (a *Arena) Fight(p1, p2 *player.Player) *player.Player { // a a pointer to the Arena type, indicating that the function is also of Arena type, both p1 and p2 are instances of player.Player type
	attacker, defender := a.determineOrder(p1, p2) // determining the order of attack

	for attacker.IsAlive() && defender.IsAlive() { // if both are alive then only they can attack each other
		a.performAttack(attacker, defender)
		attacker, defender = defender, attacker // swapping attack and defense positions
	}

	if p1.IsAlive() { // if p1 is alive, p1 wins
		return p1
	}
	return p2 // p2 wins
}

// Choosing which player will attack first

func (a *Arena) determineOrder(p1, p2 *player.Player) (*player.Player, *player.Player) {
	if p1.Health <= p2.Health { // lower health will attack first
		return p1, p2
	}

	return p2, p1
}

// Function for attacking between the players and showing the recieved damage

func (a *Arena) performAttack(attacker, defender *player.Player) {
	attackChance := a.Dice.Rolling() // calling the rolling function to determine the number of the dice
	defendChance := a.Dice.Rolling()

	attackDamage := attacker.Attack * attackChance // initial value + the dice value
	defensePower := defender.Strength * defendChance

	overallDamageTaken := attackDamage - defensePower

	if overallDamageTaken > 0 {
		defender.DamageCheck(overallDamageTaken)
	}

	// Printing all the changes and outputs recieved in this function

	fmt.Printf("%s is attacking %s\n", &attacker.Name, &defender.Name)
	fmt.Printf("Attackers Dice Roll: %d, Defenders Dice Roll: %d\n", attackChance, defendChance)
	fmt.Printf("Attack Damage: %d , Defense Power: %d\n", attackDamage, defensePower)
	fmt.Printf("Damage taken by %s is %d\n", &defender.Name, overallDamageTaken)
	fmt.Printf("Final Stats: %s health: %d, %s health: %d\n\n", attacker.Name, attacker.Health, defender.Name, defender.Health)
}
