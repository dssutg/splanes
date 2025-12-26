package main

import SDL "vendor:sdl2"

// Shared player properties
Max_Player_Health :: 100
Player_Max_Bomb_Tick_Time :: 50

Entity_Type :: enum u8 {
	None = 0, // always zero
	Player,
	Enemy_Plane,
	Bullet,
	Bomb,
	Island,
	Explosion,
	Ship,
	Healer,
}

// Entity fat struct
Entity :: struct {
	// Common
	pos:            SDL.Rect,
	crop:           SDL.Rect,
	xa:             i32,
	ya:             i32,
	texture:        i32,
	type:           Entity_Type,
	health:         i32,
	damage:         i32,
	data:           i32,
	tick_time:      i32,

	// EntityBullet
	owner_type:     Entity_Type, // entity type that spawned this bullet

	// EntityPlayer
	score:          u64,
	distance:       u64,
	has_shot:       bool,
	has_bombed:     bool,
	death_time:     i32,
	bomb_tick_time: i32,
}

// Entity polymorphic dispatch table entry
Entity_Table_Entry :: struct {
	tick:    proc(e: ^Entity),
	render:  proc(e: ^Entity),
	z_index: i32,
}

entity_pool: [100]Entity // All entities in game
entity_stub: Entity // Dummy entity returned when entity could not be created

// Reference to the player in entity list
player: ^Entity

new_entity :: proc(type: Entity_Type) -> ^Entity {
	entity := &entity_stub // default to stub
	for &e in entity_pool {
		if e.type == .None {
			entity = &e
			break
		}
	}
	entity^ = {} // zero entity
	entity.type = type
	return entity
}

remove_entity :: proc(e: ^Entity) {
	e^ = {} // zero entity
}

hurt_entity :: proc(e: ^Entity, damage: i32) {
	if e.type != .Player && e.type != .Enemy_Plane && e.type != .Ship {
		return
	}

	if e.health < damage {
		e.health = 0
	} else {
		e.health -= damage
	}

	play_sound(SoundHurt, 100)
}

render_entity_sprite :: proc(e: ^Entity) {
	render_sprite(e.texture, e.pos, e.crop)
}

remove_all_entities :: proc() {
	entity_pool = {}
}

// dummy entity handler that does nothing
entity_none_cb :: proc(_: ^Entity) {}

entity_table := [Entity_Type]Entity_Table_Entry {
	.None = {tick = entity_none_cb, render = entity_none_cb, z_index = 0},
	.Player = {tick = player_tick, render = player_render, z_index = 2},
	.Enemy_Plane = {tick = enemy_plane_tick, render = enemy_plane_render, z_index = 2},
	.Bullet = {tick = bullet_tick, render = bullet_render, z_index = 2},
	.Bomb = {tick = bomb_tick, render = bomb_render, z_index = 2},
	.Island = {tick = island_tick, render = island_render, z_index = 0},
	.Explosion = {tick = explosion_tick, render = explosion_render, z_index = 2},
	.Ship = {tick = ship_tick, render = ship_render, z_index = 1},
	.Healer = {tick = healer_tick, render = healer_render, z_index = 2},
}
