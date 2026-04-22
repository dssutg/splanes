package main

import (
	"math/rand/v2"

	"github.com/veandco/go-sdl2/sdl"
)

var shipFrames = []sdl.Rect{
	{X: 505, Y: 298, W: 41, H: 197},
	{X: 463, Y: 298, W: 41, H: 197},
}

func NewShip() *Entity {
	e := NewEntity(EntityTypeShip)

	frame := &shipFrames[0]

	e.Pos.W = frame.W * 1
	e.Pos.H = frame.H * 1
	e.Pos.X = rand.Int32N(WindowW)
	e.Pos.Y = -rand.Int32N(WindowH) - e.Pos.H

	e.VelY = 11

	e.HasShot = false
	e.HasBombed = false
	e.Health = 100

	return e
}

func ShipTick(e *Entity) {
	e.Ticks++
	if e.Ticks > 10 {
		e.Ticks = 0
	}

	if e.Health <= 0 {
		NewExplosion(e.Pos.X, e.Pos.Y)
		e.Remove()
		return
	}

	e.Pos.X += e.VelX
	e.Pos.Y += e.VelY

	if e.Pos.Y >= WindowH {
		e.Remove()
	}
}

func ShipRender(e *Entity) {
	frameNo := e.Ticks / 5 % int32(len(shipFrames))
	RenderSprite(e.Texture, e.Pos, shipFrames[frameNo])
}
