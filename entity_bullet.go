package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

var bulletFrames = []sdl.Rect{
	{X: 1, Y: 166, W: 32, H: 32},
	{X: 34, Y: 199, W: 32, H: 32},
}

func NewBullet(etype EntityType, x, y, dirX, dirY int32, ownerType EntityType, damage, bulletFrameNo int32) *Entity {
	e := NewEntity(etype)

	e.InitialFrameNo = bulletFrameNo

	e.Pos.X = x
	e.Pos.Y = y
	e.Pos.W = bulletFrames[bulletFrameNo].W * 2
	e.Pos.H = bulletFrames[bulletFrameNo].H * 2

	e.VelX = dirX * 20
	e.VelY = dirY * 20

	e.OwnerKind = ownerType
	e.Damage = damage

	return e
}

func BulletTick(e *Entity) {
	e.Pos.X += e.VelX
	e.Pos.Y += e.VelY

	dead := false

	if e.Pos.HasIntersection(&WindowRect) {
		for i := range EntityPool {
			other := &EntityPool[i]
			if other.Kind == EntityTypeNone {
				continue
			}

			canCollide := (e.OwnerKind == EntityTypePlayer && other.Kind == EntityTypeEnemyPlane) ||
				(e.OwnerKind == EntityTypeEnemyPlane && other.Kind == EntityTypePlayer)

			if !canCollide {
				continue
			}

			if !other.Pos.HasIntersection(&e.Pos) {
				continue
			}

			dead = true
			other.Hurt(e.Damage)
			if other.Kind == EntityTypePlayer {
				NewExplosion(e.Pos.X, e.Pos.Y)
				continue
			}

			if e.OwnerKind == EntityTypePlayer {
				player.Score++
			}
		}
	} else {
		dead = true
	}

	if dead {
		e.Remove()
	}
}

func BulletRender(e *Entity) {
	RenderSprite(e.Texture, e.Pos, bulletFrames[e.InitialFrameNo])
}
