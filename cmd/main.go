package main

import (
	"fmt"
	"gorp/core/services"
	"gorp/infra/input"
	"gorp/infra/output"
)

func main() {
	repo := &output.FileRepository{BasePath: "tmp/data/"}
	characterService := &services.CharacterService{Repo: repo}
	printer := &output.FmtPrinter{}

	fmt.Println("Voulez-vous (1) charger un personnage existant ou (2) créer un nouveau personnage ?")
	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		input.CharacterLoadHandler(printer, characterService)
	case 2:
		input.CharacterCreationHandler(printer, characterService)
	default:
		fmt.Println("Choix invalide.")
		main()
	}
}
