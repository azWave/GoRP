package output_test

import (
	"gorp/core/domain/entities"
	"gorp/infra/output"
	"gorp/pkg"
	"os"
	"path/filepath"
	"testing"
)

func TestFileRepository(t *testing.T) {
	tempDir, err := pkg.GetTestTempDir()
	if err != nil {
		t.Fatalf("Erreur lors de la création du dossier temporaire : %v", err)
	}
	basePath := filepath.Join(tempDir, "test_data")
	defer os.RemoveAll(basePath) // Nettoyer après test

	repo := output.NewFileRepository(basePath)
	char := entities.Character{
		Name:  "TestHero",
		Class: "Guerrier",
		Stats: entities.Stats{
			Health:       100,
			Mana:         45,
			Strength:     15,
			Intelligence: 5,
			Defense:      12,
			MagicResist:  6,
			Agility:      8,
			Luck:         5,
			Endurance:    10,
			Spirit:       4,
		},
	}

	t.Run("Test sauvegarde personnage", func(t *testing.T) {
		err := repo.SaveCharacter(char)
		if err != nil {
			t.Fatalf("Erreur sauvegarde personnage : %v", err)
		}
	})

	t.Run("Test chargement personnage", func(t *testing.T) {
		loadedChar, err := repo.LoadCharacter("TestHero")
		if err != nil {
			t.Fatalf("Erreur chargement personnage : %v", err)
		}

		if loadedChar != char {
			t.Errorf("Données incorrectes : attendu %v, obtenu %v", char, loadedChar)
		}
	})
}
