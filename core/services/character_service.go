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

	character := entities.Character{
		Name:  name,
		Class: className,
		Stats: stats,
	}

	err := cs.Repo.Save(character)
	if err != nil {
		return entities.Character{}, errors.New("erreur lors de la sauvegarde du personnage")
	}
	return character, nil
}
