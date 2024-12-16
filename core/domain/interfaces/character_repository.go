package interfaces

import "gorp/core/domain/entities"

type CharacterRepository interface {
    Save(character entities.Character) error
    FindByName(name string) (entities.Character, error)
}
