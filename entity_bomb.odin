package main

import "core:math"

import SDL "vendor:sdl2"

bomb_frames := []SDL.Rect{{265, 265, 9, 21}}

new_bomb :: proc() -> ^Entity {
	e := new_entity(.Bomb)

	e.texture = 0
	e.data = 1

	frame := bomb_frames[0]
	e.pos = {0, 0, frame.w * 2, frame.h * 2}

	e.xa = 0
	e.ya = 0

	e.damage = 1000
	e.tick_time = 1

	return e
}

bomb_calc_size :: proc(e: ^Entity) -> (w, h: i32) {
	scale := math.pow(0.90, f64(e.tick_time))

	w = (i32)(f64(e.pos.w) * scale)
	h = (i32)(f64(e.pos.h) * scale)

	if w == 0 || h == 0 {
		e.data = 0
		w = 1
		h = 1
	}

	return
}

bomb_tick :: proc(e: ^Entity) {
	e.tick_time += e.data

	e.pos.x += e.xa * 2
	e.pos.y += e.ya * 2

	if e.data != 0 {
		return
	}

	scaled_rect: SDL.Rect
	scaled_rect.x = e.pos.x
	scaled_rect.y = e.pos.y
	scaled_rect.w, scaled_rect.h = bomb_calc_size(e)

	next_entity: ^Entity = nil
	for it := entities; it != nil; it = next_entity {
		next_entity = it.next
		if it.type == .Ship && SDL.HasIntersection(&it.pos, &scaled_rect) {
			hurt_entity(it, e.damage)
			new_explosion(scaled_rect.x, scaled_rect.y)
			player.score += 1
			break
		}
	}

	e.removed = true
}

bomb_render :: proc(e: ^Entity) {
	w, h := bomb_calc_size(e)
	render_sprite(e.texture, {e.pos.x, e.pos.y, w, h}, bomb_frames[0])
}
