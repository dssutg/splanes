package entity

import "core:math/rand"

import SDL "vendor:sdl2"

import "../gfx"

healer_frames := []SDL.Rect{{166, 265, 29, 15}}

new_healer :: proc() -> ^Entity {
	e := new_entity(.Healer)

	e.texture = 0
	e.data = 0

	e.pos.w = healer_frames[e.data].w * 2
	e.pos.h = healer_frames[e.data].h * 2
	e.pos.x = rand.int31_max(gfx.Window_Width)
	e.pos.y = -rand.int31_max(gfx.Window_Height) - e.pos.h

	e.xa = 0
	e.ya = 1

	return e
}

healer_tick :: proc(e: ^Entity) {
	e.pos.x += e.xa * 10
	e.pos.y += e.ya * 10

	if e.pos.y >= gfx.Window_Height {
		e.removed = true
	}

	if SDL.HasIntersection(&e.pos, &player.pos) {
		heal_player(20)
		e.removed = true
	}
}

healer_render :: proc(e: ^Entity) {
	gfx.render_sprite(e.texture, e.pos, healer_frames[e.data])
}
