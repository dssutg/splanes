package main

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

var bombFrames = []sdl.Rect{
	{X: 265, Y: 265, W: 9, H: 21},
}

func NewBomb() *Entity {
	e := NewEntity(EntityTypeBomb)

	e.TicksDelta = 1

	frame := bombFrames[0]
	e.Pos = sdl.Rect{X: 0, Y: 0, W: frame.W * 2, H: frame.H * 2}

	e.Damage = 1000
	e.Ticks = 1

	return e
}

func bombCalcSize(e *Entity) (w, h int32) {
	scale := math.Pow(0.90, float64(e.Ticks))

	w = int32(float64(e.Pos.W) * scale)
	h = int32(float64(e.Pos.H) * scale)

	if w == 0 || h == 0 {
		e.TicksDelta = 0
		w = 1
		h = 1
	}

	return
}

func BombTick(e *Entity) {
	e.Ticks += e.TicksDelta

	e.Pos.X += e.VelX
	e.Pos.Y += e.VelY

	if e.TicksDelta != 0 {
		return
	}

	scaledRect := sdl.Rect{X: e.Pos.X, Y: e.Pos.Y}
	scaledRect.W, scaledRect.H = bombCalcSize(e)

	for i := range EntityPool {
		it := &EntityPool[i]
		if it.Kind == EntityTypeShip && it.Pos.HasIntersection(&scaledRect) {
			it.Hurt(e.Damage)
			NewExplosion(scaledRect.X, scaledRect.Y)
			player.Score++
			break
		}
	}

	e.Remove()
}

func BombRender(e *Entity) {
	w, h := bombCalcSize(e)
	RenderSprite(e.Texture, sdl.Rect{X: e.Pos.X, Y: e.Pos.Y, W: w, H: h}, bombFrames[0])
}
