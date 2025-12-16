package entity

import "core:math/rand"

import "../gfx"
import "../kbd"
import "../menu"

new_player :: proc() -> ^Entity {
	e := new_entity(.Player)

	e.texture = 0

	e.xa = 0
	e.ya = 0

	e.crop = {299, 101, 61, 49}

	e.pos.w = e.crop.w * 2
	e.pos.h = e.crop.h * 2
	e.pos.x = (gfx.Window_Width - e.pos.w) / 2
	e.pos.y = gfx.Window_Height - e.pos.h - 40

	e.has_shot = false
	e.has_bombed = false

	e.health = Max_Player_Health

	e.score = 0
	e.distance = 0
	e.death_time = 0
	e.bomb_tick_time = Player_Max_Bomb_Tick_Time

	return e
}

player_do_die :: proc(e: ^Entity) {
	e.death_time += 1

	if e.death_time == 1 {
		explosion_count: i32 : 10
		for i in 0 ..< explosion_count {
			x := e.pos.x + rand.int31_max(e.pos.w)
			y := e.pos.y + rand.int31_max(e.pos.h / 2)
			new_explosion(x, y)
		}
		return
	}

	if e.death_time > 10 {
		menu.menu_ID = .Lose
	}
}

heal_player :: proc(heal_points: i32) {
	player.health += heal_points
	if player.health > Max_Player_Health {
		player.health = Max_Player_Health
	}
}

player_tick :: proc(e: ^Entity) {
	if e.health <= 0 {
		player_do_die(e)
		return
	}

	e.xa = 0

	if kbd.keys[.Left] {
		e.xa = -20
	}

	if kbd.keys[.Right] {
		e.xa = 20
	}

	if kbd.keys[.Bomb] && !e.has_bombed {
		e.has_bombed = true
		e.bomb_tick_time = 0

		bomb := new_bomb()
		bomb.pos.x = e.pos.x + (e.pos.w - bomb.pos.w) / 2 - 10
		bomb.pos.y = e.pos.y - bomb.pos.h
	} else {
		if e.has_bombed {
			e.bomb_tick_time += 1
			if e.bomb_tick_time >= Player_Max_Bomb_Tick_Time {
				e.has_bombed = false
				e.bomb_tick_time = Player_Max_Bomb_Tick_Time
			}
		}
	}

	if kbd.keys[.Shoot] && !e.has_shot {
		e.has_shot = true

		damage: i32 : 50

		bullet := new_bullet(.Bullet, 0, 0, 0, -1, e.type, damage, 0)
		bullet.pos.x = e.pos.x + (e.pos.w - bullet.pos.w) / 2 - 10
		bullet.pos.y = e.pos.y - bullet.pos.h
	} else if e.has_shot {
		e.tick_time += 1
		if e.tick_time > 3 {
			e.has_shot = false
			e.tick_time = 0
		}
	}

	xn := e.pos.x + e.xa
	yn := e.pos.y + e.ya

	if xn + e.pos.w >= gfx.Window_Width + 1 {
		xn = gfx.Window_Width - e.pos.w
	}

	if xn < 0 {
		xn = 0
	}

	e.pos.x = xn
	e.pos.y = yn

	e.distance += 1
}

player_render :: proc(e: ^Entity) {
	if e.death_time < 5 {
		render_entity_sprite(e)
	}
}
