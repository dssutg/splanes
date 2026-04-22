package main

import (
	"math/rand/v2"

	"github.com/veandco/go-sdl2/sdl"
)

func newPlayer() *Entity {
	e := newEntity(EntityTypePlayer)

	e.texture = 0

	e.xa = 0
	e.ya = 0

	e.crop = sdl.Rect{X: 299, Y: 101, W: 61, H: 49}

	e.pos.W = e.crop.W * 2
	e.pos.H = e.crop.H * 2
	e.pos.X = (WindowWidth - e.pos.W) / 2
	e.pos.Y = WindowHeight - e.pos.H - 40

	e.hasShot = false
	e.hasBombed = false

	e.health = MaxPlayerHealth

	e.score = 0
	e.distance = 0
	e.deathTime = 0
	e.bombTickTime = PlayerMaxBombTickTime

	return e
}

func playerDoDie(e *Entity) {
	e.deathTime++

	if e.deathTime == 1 {
		const explosionCount = 10
		for range explosionCount {
			x := e.pos.X + rand.Int32N(e.pos.W)
			y := e.pos.Y + rand.Int32N(e.pos.H/2)
			newExplosion(x, y)
		}
		return
	}

	if e.deathTime > 10 {
		menuID = MenuTypeLose
	}
}

func healPlayer(healPoints int32) {
	player.health += healPoints
	if player.health > MaxPlayerHealth {
		player.health = MaxPlayerHealth
	}
}

func playerTick(e *Entity) {
	if e.health <= 0 {
		playerDoDie(e)
		return
	}

	e.xa = 0

	if keys[KeyLeft] {
		e.xa = -20
	}

	if keys[KeyRight] {
		e.xa = 20
	}

	if keys[KeyBomb] && !e.hasBombed {
		e.hasBombed = true
		e.bombTickTime = 0

		bomb := newBomb()
		bomb.pos.X = e.pos.X + (e.pos.W-bomb.pos.W)/2 - 10
		bomb.pos.Y = e.pos.Y - bomb.pos.H
	} else {
		if e.hasBombed {
			e.bombTickTime++
			if e.bombTickTime >= PlayerMaxBombTickTime {
				e.hasBombed = false
				e.bombTickTime = PlayerMaxBombTickTime
			}
		}
	}

	if keys[KeyShoot] && !e.hasShot {
		e.hasShot = true

		const damage = 50

		bullet := newBullet(EntityTypeBullet, 0, 0, 0, -1, e.etype, damage, 0)
		bullet.pos.X = e.pos.X + (e.pos.W-bullet.pos.W)/2 - 10
		bullet.pos.Y = e.pos.Y - bullet.pos.H
	} else if e.hasShot {
		e.tickTime++
		if e.tickTime > 3 {
			e.hasShot = false
			e.tickTime = 0
		}
	}

	xn := e.pos.X + e.xa
	yn := e.pos.Y + e.ya

	if xn+e.pos.W >= WindowWidth+1 {
		xn = WindowWidth - e.pos.W
	}

	if xn < 0 {
		xn = 0
	}

	e.pos.X = xn
	e.pos.Y = yn

	e.distance++
}

func playerRender(e *Entity) {
	if e.deathTime < 5 {
		renderEntitySprite(e)
	}
}
