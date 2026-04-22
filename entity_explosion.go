package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

var explosionFrames = []sdl.Rect{
	{X: 67, Y: 166, W: 32, H: 32},
	{X: 100, Y: 166, W: 32, H: 32},
	{X: 133, Y: 166, W: 32, H: 32},
	{X: 166, Y: 166, W: 32, H: 32},
	{X: 199, Y: 166, W: 32, H: 32},
	{X: 232, Y: 166, W: 32, H: 32},
}

func newExplosion(x, y int32) *Entity {
	e := newEntity(EntityTypeExplosion)

	e.texture = 0

	frame := explosionFrames[0]
	e.pos = sdl.Rect{X: x, Y: y, W: frame.W * 2, H: frame.H * 2}

	playSound(soundExplosion1, 100)

	return e
}

func explosionTick(e *Entity) {
	e.tickTime++
	if e.tickTime >= int32(len(explosionFrames)) {
		removeEntity(e)
	}
}

func explosionRender(e *Entity) {
	frameNo := e.tickTime % int32(len(explosionFrames))
	renderSprite(e.texture, e.pos, explosionFrames[frameNo])
}
