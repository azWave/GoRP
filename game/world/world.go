package world

type World struct {
	Map [][]rune
}

func GenerateWorld(size int) *World {
	size += 2
	world := &World{
		Map: make([][]rune, size),
	}

	for i := range world.Map {
		world.Map[i] = make([]rune, size)
		for j := range world.Map[i] {
			if i == 0 || i == size-1 || j == 0 || j == size-1 {
				world.Map[i][j] = '#'
			} else {
				world.Map[i][j] = '.'
			}
		}
	}

	return world
}

func (w *World) IsWalkable(x, y int) bool {
	if y < 0 || y >= len(w.Map) || x < 0 || x >= len(w.Map[y]) {
		return false
	}
	return w.Map[y][x] == '.'
}
