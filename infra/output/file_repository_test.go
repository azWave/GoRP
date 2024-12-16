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
	repo := output.NewFileRepository(basePath)
	defer os.RemoveAll(basePath) // Nettoyer après test

	char := entities.Character{
		Name:  "TestHero",
		Class: "Guerrier",
		Stats: entities.Stats{
			Health:       100,
			Mana:         50,
			Strength:     15,
			Intelligence: 5,
			Defense:      10,
			MagicResist:  8,
			Agility:      7,
			Luck:         4,
			Endurance:    9,
			Spirit:       6,
		},
	}

	// Test sauvegarde personnage
	err = repo.SaveCharacter(char)
	if err != nil {
		t.Fatalf("Erreur sauvegarde personnage : %v", err)
	}

	// Test chargement personnage
	loadedChar, err := repo.LoadCharacter("TestHero")
	if err != nil {
		t.Fatalf("Erreur chargement personnage : %v", err)
	}

	// Test de la cohérence des données chargées avec celles sauvegardées
	if loadedChar.Name != char.Name {
		t.Errorf("Nom incorrect : attendu %s, obtenu %s", char.Name, loadedChar.Name)
	}
}
