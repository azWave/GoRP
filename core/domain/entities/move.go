package entities

func (c *Character) Move(dx, dy int, gameMap *Map) {
	newX := c.Position.X + dx
	newY := c.Position.Y + dy

	// Check for collision with rocks
	if newX >= 0 && newX < len(gameMap.Cells[0]) && newY >= 0 && newY < len(gameMap.Cells) {
		if gameMap.Cells[newY][newX].Type != Rock {
			c.Position.X = newX
			c.Position.Y = newY
		}
	}
}
