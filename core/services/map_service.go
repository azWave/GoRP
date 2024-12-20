package services

import (
	"gorp/core/domain/entities"
	"math/rand"
)

type MapService struct{}

func (ms *MapService) GenerateMap(with, height int) *entities.Map {

	cells := make([][]entities.Cell, height)
	for y := range cells {
		row := make([]entities.Cell, with)
		for x := range row {
			row[x] = entities.Cell{Type: entities.Land}
		}
		cells[y] = row

	}

	GenerateRandomRocks(cells)
	return &entities.Map{With: with, Height: height, Cells: cells}
}

func GenerateRandomRocks(cells [][]entities.Cell) {
	rand.Seed(42)

	for i := 0; i < 5; i++ {
		x := rand.Intn(len(cells[0]))
		y := rand.Intn(len(cells))
		cells[y][x] = entities.Cell{Type: entities.Rock}
	}
}

func (ms *MapService) IsWallkable(x, y int, gameMap entities.Map) bool {
	if x < 0 || x >= gameMap.With || y < 0 || y >= gameMap.Height {
		return false
	}
	cell := gameMap.Cells[y][x]
	return entities.ClassCellProperties[cell.Type].Wallkable
}

func (ms *MapService) Move(currentX, currentY, dx, dy int, gameMap entities.Map) (int, int, bool) {
	newX, newY := currentX+dx, currentY+dy
	if !ms.IsWallkable(newX, newY, gameMap) {
		return currentX, currentY, false
	}
	return newX, newY, true
}
