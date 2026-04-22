package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

var bulletFrames = []sdl.Rect{
	{X: 1, Y: 166, W: 32, H: 32},
	{X: 34, Y: 199, W: 32, H: 32},
}

func newBullet(etype EntityType, x, y, xa, ya int32, ownerType EntityType, damage, bulletFrameNo int32) *Entity {
	e := newEntity(etype)

	e.texture = 0
	e.data = bulletFrameNo

	e.pos.X = x
	e.pos.Y = y
	e.pos.W = bulletFrames[bulletFrameNo].W * 2
	e.pos.H = bulletFrames[bulletFrameNo].H * 2

	e.xa = xa
	e.ya = ya

	e.ownerType = ownerType
	e.damage = damage

	return e
}

func bulletTick(e *Entity) {
	e.pos.X += e.xa * 20
	e.pos.Y += e.ya * 20

	dead := false

	if e.pos.HasIntersection(&WindowRect) {
		for i := range entityPool {
			other := &entityPool[i]
			if other.etype == EntityTypeNone {
				continue
			}

			canCollide := (e.ownerType == EntityTypePlayer && other.etype == EntityTypeEnemyPlane) ||
				(e.ownerType == EntityTypeEnemyPlane && other.etype == EntityTypePlayer)

			if !canCollide {
				continue
			}

			if !other.pos.HasIntersection(&e.pos) {
				continue
			}

			dead = true
			hurtEntity(other, e.damage)
			if other.etype == EntityTypePlayer {
				newExplosion(e.pos.X, e.pos.Y)
				continue
			}

			if e.ownerType == EntityTypePlayer {
				player.score++
			}
		}
	} else {
		dead = true
	}

	if dead {
		removeEntity(e)
	}
}

func bulletRender(e *Entity) {
	renderSprite(e.texture, e.pos, bulletFrames[e.data])
}
