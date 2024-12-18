package services

import (
	"errors"
	"gorp/core/domain/entities"
	"gorp/core/domain/interfaces"
)

type CharacterService struct {
	Repo interfaces.CharacterRepository
}

func (cs *CharacterService) CreateCharacter(name, className string) (entities.Character, error) {
	stats, exists := entities.ClassStats[className]
	if !exists {
		return entities.Character{}, errors.New("classe inconnue")
	}

	if len(name) < 3 {
		return entities.Character{}, errors.New("nom trop court")
	}
	if len(name) > 20 {
		return entities.Character{}, errors.New("nom trop long")
	}

	character := entities.Character{
		Name:  name,
		Class: className,
		Stats: stats,
	}

	err := cs.Repo.SaveCharacter(character)
	if err != nil {
		return entities.Character{}, errors.New("erreur lors de la sauvegarde du personnage")
	}
	return character, nil
}
