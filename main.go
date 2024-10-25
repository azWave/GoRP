package main

import (
    "fmt"
    "time"
    "gorp/game"
)

func main() {
    var startTime = time.Now()
    fmt.Println("Hello, the game is on!")
    player := game.NewPlayer(1, 1)
    world := game.GenerateWorld(5)

    for {
        game.Render(player, world)
        inputs := game.ReadInput()

        for _, input := range inputs {
            player.Move(input, world)
        }
        time.Sleep(100 * time.Millisecond)
    }
    fmt.Println("Time taken: ", time.Now().Sub(startTime))
}