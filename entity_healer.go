package main

import (
	"math/rand/v2"

	"github.com/veandco/go-sdl2/sdl"
)

var healerFrames = []sdl.Rect{
	{X: 166, Y: 265, W: 29, H: 15},
}

func newHealer() *Entity {
	e := newEntity(EntityTypeHealer)

	e.texture = 0
	e.data = 0

	e.pos.W = healerFrames[e.data].W * 2
	e.pos.H = healerFrames[e.data].H * 2
	e.pos.X = rand.Int32N(WindowWidth)
	e.pos.Y = -rand.Int32N(WindowHeight) - e.pos.H

	e.xa = 0
	e.ya = 1

	return e
}

func healerTick(e *Entity) {
	e.pos.X += e.xa * 10
	e.pos.Y += e.ya * 10

	if e.pos.Y >= WindowHeight {
		removeEntity(e)
	}

	if e.pos.HasIntersection(&player.pos) {
		healPlayer(20)
		removeEntity(e)
	}
}

func healerRender(e *Entity) {
	renderSprite(e.texture, e.pos, healerFrames[e.data])
}
