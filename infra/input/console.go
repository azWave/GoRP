package input

import (
	"fmt"
	"gorp/core/domain/entities"
	"gorp/core/domain/interfaces"
	"gorp/core/services"
)

func CharacterLoadHandler(printer interfaces.Printer, service *services.CharacterService) entities.Character {
	var name string

	fmt.Println("-=-=-=-=-=-=-=-=-=-=-=-=-=-=-")
	fmt.Println("Chargement d'un personnage")
	fmt.Println("-=-=-=-=-=-=-=-=-=-=-=-=-=-=-")
	fmt.Println("Entrez le nom du personnage :")
	fmt.Scanln(&name)

	character, err := service.LoadCharacter(name)
	if err != nil {
		fmt.Println("Erreur :", err)
		return entities.Character{}
	}

	fmt.Println("\n-=-=-=-=-=-=-=-=-=-=-=-=-=-=-")
	fmt.Println("Personnage chargé avec succès.")
	fmt.Println("-=-=-=-=-=-=-=-=-=-=-=-=-=-=-")
	PrintCharacterSummary(printer, character)

	return character
}

func CharacterCreationHandler(printer interfaces.Printer, service *services.CharacterService) entities.Character {
	var name, class string

	fmt.Println("-=-=-=-=-=-=-=-=-=-=-=-=-=-=-")
	fmt.Println("Création d'un personnage")
	fmt.Println("-=-=-=-=-=-=-=-=-=-=-=-=-=-=-")
	fmt.Println("Entrez le nom du personnage :")
	fmt.Scanln(&name)

	fmt.Println("\nChoisissez une classe (Guerrier, Mage, Voleur) :")
	fmt.Scanln(&class)

	character, err := service.CreateCharacter(name, class)
	if err != nil {
		fmt.Println("Erreur :", err)
		return entities.Character{}
	}

	fmt.Println("\n-=-=-=-=-=-=-=-=-=-=-=-=-=-=-")
	fmt.Println("Personnage créé avec succès.")
	fmt.Println("-=-=-=-=-=-=-=-=-=-=-=-=-=-=-")
	PrintCharacterSummary(printer, character)

	return character
}

func PrintCharacterSummary(printer interfaces.Printer, character entities.Character) {
	printer.Printf(
		"Nom : %v\n"+
			"Classe : %v\n"+
			"PV : %v\n"+
			"PM : %v\n"+
			"Force : %v\n"+
			"Intelligence : %v\n"+
			"Défense : %v\n"+
			"Résistance Magique : %v\n"+
			"Agilité : %v\n"+
			"Chance : %v\n"+
			"Endurance : %v\n"+
			"Esprit : %v\n",
		character.Name,
		character.Class,
		character.Stats.Health,
		character.Stats.Mana,
		character.Stats.Strength,
		character.Stats.Intelligence,
		character.Stats.Defense,
		character.Stats.MagicResist,
		character.Stats.Agility,
		character.Stats.Luck,
		character.Stats.Endurance,
		character.Stats.Spirit,
	)
}

func PrintMap(printer interfaces.Printer, gameMap *entities.Map, character *entities.Character) {
	for y := 0; y < gameMap.Height; y++ {
		for x := 0; x < gameMap.With; x++ {
			if character.Position.X == x && character.Position.Y == y {
				printer.Printf("$")
			} else {
				cell := gameMap.Cells[y][x]
				printer.Printf("%s", entities.ClassCellProperties[cell.Type].Display)
			}
			printer.Printf(" ")
		}
		printer.Println()
	}
}
