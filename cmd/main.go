package main

import (
	"fmt"
	"gorp/core/domain/entities"
	"gorp/core/services"
	"gorp/infra/input"
	"gorp/infra/output"
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
	mapService := services.MapService{}
	gameMap := mapService.GenerateMap(10, 10)

	character := entities.Character{
		Name:  "Hero",
		Class: "Warrior",
		Stats: entities.Stats{
			Health: 100,
			Mana:   50,
			// ... other stats
		},
		Position: entities.Position{X: 3, Y: 3},
	}

	fmt.Println("Generated Map:")
	printer := &output.FmtPrinter{}
	input.PrintMap(printer, gameMap, &character)
}
