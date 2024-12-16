package input

import (
	"fmt"
	"gorp/core/services"
)

func CharacterCreationHandler(service *services.CharacterService) {
	var name, class string

	fmt.Println("Entrez le nom du personnage :")
	fmt.Scanln(&name)

	fmt.Println("Choisissez une classe (Guerrier, Mage, Voleur) :")
	fmt.Scanln(&class)

	character, err := service.CreateCharacter(name, class)
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	fmt.Printf("Personnage créé avec succès :\n%+v\n", character)
}
