package game

import (
    "fmt"
)

var errorMessages []string

func Render(player *Player, world *World) {
    // Clear the screen if there are no new error messages
    if len(errorMessages) == 0 {
        fmt.Print("\033[H\033[2J") // Clear the screen
    }

    // Render the world and player position
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

    // Print and clear error messages after displaying them
    for _, msg := range errorMessages {
        fmt.Println(msg)
    }
    errorMessages = []string{} // Clear error messages after displaying
}

// AddErrorMessage adds an error message to be displayed
func AddErrorMessage(msg string) {
    if len(msg) > 0 && msg[:7] == "[ERROR]" {
        errorMessages = append(errorMessages, msg)
    }
}