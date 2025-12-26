package main

import "core:math/rand"

import SDL "vendor:sdl2"

ship_frames := []SDL.Rect{{505, 298, 41, 197}, {463, 298, 41, 197}}

new_ship :: proc() -> ^Entity {
	e := new_entity(.Ship)

	e.texture = 0

	frame := &ship_frames[0]

	e.pos.w = frame.w * 1
	e.pos.h = frame.h * 1
	e.pos.x = rand.int31_max(Window_Width)
	e.pos.y = -rand.int31_max(Window_Height) - e.pos.h

	e.xa = 0
	e.ya = 1

	e.has_shot = false
	e.has_bombed = false
	e.health = 100

	return e
}

ship_tick :: proc(e: ^Entity) {
	e.tick_time += 1
	if e.tick_time > 10 {
		e.tick_time = 0
	}

	if e.health <= 0 {
		new_explosion(e.pos.x, e.pos.y)
		e.removed = true
		return
	}

	e.pos.x += e.xa * 11
	e.pos.y += e.ya * 11

	if e.pos.y >= Window_Height {
		e.removed = true
	}
}

ship_render :: proc(e: ^Entity) {
	frame_no := e.tick_time / 5 % i32(len(ship_frames))
	render_sprite(e.texture, e.pos, ship_frames[frame_no])
}
