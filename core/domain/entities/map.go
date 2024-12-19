package entities

import "fmt"

type CellType int

const (
	Land CellType = iota
	DonjonRoom
)

type RoomType int

const (
	Empty RoomType = iota
	Chest
	Monsters
)

type Cell struct {
	Type CellType
	Room *Room
}

type Room struct {
	Type RoomType
}

type Map struct {
	Cells [][]Cell
}

func (m *Map) Display(character *Character) {
	for y, row := range m.Cells {
		for x, cell := range row {
			if character.Position.X == x && character.Position.Y == y {
				fmt.Print("C ")
			} else if cell.Type == DonjonRoom {
				fmt.Print("R ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
}
