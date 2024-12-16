package main

import (
	"fmt"
	"gorp/game"
	"gorp/game/console"
	"gorp/game/player"
	"gorp/game/world"
	"gorp/movements"
)

func main() {
	fmt.Print("Let's create your character. \n" +
		"What is your name?\n" +
		"> ")
	name := console.ReadInputWithConfirm()
	player := player.New(1, 1, name)
	world := world.GenerateWorld(5)

	for {
		game.Render(player, world)
		input := console.ReadInput()

		movements.Move(player, input, world)
	}

}
