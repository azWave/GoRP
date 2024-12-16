package output

import (
	"encoding/json"
	"gorp/core/domain/entities"
	"os"
)

type FileRepository struct {
	FilePath string
}

func (fr *FileRepository) Save(character entities.Character) error {
	file, err := os.OpenFile(fr.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := json.Marshal(character)
	if err != nil {
		return err
	}
	_, err = file.Write(append(data, '\n'))
	return err
}

func (fr *FileRepository) FindByName(name string) (entities.Character, error) {
	// Implémentation à venir
	return entities.Character{}, nil
}
