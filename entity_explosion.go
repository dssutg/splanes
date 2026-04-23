package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

// explosionFrames are the animation frames for explosions.
var explosionFrames = []sdl.Rect{
	{X: 67, Y: 166, W: 32, H: 32},
	{X: 100, Y: 166, W: 32, H: 32},
	{X: 133, Y: 166, W: 32, H: 32},
	{X: 166, Y: 166, W: 32, H: 32},
	{X: 199, Y: 166, W: 32, H: 32},
	{X: 232, Y: 166, W: 32, H: 32},
}

// NewExplosion creates an explosion effect at the given position.
func NewExplosion(x, y int32) {
	e := NewEntity(EntityTypeExplosion)

	frame := explosionFrames[0]
	e.Pos = sdl.Rect{X: x, Y: y, W: frame.W * 2, H: frame.H * 2}

	PlaySound(soundExplosion1, 100)
}

// ExplosionTick advances the explosion animation.
// The entity is removed when the animation completes.
func ExplosionTick(e *Entity) {
	e.Ticks++
	if e.Ticks >= int32(len(explosionFrames)) {
		e.Remove()
		return
	}
}

// ExplosionRender draws the current frame of the explosion animation.
func ExplosionRender(e *Entity) {
	frameNo := e.Ticks % int32(len(explosionFrames))
	RenderSprite(e.Texture, e.Pos, explosionFrames[frameNo])
}
