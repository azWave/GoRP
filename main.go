package main

import (
	"gorp/game"
	"gorp/game/console"
	"gorp/game/player"
	"gorp/game/world"
)

func main() {
	player := player.CreatePlayer(1, 1)
	world := world.GenerateWorld(5)

	for {
		game.Render(player, world)
		input := console.ReadInput()

		player.Move(input, world)
	}

}
