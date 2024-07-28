package player

// creating the type for the player who will be playing in our magical arena

type Player struct {
	Name     string
	Health   int
	Strength int
	Attack   int
}

//  Function to create a new Player with the defined Type "Player" by pointing to it

func NewPlayer(name string, health, strength, attack int) *Player {
	return &Player{
		Name:     name,
		Health:   health,
		Strength: strength,
		Attack:   attack,
	}
}

/*  Two more functions are required one for checking if the player is alive or dead and other
to check how much the health of the player has reduced due to the damage taken */

func (p *Player) IsAlive() bool {
	return p.Health > 0 // if health is greater than 0, we can say that the player is alive
}

func (p *Player) DamageCheck(damage int) {
	p.Health -= damage // subtract the damage recieved from the current health of the player
	if p.Health < 0 {
		p.Health = 0 // if the health reaches below zero, set it to zero meaning the player is dead
	}
}
