package input

import (
	"fmt"
	"gorp/core/domain/entities"
	"gorp/core/services"
)

func CharacterCreationHandler(service *services.CharacterService) entities.Character {
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
	PrintCharacterSummary(character)

	return character
}

func PrintCharacterSummary(character entities.Character) {
	fmt.Printf(
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
