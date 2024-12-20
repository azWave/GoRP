package services

import (
	"errors"
	"gorp/core/domain/constants"
	"gorp/core/domain/entities"
	"gorp/core/domain/interfaces"
)

type CharacterService struct {
	Character *entities.Character
	Repo      interfaces.CharacterRepository
}

func (cs *CharacterService) CreateCharacter(name, className string) (entities.Character, error) {
	stats, exists := entities.ClassStats[className]
	if !exists {
		return entities.Character{}, errors.New("classe inconnue")
	}

	if len(name) < constants.MinCharacterNameLength {
		return entities.Character{}, errors.New("nom trop court")
	}
	if len(name) > constants.MaxCharacterNameLength {
		return entities.Character{}, errors.New("nom trop long")
	}

	character := entities.Character{
		Name:  name,
		Class: className,
		Stats: stats,
		Position: entities.Position{
			X:       0,
			Y:       0,
			MapName: constants.FirstWorldMapName,
		},
	}

	err := cs.SaveCharacter(character)
	if err != nil {
		return entities.Character{}, err
	}
	cs.Character = &character
	return character, nil
}

func (cs *CharacterService) SaveCharacter(character entities.Character) error {
	err := cs.Repo.SaveCharacter(character)
	if err != nil {
		return errors.New("erreur lors de la sauvegarde du personnage")
	}
	return nil
}

func (cs *CharacterService) LoadCharacter(name string) (entities.Character, error) {
	character, err := cs.Repo.LoadCharacter(name)
	if err != nil {
		return entities.Character{}, errors.New("personnage introuvable")
	}
	cs.Character = &character
	return character, nil
}

func (cs *CharacterService) SetMap(character entities.Character, mapName string) error {
	character.Position.MapName = mapName
	return cs.SaveCharacter(character)
}
