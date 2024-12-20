package services

import (
	"gorp/core/domain/entities"
	"math/rand"
	"time"
)

type MapService struct {
	GameMap *entities.Map
}

func (ms *MapService) InitializeMap(width, height int) {
	ms.GameMap = entities.NewMap(width, height, entities.TileTypes[entities.Land])
}

func (ms *MapService) AddTile(x, y int, tileType entities.TileType, hasCollision bool) error {
	tile := entities.Tile{
		Type:         tileType,
		HasCollision: hasCollision,
		DisplayChar:  entities.TileTypes[tileType].DisplayChar,
	}
	return ms.GameMap.SetTile(x, y, tile)
}

func (ms *MapService) AddRandomObstacles(count int) {
	rand.Seed(time.Now().UnixNano())
	placed := 0
	for placed < count {
		x := rand.Intn(ms.GameMap.Width)
		y := rand.Intn(ms.GameMap.Height)

		// Vérifie si la cellule est déjà un obstacle
		tile, _ := ms.GameMap.GetTile(x, y)
		if tile.Type == entities.Land {
			obstacleType := entities.Water
			if rand.Intn(2) == 1 {
				obstacleType = entities.Rock
			}
			ms.AddTile(x, y, obstacleType, entities.TileTypes[obstacleType].HasCollision)
			placed++
		}
	}
}

func (ms *MapService) AddObstacles(obstacles []struct {
	X            int
	Y            int
	TileType     entities.TileType
	HasCollision bool
}) {
	for _, obstacle := range obstacles {
		ms.AddTile(obstacle.X, obstacle.Y, obstacle.TileType, obstacle.HasCollision)
	}
}
