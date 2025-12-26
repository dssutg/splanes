package main

import "core:math/rand"

import SDL "vendor:sdl2"

island_frames := []SDL.Rect{{100, 496, 64, 65}, {165, 496, 64, 65}, {230, 496, 64, 65}}

new_island :: proc() -> ^Entity {
	e := new_entity(.Island)

	e.texture = 0
	e.data = rand.int31_max(3)

	e.pos.w = island_frames[e.data].w * 3
	e.pos.h = island_frames[e.data].h * 3
	e.pos.x = rand.int31_max(Window_Width)
	e.pos.y = -rand.int31_max(Window_Height) - e.pos.h

	e.xa = 0
	e.ya = 1

	return e
}

island_tick :: proc(e: ^Entity) {
	e.pos.x += e.xa * 10
	e.pos.y += e.ya * 10
	if e.pos.y >= Window_Height {
		remove_entity(e)
	}
}

island_render :: proc(e: ^Entity) {
	render_sprite(e.texture, e.pos, island_frames[e.data])
}
