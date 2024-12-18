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
	defer os.RemoveAll(basePath) // Nettoyer après test

	repo := &output.FileRepository{BasePath: basePath}
	service := &services.CharacterService{Repo: repo}

	tests := []struct {
		name       string
		inputName  string
		inputClass string
		wantErr    string
	}{
		{"Character valide", "Aragorn", "Guerrier", ""},
		{"Nom trop court", "Al", "Guerrier", "nom trop court"},
		{"Nom trop long", "UnNomExcessivementLongPourUnHeros", "Guerrier", "nom trop long"},
		{"Classe inconnue", "Aragorn", "UnknownClass", "classe inconnue"},
	}

	// Test CreateCharacter
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := service.CreateCharacter(tt.inputName, tt.inputClass)
			if err == nil && tt.wantErr != "" {
				t.Fatalf("Erreur attendue : %v, obtenue : nil", tt.wantErr)
			}
			if err != nil && err.Error() != tt.wantErr {
				t.Fatalf("Erreur attendue : %v, obtenue : %v", tt.wantErr, err.Error())
			}
		})
	}
}
