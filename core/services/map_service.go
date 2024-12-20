package services

import (
	"gorp/core/domain/entities"
	"math/rand"
)

type MapService struct{}

func NewMapService() *MapService {
	return &MapService{}
}

func (ms *MapService) GenerateMap(width, height int) *entities.Map {

	cells := make([][]entities.Cell, height)
	for i := range cells {
		cells[i] = make([]entities.Cell, width)
		for j := range cells[i] {
			cells[i][j] = entities.Cell{Type: entities.Land}
		}
	}

	for i := 0; i < 5; i++ {
		x := rand.Intn(width)
		y := rand.Intn(height)
		cells[y][x] = entities.Cell{Type: entities.Rock}
	}

	for i := 0; i < 2; i++ {
		x := rand.Intn(width)
		y := rand.Intn(height)
		roomType := entities.RoomType(rand.Intn(3))
		cells[y][x] = entities.Cell{
			Type: entities.DonjonRoom,
			Room: &entities.Room{Type: roomType},
		}
	}

	return &entities.Map{Cells: cells}
}
