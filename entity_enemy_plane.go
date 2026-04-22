package main

import (
	"math/rand/v2"

	"github.com/veandco/go-sdl2/sdl"
)

var enemyPlaneFrames = []sdl.Rect{
	{X: 1, Y: 1, W: 32, H: 32},
	{X: 1, Y: 34, W: 32, H: 32},
	{X: 1, Y: 67, W: 32, H: 32},
	{X: 1, Y: 100, W: 32, H: 32},
	{X: 1, Y: 133, W: 32, H: 32},
}

func newEnemyPlane() *Entity {
	e := newEntity(EntityTypeEnemyPlane)

	e.texture = 0

	e.crop = enemyPlaneFrames[rand.IntN(len(enemyPlaneFrames))]

	e.pos.W = e.crop.W * 2
	e.pos.H = e.crop.H * 2
	e.pos.X = rand.Int32N(WindowWidth)
	e.pos.Y = -rand.Int32N(WindowHeight) - e.pos.H

	e.xa = 0
	e.ya = 1

	e.hasShot = false
	e.hasBombed = false
	e.health = 100

	return e
}

func enemyPlaneTick(e *Entity) {
	if e.health <= 0 {
		newExplosion(e.pos.X, e.pos.Y)
		removeEntity(e)
		return
	}

	if !e.hasShot {
		e.hasShot = true
		x := e.pos.X + (e.pos.W-128)/2
		y := e.pos.Y + e.pos.W
		newBullet(EntityTypeBullet, x, y, 0, 2, e.etype, 2, 1)
	} else {
		if e.hasShot {
			e.tickTime++
			if e.tickTime > 20 {
				e.hasShot = false
				e.tickTime = 0
			}
		}
	}

	e.pos.X += e.xa * 20
	e.pos.Y += e.ya * 20

	if e.pos.Y >= WindowHeight {
		removeEntity(e)
	}

	if e.pos.HasIntersection(&player.pos) {
		player.health = 0
	}
}

func enemyPlaneRender(e *Entity) {
	renderEntitySprite(e)
}
