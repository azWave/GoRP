package game

import (
	"fmt"
	"gorp/game/player"
	"gorp/game/world"
)

var errorMessages []string

func Render(player *player.Player, world *world.World) {
	// Clear the screen
	fmt.Print("\033[H\033[2J")

	for y, row := range world.Map {
		for x, cell := range row {
			if player.X == x && player.Y == y {
				fmt.Print("P") // Display player
			} else {
				fmt.Print(string(cell))
			}
		}
		fmt.Println()
	}

	// Print error messages if any
	for _, msg := range errorMessages {
		fmt.Println(msg)
	}
}