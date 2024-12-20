package services_test

import (
	"gorp/core/domain/entities"
	"gorp/core/services"
	"testing"
)

func TestMapService(t *testing.T) {
	serviceMap := services.MapService{}
	t.Run("Test ajout d'obstacles", func(t *testing.T) {
		// Test d'ajout d'obstacles
		serviceMap.InitializeMap(10, 10)
		serviceMap.AddTile(0, 0, entities.Water, entities.TileTypes[entities.Water].HasCollision)
		if serviceMap.GameMap.Tiles[0][0].Type != entities.Water {
			t.Error("L'ajout d'obstacle a échoué")
		}
	})

	t.Run("Test ajout d'obstacles aléatoires", func(t *testing.T) {
		// Test d'ajout d'obstacles aléatoires
		serviceMap.InitializeMap(10, 10)
		serviceMap.AddRandomObstacles(5)
		count := 0
		for _, line := range serviceMap.GameMap.Tiles {
			for _, tile := range line {
				if tile.Type == entities.Water || tile.Type == entities.Rock {
					count++
				}
			}
		}
		if count != 5 {
			t.Error("Le nombre d'obstacles ajoutés est incorrect")
		}
	})

}
