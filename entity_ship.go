package main

import (
	"math/rand/v2"

	"github.com/veandco/go-sdl2/sdl"
)

var shipFrames = []sdl.Rect{
	{X: 505, Y: 298, W: 41, H: 197},
	{X: 463, Y: 298, W: 41, H: 197},
}

func newShip() *Entity {
	e := newEntity(EntityTypeShip)

	e.texture = 0

	frame := &shipFrames[0]

	e.pos.W = frame.W * 1
	e.pos.H = frame.H * 1
	e.pos.X = rand.Int32N(WindowWidth)
	e.pos.Y = -rand.Int32N(WindowHeight) - e.pos.H

	e.xa = 0
	e.ya = 1

	e.hasShot = false
	e.hasBombed = false
	e.health = 100

	return e
}

func shipTick(e *Entity) {
	e.tickTime++
	if e.tickTime > 10 {
		e.tickTime = 0
	}

	if e.health <= 0 {
		newExplosion(e.pos.X, e.pos.Y)
		removeEntity(e)
		return
	}

	e.pos.X += e.xa * 11
	e.pos.Y += e.ya * 11

	if e.pos.Y >= WindowHeight {
		removeEntity(e)
	}
}

func shipRender(e *Entity) {
	frameNo := e.tickTime / 5 % int32(len(shipFrames))
	renderSprite(e.texture, e.pos, shipFrames[frameNo])
}
