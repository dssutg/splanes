package main

import (
	"math"
	"math/rand/v2"

	"github.com/veandco/go-sdl2/sdl"
)

// NewPlayer creates a new player entity and initializes its properties.
// The player starts at the bottom-center of the screen with full health.
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

// PlayerDoDie handles the player death sequence.
// It spawns explosions and transitions to the game over menu after a delay.
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

// PlayerHeal increases the player's health, clamping to the maximum.
func PlayerHeal(e *Entity, healPoints int32) {
	e.Health += healPoints
	if e.Health > MaxPlayerHealth {
		e.Health = MaxPlayerHealth
	}
}

// PlayerTick updates the player entity each game tick.
// It handles input, movement, shooting, bombing, and rotation.
func PlayerTick(e *Entity) {
	if e.Health <= 0 {
		PlayerDoDie(e)
		return
	}

	// Handle rotation input.
	const rotationDelta = 10
	if Keys[KeyRotateLeft] {
		e.Rotation -= rotationDelta
	}
	if Keys[KeyRotateRight] {
		e.Rotation += rotationDelta
	}
	// Restrict rotation freedom.
	e.Rotation = Clamp(e.Rotation, -30, 30)

	// Movement constants.
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

	// Apply acceleration and friction.
	{
		accelX := e.AccelX

		// Apply friction if no input.
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

	// Handle bomb dropping.
	if Keys[KeyBomb] && !e.HasBombed {
		e.HasBombed = true
		e.BombTicks = 0

		bomb := NewBomb()
		bomb.Pos.X = e.Pos.X + (e.Pos.W-bomb.Pos.W)/2 - 10
		bomb.Pos.Y = e.Pos.Y - bomb.Pos.H
	} else if e.HasBombed {
		e.BombTicks++
		if e.BombTicks >= PlayerMaxBombTickTime {
			e.HasBombed = false
			e.BombTicks = PlayerMaxBombTickTime
		}
	}

	// Handle shooting.
	if Keys[KeyShoot] && !e.HasShot {
		// Apply cool down.
		e.HasShot = true

		sin, cos := math.Sincos(DegToRad(float64(e.Rotation) - 90))

		// Determine bullet direction.
		const dirRadius = 3 // must be small but precise enough
		dirX := int32(cos * dirRadius)
		dirY := int32(sin * dirRadius)

		// Spawn the bullet.
		const damage = 50
		bullet := NewBullet(EntityTypeBullet, 0, 0, dirX, dirY, e.Kind, damage, 0, e.Rotation)

		// Adjust the bullet initial position.
		centerX := float64(e.Pos.X + (e.Pos.W-bullet.Pos.W)/2)
		centerY := float64(e.Pos.Y + (e.Pos.H-bullet.Pos.H)/2)
		playerRadius := math.Hypot(float64(e.Pos.W)/2, float64(e.Pos.H)/2)
		const awayBias = 50
		awayDistance := playerRadius - awayBias
		bullet.Pos.X = int32(centerX + cos*awayDistance)
		bullet.Pos.Y = int32(centerY + sin*awayDistance)
	} else if e.HasShot {
		e.Ticks++
		if e.Ticks > 3 {
			// Cool down is done, can shoot again.
			e.HasShot = false
			e.Ticks = 0
		}
	}

	// Apply velocity.
	xn := e.Pos.X + e.VelX
	yn := e.Pos.Y + e.VelY

	// Keep player within screen bounds.
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

// PlayerRender draws the player entity.
// The player is only rendered during the death animation briefly.
func PlayerRender(e *Entity) {
	if e.DeathTime < 5 {
		e.RenderSprite()
	}
}
