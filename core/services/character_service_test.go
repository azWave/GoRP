package services_test

import (
	"gorp/core/services"
	"gorp/infra/output"
	"gorp/pkg"
	"os"
	"path/filepath"
	"testing"
)

func TestCharacterService(t *testing.T) {
	tempDir, err := pkg.GetTestTempDir()
	if err != nil {
		t.Fatalf("Erreur lors de la création du dossier temporaire : %v", err)
	}
	basePath := filepath.Join(tempDir, "test_data")
	repo := &output.FileRepository{BasePath: basePath}
	defer os.RemoveAll(basePath) // Nettoyer après test

	// Test CreateCharacter with unknown class
	_, err = (&services.CharacterService{Repo: repo}).CreateCharacter("TestHero", "UnknownClass")
	if err == nil {
		t.Fatalf("Erreur attendue lors de la creation du personnage : %v", err)
	}

	// Test CreateCharacter with known class
	_, err = (&services.CharacterService{Repo: repo}).CreateCharacter("TestHero", "Guerrier")
	if err != nil {
		t.Fatalf("Erreur lors de la creation du personnage : %v", err)
	}
}
