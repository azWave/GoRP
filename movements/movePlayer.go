package movements

import (
	"gorp/game/player"
	"gorp/game/world"
)

func Move(p *player.Player, direction string, world *world.World) {
	switch direction {
	case "s":
		if world.IsWalkable(p.X, p.Y+1) {
			p.Y++
		}
	case "z":
		if world.IsWalkable(p.X, p.Y-1) {
			p.Y--
		}
	case "q":
		if world.IsWalkable(p.X-1, p.Y) {
			p.X--
		}
	case "d":
		if world.IsWalkable(p.X+1, p.Y) {
			p.X++
		}
	}
}
