package main

import (
	"math/rand/v2"

	"github.com/veandco/go-sdl2/sdl"
)

// healerFrames are the sprite frames for health power-ups.
var healerFrames = []sdl.Rect{
	{X: 166, Y: 265, W: 29, H: 15},
}

// NewHealer creates a health power-up.
// Collecting it restores 20 health points.
func NewHealer() *Entity {
	e := NewEntity(EntityTypeHealer)

	e.Pos.W = healerFrames[0].W * 2
	e.Pos.H = healerFrames[0].H * 2
	e.Pos.X = rand.Int32N(WindowW)
	e.Pos.Y = -rand.Int32N(WindowH) - e.Pos.H

	e.VelY = 1

	return e
}

// HealerTick moves the power-up and checks for player pickup.
func HealerTick(e *Entity) {
	e.Pos.X += e.VelX * 10
	e.Pos.Y += e.VelY * 10

	if e.Pos.Y >= WindowH {
		e.Remove()
		return
	}

	if e.Pos.HasIntersection(&player.Pos) {
		PlayerHeal(player, 20)
		e.Remove()
		return
	}
}

// HealerRender draws the power-up.
func HealerRender(e *Entity) {
	RenderSprite(e.Texture, e.Pos, healerFrames[0])
}
