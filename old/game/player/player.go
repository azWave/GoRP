package player

type Player struct {
	X, Y int
	Name string
}

func New(x, y int, name string) *Player {
	return &Player{x, y, name}
}
