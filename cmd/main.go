package main

import (
	"fmt"
	"gorp/core/services"
	input_character "gorp/infra/input/character"
	input_map "gorp/infra/input/map"
	"gorp/infra/output"
)

func main() {
	repo := &output.FileRepository{BasePath: "tmp/data/"}
	mapService := &services.MapService{MapRepo: repo}
	characterService := &services.CharacterService{Repo: repo}
	printer := &output.FmtPrinter{}

Start:

	fmt.Println("Voulez-vous (1) charger un personnage existant ou (2) créer un nouveau personnage ?")
	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		input_character.CharacterLoadHandler(printer, characterService, mapService)
	case 2:
		input_character.CharacterCreationHandler(printer, characterService)
	default:
		fmt.Println("Choix invalide.")
		goto Start
	}

	if mapService.GameMap == nil {
		input_map.MapCreationHandler(printer, mapService)
	}
	input_map.PrintMap(printer, mapService.GameMap)
}

// func main() {
// 	mapService := services.MapService{}
// 	gameMap := mapService.GenerateMap(10, 10)

// 	character := entities.Character{
// 		Name:  "Hero",
// 		Class: "Warrior",
// 		Stats: entities.Stats{
// 			Health: 100,
// 			Mana:   50,
// 			// ... other stats
// 		},
// 		Position: entities.Position{X: 3, Y: 3},
// 	}

// 	fmt.Println("Generated Map:")
// 	printer := &output.FmtPrinter{}
// 	input.PrintMap(printer, gameMap, &character)
// }
