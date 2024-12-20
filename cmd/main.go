package main

import (
	"fmt"
	"gorp/core/domain/entities"
	"gorp/core/services"
)

// func main() {
// 	repo := &output.FileRepository{BasePath: "tmp/data/"}
// 	characterService := &services.CharacterService{Repo: repo}
// 	printer := &output.FmtPrinter{}

// 	fmt.Println("Voulez-vous (1) charger un personnage existant ou (2) créer un nouveau personnage ?")
// 	var choice int
// 	fmt.Scanln(&choice)

//		switch choice {
//		case 1:
//			input.CharacterLoadHandler(printer, characterService)
//		case 2:
//			input.CharacterCreationHandler(printer, characterService)
//		default:
//			fmt.Println("Choix invalide.")
//			main()
//		}
//	}
func main() {
	mapService := services.NewMapService()
	gameMap := mapService.GenerateMap(10, 10)

	character := entities.Character{
		Name:  "Hero",
		Class: "Warrior",
		Stats: entities.Stats{
			Health: 100,
			Mana:   50,
			// ... other stats
		},
		Position: entities.Position{X: 0, Y: 0},
	}

	fmt.Println("Generated Map:")
	gameMap.Display(&character)

	fmt.Printf("Character %s is at position (%d, %d)\n", character.Name, character.Position.X, character.Position.Y)

	// Movement loop
	var command string
	for {
		fmt.Print("Enter command (z/q/s/d to move, quit to quit): ")
		fmt.Scan(&command)

		switch command {
		case "z":
			character.Move(0, -1, gameMap)
		case "q":
			character.Move(-1, 0, gameMap)
		case "s":
			character.Move(0, 1, gameMap)
		case "d":
			character.Move(1, 0, gameMap)
		case "quit":
			fmt.Println("Quitting the game.")
			return
		default:
			fmt.Println("Invalid command.")
		}

		fmt.Println("Updated Map:")
		gameMap.Display(&character)
		fmt.Printf("Character %s is at position (%d, %d)\n", character.Name, character.Position.X, character.Position.Y)
	}
}
