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

func NewExplosion(x, y int32) *Entity {
	e := NewEntity(EntityTypeExplosion)

	frame := explosionFrames[0]
	e.Pos = sdl.Rect{X: x, Y: y, W: frame.W * 2, H: frame.H * 2}

	PlaySound(soundExplosion1, 100)

	return e
}

func ExplosionTick(e *Entity) {
	e.Ticks++
	if e.Ticks >= int32(len(explosionFrames)) {
		e.Remove()
	}
}

func ExplosionRender(e *Entity) {
	frameNo := e.Ticks % int32(len(explosionFrames))
	RenderSprite(e.Texture, e.Pos, explosionFrames[frameNo])
}
