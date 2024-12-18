package input_test

import (
	"gorp/core/domain/entities"
	"gorp/infra/input"
	"gorp/infra/output"
	"testing"
)

func TestPrintCharacterSummary(t *testing.T) {
	// Configuration du personnage
	character := entities.Character{
		Name:  "Aragorn",
		Class: "Guerrier",
		Stats: entities.ClassStats["Guerrier"],
	}

	// actualOutput, err := pkg.CaptureFunctionConsoleOutput(func() { input.PrintCharacterSummary(character) })
	// if err != nil {
	// 	t.Errorf("CaptureFunctionConsoleOutput() a retourné une erreur : %v", err)
	// }
	mockPrinter := &output.MockPrinter{}
	input.PrintCharacterSummary(mockPrinter, character)
	actualOutput := mockPrinter.Output.String()

	// Sortie attendue
	expectedOutput := "Nom : Aragorn\n" +
		"Classe : Guerrier\n" +
		"PV : 150\n" +
		"PM : 50\n" +
		"Force : 15\n" +
		"Intelligence : 5\n" +
		"Défense : 12\n" +
		"Résistance Magique : 6\n" +
		"Agilité : 8\n" +
		"Chance : 5\n" +
		"Endurance : 10\n" +
		"Esprit : 4\n"

	// Comparaison de la sortie
	if actualOutput != expectedOutput {
		t.Errorf("PrintCharacterSummary() = %v, want %v", actualOutput, expectedOutput)
	}
}
