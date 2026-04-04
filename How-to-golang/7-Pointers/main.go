package main

import "fmt"

type Player struct {
	health int
}

func (p *Player) takeDamageFromPlayer(dmg int) {
	fmt.Println("player is taking damage from explosion")
	p.health -= dmg
}

func takeDamageFromPlayer(player *Player, dmg int) {
	fmt.Println("player is taking damage from explosion")
	player.health -= dmg
}

func main() {
	fmt.Println("introduction to pointers")
	player := &Player{health: 100} //8 byte long integer
	fmt.Printf("before explosion %+v \n", player)
	takeDamageFromPlayer(player, 50)
	fmt.Printf("after  explosion %+v \n", player)
	player.takeDamageFromPlayer(50)
	fmt.Println("after second explosion", player)

}
