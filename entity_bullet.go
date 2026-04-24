package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

// bulletFrames are the sprite frames for bullets.
var bulletFrames = []sdl.Rect{
	{X: 1, Y: 166, W: 32, H: 32},
	{X: 34, Y: 199, W: 32, H: 32},
}

// NewBullet creates a bullet entity moving in the given direction.
// dirX and dirY should be -1, 0, or 1 to specify direction.
func NewBullet(
	etype EntityType,
	x, y int32,
	dirX, dirY int32,
	ownerType EntityType,
	damage int32,
	bulletFrameNo int32,
	rotation float32,
) *Entity {
	e := NewEntity(etype)

	e.InitialFrameNo = bulletFrameNo

	e.Pos.X = x
	e.Pos.Y = y
	e.Pos.W = bulletFrames[bulletFrameNo].W * 2
	e.Pos.H = bulletFrames[bulletFrameNo].H * 2

	e.VelX = dirX * 20
	e.VelY = dirY * 20

	e.Rotation = rotation

	e.OwnerKind = ownerType
	e.Damage = damage

	return e
}

// BulletTick moves the bullet and checks for collisions with enemies.
func BulletTick(e *Entity) {
	e.Pos.X += e.VelX
	e.Pos.Y += e.VelY

	// Remove the bullet if it's off screen.
	if !e.Pos.HasIntersection(&WindowRect) {
		e.Remove()
		return
	}

	dead := false

	// Iterate over all entities and see if the bullet collides with them.
	for i := range EntityPool {
		other := &EntityPool[i]
		if other.Kind == EntityTypeNone {
			continue
		}

		// Bullet can collide only with the entities it can attack and
		// who it does not belong to. Otherwise, it would cause friendly-fire.
		canCollide := (e.OwnerKind == EntityTypePlayer && other.Kind == EntityTypeEnemyPlane) ||
			(e.OwnerKind == EntityTypeEnemyPlane && other.Kind == EntityTypePlayer)

		if !canCollide {
			continue
		}

		// Check if bullet collides with the entity.
		if !other.Pos.HasIntersection(&e.Pos) {
			continue
		}

		// Bullet collides: destroy it, damage the target, and add relevant scores.
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

	if dead {
		e.Remove()
		return
	}
}

// BulletRender draws the bullet using the correct animation frame.
func BulletRender(e *Entity) {
	e.Crop = bulletFrames[e.InitialFrameNo]
	e.RenderSprite()
}
