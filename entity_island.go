package main

import (
	"math/rand/v2"

	"github.com/veandco/go-sdl2/sdl"
)

var islandFrames = []sdl.Rect{
	{X: 100, Y: 496, W: 64, H: 65},
	{X: 165, Y: 496, W: 64, H: 65},
	{X: 230, Y: 496, W: 64, H: 65},
}

func newIsland() *Entity {
	e := newEntity(EntityTypeIsland)

	e.texture = 0
	e.data = rand.Int32N(3)

	e.pos.W = islandFrames[e.data].W * 3
	e.pos.H = islandFrames[e.data].H * 3
	e.pos.X = rand.Int32N(WindowWidth)
	e.pos.Y = -rand.Int32N(WindowHeight) - e.pos.H

	e.xa = 0
	e.ya = 1

	return e
}

func islandTick(e *Entity) {
	e.pos.X += e.xa * 10
	e.pos.Y += e.ya * 10
	if e.pos.Y >= WindowHeight {
		removeEntity(e)
	}
}

func islandRender(e *Entity) {
	renderSprite(e.texture, e.pos, islandFrames[e.data])
}
