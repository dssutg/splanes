package main

import (
	"math/rand/v2"

	"github.com/veandco/go-sdl2/sdl"
)

func NewPlayer() *Entity {
	e := NewEntity(EntityTypePlayer)

	e.Crop = sdl.Rect{X: 299, Y: 101, W: 61, H: 49}

	e.Pos.W = e.Crop.W * 2
	e.Pos.H = e.Crop.H * 2
	e.Pos.X = (WindowW - e.Pos.W) / 2
	e.Pos.Y = WindowH - e.Pos.H - 40

	e.Health = MaxPlayerHealth

	e.BombTicks = PlayerMaxBombTickTime

	return e
}

func PlayerDoDie(e *Entity) {
	e.DeathTime++

	if e.DeathTime == 1 {
		const explosionCount = 10
		for range explosionCount {
			x := e.Pos.X + rand.Int32N(e.Pos.W)
			y := e.Pos.Y + rand.Int32N(e.Pos.H/2)
			NewExplosion(x, y)
		}
		return
	}

	if e.DeathTime > 10 {
		menuID = MenuTypeLose
	}
}

func PlayerHeal(e *Entity, healPoints int32) {
	e.Health += healPoints
	if e.Health > MaxPlayerHealth {
		e.Health = MaxPlayerHealth
	}
}

func PlayerTick(e *Entity) {
	if e.Health <= 0 {
		PlayerDoDie(e)
		return
	}

	const (
		minVel         = -20
		maxVel         = 20
		accelFactor    = 4
		slowdownFactor = 1
	)

	// Determine current acceleration based on input.
	// If the player holds the two keys at the same
	// time, acceleration becomes zero because the
	// forces negate themselves.
	e.AccelX = 0
	if Keys[KeyLeft] {
		e.AccelX -= accelFactor
	}
	if Keys[KeyRight] {
		e.AccelX += accelFactor
	}

	// Gradually update the velocity.
	{
		accelX := e.AccelX

		if accelX == 0 {
			// Gradually slow down the velocity by the inverse of the acceleration.
			// But make sure that we don't get pass the zero velocity,
			// otherwise the player won't stop.
			if e.VelX > 0 {
				accelX = -min(e.VelX, slowdownFactor)
			} else {
				accelX = min(-e.VelX, slowdownFactor)
			}
		}

		e.VelX = Clamp(e.VelX+accelX, minVel, maxVel)
	}

	if Keys[KeyBomb] && !e.HasBombed {
		e.HasBombed = true
		e.BombTicks = 0

		bomb := NewBomb()
		bomb.Pos.X = e.Pos.X + (e.Pos.W-bomb.Pos.W)/2 - 10
		bomb.Pos.Y = e.Pos.Y - bomb.Pos.H
	} else {
		if e.HasBombed {
			e.BombTicks++
			if e.BombTicks >= PlayerMaxBombTickTime {
				e.HasBombed = false
				e.BombTicks = PlayerMaxBombTickTime
			}
		}
	}

	if Keys[KeyShoot] && !e.HasShot {
		e.HasShot = true

		const damage = 50

		bullet := NewBullet(EntityTypeBullet, 0, 0, 0, -1, e.Kind, damage, 0)
		bullet.Pos.X = e.Pos.X + (e.Pos.W-bullet.Pos.W)/2 - 10
		bullet.Pos.Y = e.Pos.Y - bullet.Pos.H
	} else if e.HasShot {
		e.Ticks++
		if e.Ticks > 3 {
			e.HasShot = false
			e.Ticks = 0
		}
	}

	xn := e.Pos.X + e.VelX
	yn := e.Pos.Y + e.VelY

	if xn+e.Pos.W >= WindowW+1 {
		xn = WindowW - e.Pos.W
	}

	if xn < 0 {
		xn = 0
	}

	e.Pos.X = xn
	e.Pos.Y = yn

	e.Distance++
}

func PlayerRender(e *Entity) {
	if e.DeathTime < 5 {
		e.RenderSprite()
	}
}
