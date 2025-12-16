package entity

import SDL "vendor:sdl2"

import "../gfx"
import "../snd"

explosion_frames := []SDL.Rect {
	{67, 166, 32, 32},
	{100, 166, 32, 32},
	{133, 166, 32, 32},
	{166, 166, 32, 32},
	{199, 166, 32, 32},
	{232, 166, 32, 32},
}

new_explosion :: proc(x, y: i32) -> ^Entity {
	e := new_entity(.Explosion)

	e.texture = 0

	frame := explosion_frames[0]
	e.pos = {x, y, frame.w * 2, frame.h * 2}

	snd.play_sound(.Explosion1, 100)

	return e
}

explosion_tick :: proc(e: ^Entity) {
	e.tick_time += 1
	if e.tick_time >= i32(len(explosion_frames)) {
		e.removed = true
	}
}

explosion_render :: proc(e: ^Entity) {
	frame_no := e.tick_time % i32(len(explosion_frames))
	gfx.render_sprite(e.texture, e.pos, explosion_frames[frame_no])
}
