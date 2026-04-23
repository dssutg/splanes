package main

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

// bombFrames are the sprite frames for bombs.
var bombFrames = []sdl.Rect{
	{X: 265, Y: 265, W: 9, H: 21},
}

// NewBomb creates a new bomb entity at the player's position.
func NewBomb() *Entity {
	e := NewEntity(EntityTypeBomb)

	e.TicksDelta = 1

	frame := bombFrames[0]
	e.Pos = sdl.Rect{X: 0, Y: 0, W: frame.W * 2, H: frame.H * 2}

	e.Damage = 1000
	e.Ticks = 1

	return e
}

// bombCalcSize calculates the current size of a shrinking bomb.
// The bomb shrinks over time as it falls.
func bombCalcSize(e *Entity) (w, h int32) {
	scale := math.Pow(0.90, float64(e.Ticks))

	w = int32(float64(e.Pos.W) * scale)
	h = int32(float64(e.Pos.H) * scale)

	if w == 0 || h == 0 {
		e.TicksDelta = 0
		w = 1
		h = 1
	}

	return w, h
}

// BombTick moves the bomb and checks for collisions with ships/submarines.
func BombTick(e *Entity) {
	e.Ticks += e.TicksDelta

	e.Pos.X += e.VelX
	e.Pos.Y += e.VelY

	if e.TicksDelta != 0 {
		return
	}

	scaledRect := sdl.Rect{X: e.Pos.X, Y: e.Pos.Y}
	scaledRect.W, scaledRect.H = bombCalcSize(e)

	dead := false

	for i := range EntityPool {
		it := &EntityPool[i]

		canAttack := it.Kind == EntityTypeShip || it.Kind == EntityTypeSubmarine

		if !canAttack {
			continue
		}

		if !it.Pos.HasIntersection(&scaledRect) {
			continue
		}

		// Bomb hits the target.
		dead = true
		it.Hurt(e.Damage)

		NewExplosion(scaledRect.X, scaledRect.Y)

		switch it.Kind {
		case EntityTypeShip:
			player.Score += 50
		case EntityTypeSubmarine:
			player.Score += 200
		}
		break
	}

	if dead {
		e.Remove()
		return
	}
}

// BombRender draws the bomb with a shrinking animation.
func BombRender(e *Entity) {
	w, h := bombCalcSize(e)
	RenderSprite(e.Texture, sdl.Rect{X: e.Pos.X, Y: e.Pos.Y, W: w, H: h}, bombFrames[0])
}
