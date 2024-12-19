package entities

func (c *Character) Move(dx, dy int) {
	c.Position.X += dx
	c.Position.Y += dy
}
