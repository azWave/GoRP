package output

import (
	"encoding/json"
	"fmt"
	"gorp/core/domain/entities"
	"os"
	"path/filepath"
)

type FileRepository struct {
	BasePath string
}

func NewFileRepository(basePath string) *FileRepository {
	return &FileRepository{BasePath: basePath}
}

func (fr *FileRepository) ensureDirectoryExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.MkdirAll(path, 0755)
	}
	return nil
}

func (fr *FileRepository) SaveCharacter(character entities.Character) error {
	charDir := filepath.Join(fr.BasePath, "characters")
	if err := fr.ensureDirectoryExists(charDir); err != nil {
		return fmt.Errorf("erreur création dossier personnages : %w", err)
	}

	charFile := filepath.Join(charDir, fmt.Sprintf("%s.json", character.Name))
	file, err := os.Create(charFile)
	if err != nil {
		return fmt.Errorf("erreur création fichier personnage : %w", err)
	}
	defer file.Close()

	data, err := json.MarshalIndent(character, "", "  ")
	if err != nil {
		return fmt.Errorf("erreur encodage JSON : %w", err)
	}

	_, err = file.Write(data)
	return err
}

func (fr *FileRepository) LoadCharacter(name string) (entities.Character, error) {
	charFile := filepath.Join(fr.BasePath, "characters", fmt.Sprintf("%s.json", name))
	file, err := os.Open(charFile)
	if err != nil {
		return entities.Character{}, fmt.Errorf("erreur ouverture fichier personnage : %w", err)
	}
	defer file.Close()

	var character entities.Character
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&character); err != nil {
		return entities.Character{}, fmt.Errorf("erreur décodage JSON : %w", err)
	}

	return character, nil
}

func (fr *FileRepository) SaveMap(mapName string, mapData string) error {
	mapDir := filepath.Join(fr.BasePath, "maps")
	if err := fr.ensureDirectoryExists(mapDir); err != nil {
		return fmt.Errorf("erreur création dossier maps : %w", err)
	}

	mapFile := filepath.Join(mapDir, fmt.Sprintf("%s.json", mapName))
	file, err := os.Create(mapFile)
	if err != nil {
		return fmt.Errorf("erreur création fichier carte : %w", err)
	}
	defer file.Close()

	_, err = file.Write([]byte(mapData))
	return err
}

func (fr *FileRepository) LoadMap(mapName string) (string, error) {
	mapFile := filepath.Join(fr.BasePath, "maps", fmt.Sprintf("%s.json", mapName))
	file, err := os.Open(mapFile)
	if err != nil {
		return "", fmt.Errorf("erreur ouverture fichier carte : %w", err)
	}
	defer file.Close()

	data, err := os.ReadFile(mapFile)
	if err != nil {
		return "", fmt.Errorf("erreur lecture fichier carte : %w", err)
	}

	return string(data), nil
}
