package player

import (
	"fmt"
	"gorp/game/console"
	"gorp/game/world"
)

type Player struct {
	X, Y int
	Name string
}

func CreatePlayer(x, y int) *Player {
	fmt.Print("Let's create your character. \n" +
		"What is your name?\n" +
		"> ")
	name := console.ReadInputWithConfirm()
	return &Player{x, y, name}
}

func (p *Player) Move(direction string, world *world.World) {
	switch direction {
	case "s":
		if world.IsWalkable(p.X, p.Y+1) {
			p.Y++
		}
	case "z":
		if world.IsWalkable(p.X, p.Y-1) {
			p.Y--
		}
	case "q":
		if world.IsWalkable(p.X-1, p.Y) {
			p.X--
		}
	case "d":
		if world.IsWalkable(p.X+1, p.Y) {
			p.X++
		}
	}
}
