package main

import (
	"fmt"
	"game/arena"
	"game/player"
)

/* Here we will use all the functions created in other three folders:
1. Firstly we will initialize the players using the NewPlayer method from player.go
2. Next, we initialize the arena from the arena.go function NewArena
3. Finally, we can call the fight function to determine the winner of this fight
*/

func main() {

	// Creating the players

	playerA := player.NewPlayer("Player A", 50, 5, 10)
	playerB := player.NewPlayer("Player B", 100, 10, 5)

	// Initializing the Arena
	magicalArena := arena.NewArena()

	// Starting the fight between the players
	fmt.Println("Let the fight begin!!!")
	fmt.Printf("%s (Health: %d, Strength: %d, Attack: %d) vs %s (Health: %d, Strength: %d, Attack: %d)\n\n",
		playerA.Name, playerA.Health, playerA.Strength, playerA.Attack,
		playerB.Name, playerB.Health, playerB.Strength, playerB.Attack)

	// Figuring out the winner after the fight

	winner := magicalArena.Fight(playerA, playerB)

	fmt.Printf("\nThe winner is %s with %d health remaining!\n", winner.Name, winner.Health)
}
