package main

import (
	"gorp/game"
)

func main() {
	player := game.CreatePlayer(1, 1)
	world := game.GenerateWorld(5)

	for {
		game.Render(player, world)
		input := game.ReadInput()

		player.Move(input, world)
	}

}
