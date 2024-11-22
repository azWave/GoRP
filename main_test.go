package main

import (
	"gorp/game/player"
	"gorp/game/world"
	"gorp/movements"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupPlayer() *player.Player {
	return player.New(0, 0, "John")
}

func setupWorld() *world.World {
	return world.GenerateWorld(5)
}

func TestNewPlayer(t *testing.T) {
	name := "John"
	p := setupPlayer()
	assert.Equal(t, name, p.Name, "Player name should be John")
}

func TestMovePlayer(t *testing.T) {
	p := setupPlayer()
	world := setupWorld()
	movements.Move(p, "d", world)
	assert.Equal(t, 1, p.X) // 1 right
}
