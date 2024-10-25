package game

type Player struct {
	X, Y int
}

func NewPlayer(x, y int) *Player {
	return &Player{x, y}
}

func (p *Player) Move(direction string, world *World) {
	switch direction {
	case "s":
		if(world.IsWalkable(p.X, p.Y+1)) {
			p.Y++
		}
	case "z":
		if(world.IsWalkable(p.X, p.Y-1)) {
			p.Y--
		}
	case "q":
		if(world.IsWalkable(p.X-1, p.Y)) {
			p.X--
		}
	case "d":
		if(world.IsWalkable(p.X+1, p.Y)) {
			p.X++
		}
	}
}