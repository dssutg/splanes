package main

import "core:math/rand"

import SDL "vendor:sdl2"

enemy_plane_frames := []SDL.Rect {
	{1, 1, 32, 32},
	{1, 34, 32, 32},
	{1, 67, 32, 32},
	{1, 100, 32, 32},
	{1, 133, 32, 32},
}

new_enemy_plane :: proc() -> ^Entity {
	e := new_entity(.Enemy_Plane)

	e.texture = 0

	e.crop = enemy_plane_frames[rand.int31_max(i32(len(enemy_plane_frames)))]

	e.pos.w = e.crop.w * 2
	e.pos.h = e.crop.h * 2
	e.pos.x = rand.int31_max(Window_Width)
	e.pos.y = -rand.int31_max(Window_Height) - e.pos.h

	e.xa = 0
	e.ya = 1

	e.has_shot = false
	e.has_bombed = false
	e.health = 100

	return e
}

enemy_plane_tick :: proc(e: ^Entity) {
	if e.health <= 0 {
		new_explosion(e.pos.x, e.pos.y)
		e.removed = true
		return
	}

	if !e.has_shot {
		e.has_shot = true
		x := e.pos.x + (e.pos.w - 128) / 2
		y := e.pos.y + e.pos.w
		new_bullet(.Bullet, x, y, 0, 2, e.type, 2, 1)
	} else {
		if e.has_shot {
			e.tick_time += 1
			if e.tick_time > 20 {
				e.has_shot = false
				e.tick_time = 0
			}
		}
	}

	e.pos.x += e.xa * 20
	e.pos.y += e.ya * 20

	if e.pos.y >= Window_Height {
		e.removed = true
	}

	if SDL.HasIntersection(&e.pos, &player.pos) {
		player.health = 0
	}
}

enemy_plane_render :: proc(e: ^Entity) {
	render_entity_sprite(e)
}
