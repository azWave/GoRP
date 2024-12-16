package interfaces

import "gorp/core/domain/entities"

type CharacterRepository interface {
	SaveCharacter(character entities.Character) error
	LoadCharacter(name string) (entities.Character, error)
}
