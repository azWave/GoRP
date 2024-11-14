package main

import (
	"fmt"
	"gorp/game"
)

func main() {
	fmt.Println("Hello, the game is on!")
	player := game.NewPlayer(1, 1)
	world := game.GenerateWorld(5)

	for {
		game.Render(player, world)
		input := game.ReadInput()

		player.Move(input, world)
	}

}
