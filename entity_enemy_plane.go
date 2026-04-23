package main

import (
	"math/rand/v2"

	"github.com/veandco/go-sdl2/sdl"
)

// enemyPlaneFrames are the animation frames for enemy planes.
var enemyPlaneFrames = []sdl.Rect{
	{X: 1, Y: 1, W: 32, H: 32},
	{X: 1, Y: 34, W: 32, H: 32},
	{X: 1, Y: 67, W: 32, H: 32},
	{X: 1, Y: 100, W: 32, H: 32},
	{X: 1, Y: 133, W: 32, H: 32},
}

// NewEnemyPlane creates an enemy plane that flies downward.
// It spawns at a random X position above the screen.
func NewEnemyPlane() *Entity {
	e := NewEntity(EntityTypeEnemyPlane)

	e.Crop = enemyPlaneFrames[rand.IntN(len(enemyPlaneFrames))]

	e.Pos.W = e.Crop.W * 2
	e.Pos.H = e.Crop.H * 2
	e.Pos.X = rand.Int32N(WindowW)
	e.Pos.Y = -rand.Int32N(WindowH) - e.Pos.H

	e.VelY = 20

	e.Health = 100

	return e
}

// EnemyPlaneTick moves the enemy plane and makes it shoot at the player.
func EnemyPlaneTick(e *Entity) {
	// If dead, explode and remove it.
	if e.Health <= 0 {
		NewExplosion(e.Pos.X, e.Pos.Y)
		e.Remove()
		return
	}

	// Handle shooting with cooldown.
	if e.HasShot {
		// Handle reload delay.
		e.Ticks++
		if e.Ticks > 20 {
			e.HasShot = false
			e.Ticks = 0
		}
	} else {
		// Shoot.
		e.HasShot = true
		x := e.Pos.X + (e.Pos.W-128)/2
		y := e.Pos.Y + e.Pos.W
		NewBullet(EntityTypeBullet, x, y, 0, 2, e.Kind, 2, 1, 0)
	}

	// Move downward.
	e.Pos.X += e.VelX
	e.Pos.Y += e.VelY

	// Remove when off screen.
	if e.Pos.Y >= WindowH {
		e.Remove()
		return
	}

	// Check collision with player.
	if e.Pos.HasIntersection(&player.Pos) {
		player.Health = 0
	}
}

// EnemyPlaneRender draws the enemy plane.
func EnemyPlaneRender(e *Entity) {
	e.RenderSprite()
}
