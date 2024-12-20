package entities

import "errors"

type Map struct {
	Width  int
	Height int
	Tiles  [][]Tile
}

func NewMap(width, height int, defaultTile Tile) *Map {
	tiles := make([][]Tile, height)
	for i := range tiles {
		tiles[i] = make([]Tile, width)
		for j := range tiles[i] {
			tiles[i][j] = defaultTile
		}
	}
	return &Map{
		Width:  width,
		Height: height,
		Tiles:  tiles,
	}
}

func (m *Map) SetTile(x, y int, tile Tile) error {
	if x < 0 || x >= m.Width || y < 0 || y >= m.Height {
		return errors.New("coordonnées hors limites")
	}
	m.Tiles[y][x] = tile
	return nil
}

func (m *Map) GetTile(x, y int) (Tile, error) {
	if x < 0 || x >= m.Width || y < 0 || y >= m.Height {
		return Tile{}, errors.New("coordonnées hors limites")
	}
	return m.Tiles[y][x], nil
}
