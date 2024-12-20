package input_test

import (
	"gorp/core/domain/entities"
	"gorp/core/services"
	input "gorp/infra/input/map"
	"gorp/infra/output"
	"testing"
)

func TestPrintMap(t *testing.T) {
	// Configuration de la map
	mapService := services.MapService{}
	mapService.InitializeMap(10, 10)
	mapService.AddTile(1, 1, entities.Water, entities.TileTypes[entities.Water].HasCollision)

	character := entities.Character{
		Position: entities.Position{X: 0, Y: 0, MapName: "world-1"},
	}

	mockPrinter := &output.MockPrinter{}
	input.PrintMap(mockPrinter, mapService.GameMap, character.Position)
	actualOutput := mockPrinter.Output.String()

	// Sortie attendue
	expectedOutput := "£ . . . . . . . . . \n" +
		". ~ . . . . . . . . \n" +
		". . . . . . . . . . \n" +
		". . . . . . . . . . \n" +
		". . . . . . . . . . \n" +
		". . . . . . . . . . \n" +
		". . . . . . . . . . \n" +
		". . . . . . . . . . \n" +
		". . . . . . . . . . \n" +
		". . . . . . . . . . \n"

	// Comparaison de la sortie
	if actualOutput != expectedOutput {
		t.Errorf("PrintMap() = %v, attendu %v", actualOutput, expectedOutput)
	}
}
