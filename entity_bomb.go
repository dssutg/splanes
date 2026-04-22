package main

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

var bombFrames = []sdl.Rect{
	{X: 265, Y: 265, W: 9, H: 21},
}

func newBomb() *Entity {
	e := newEntity(EntityTypeBomb)

	e.texture = 0
	e.data = 1

	frame := bombFrames[0]
	e.pos = sdl.Rect{X: 0, Y: 0, W: frame.W * 2, H: frame.H * 2}

	e.xa = 0
	e.ya = 0

	e.damage = 1000
	e.tickTime = 1

	return e
}

func bombCalcSize(e *Entity) (w, h int32) {
	scale := math.Pow(0.90, float64(e.tickTime))

	w = int32(float64(e.pos.W) * scale)
	h = int32(float64(e.pos.H) * scale)

	if w == 0 || h == 0 {
		e.data = 0
		w = 1
		h = 1
	}

	return
}

func bombTick(e *Entity) {
	e.tickTime += e.data

	e.pos.X += e.xa * 2
	e.pos.Y += e.ya * 2

	if e.data != 0 {
		return
	}

	scaledRect := sdl.Rect{
		X: e.pos.X,
		Y: e.pos.Y,
	}

	scaledRect.W, scaledRect.H = bombCalcSize(e)

	for i := range entityPool {
		it := &entityPool[i]
		if it.etype == EntityTypeShip && it.pos.HasIntersection(&scaledRect) {
			hurtEntity(it, e.damage)
			newExplosion(scaledRect.X, scaledRect.Y)
			player.score++
			break
		}
	}

	removeEntity(e)
}

func bombRender(e *Entity) {
	w, h := bombCalcSize(e)
	renderSprite(e.texture, sdl.Rect{X: e.pos.X, Y: e.pos.Y, W: w, H: h}, bombFrames[0])
}
