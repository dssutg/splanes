package main

import SDL "vendor:sdl2"

bullet_frames := []SDL.Rect{{1, 166, 32, 32}, {34, 199, 32, 32}}

new_bullet :: proc(
	type: Entity_Type,
	x, y, xa, ya: i32,
	owner_type: Entity_Type,
	damage, bullet_frame_no: i32,
) -> ^Entity {
	e := new_entity(type)

	e.texture = 0
	e.data = bullet_frame_no

	e.pos.x = x
	e.pos.y = y
	e.pos.w = bullet_frames[bullet_frame_no].w * 2
	e.pos.h = bullet_frames[bullet_frame_no].h * 2

	e.xa = xa
	e.ya = ya

	e.owner_type = owner_type
	e.damage = damage

	return e
}

bullet_tick :: proc(e: ^Entity) {
	e.pos.x += e.xa * 20
	e.pos.y += e.ya * 20

	dead := false

	for other := entities; other != nil; other = other.next {
		can_collide :=
			(e.owner_type == .Player && other.type == .Enemy_Plane) ||
			(e.owner_type == .Enemy_Plane && other.type == .Player)

		if !can_collide {
			continue
		}

		if !SDL.HasIntersection(&other.pos, &e.pos) {
			continue
		}

		dead = true
		hurt_entity(other, e.damage)
		if other.type == .Player {
			new_explosion(e.pos.x, e.pos.y)
			continue
		}

		if e.owner_type == .Player {
			player.score += 1
		}
	}

	window_rect := Window_Rect
	if dead || !SDL.HasIntersection(&e.pos, &window_rect) {
		e.removed = true
	}
}

bullet_render :: proc(e: ^Entity) {
	render_sprite(e.texture, e.pos, bullet_frames[e.data])
}
