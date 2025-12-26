package main

import SDL "vendor:sdl2"

// Shared player properties
Max_Player_Health :: 100
Player_Max_Bomb_Tick_Time :: 50

Entity_Type :: enum u8 {
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
	next:           ^Entity,
	prev:           ^Entity,
	pos:            SDL.Rect,
	crop:           SDL.Rect,
	xa:             i32,
	ya:             i32,
	texture:        i32,
	type:           Entity_Type,
	removed:        bool,
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

// All entities in game
entities: ^Entity

// Reference to the player in entity list
player: ^Entity

new_entity :: proc(type: Entity_Type) -> ^Entity {
	e := new(Entity)

	e.type = type
	e.tick_time = 0
	e.removed = false

	// insert to the head of entity list
	e.prev = nil
	e.next = entities
	if entities != nil {
		entities.prev = e // not first element?
	}

	entities = e

	return e
}

free_entity :: proc(e: ^Entity) {
	if e == nil || entities == nil {
		return
	}

	if e == entities {
		entities = e.next
	}

	if e.next != nil {
		e.next.prev = e.prev
	}
	if e.prev != nil {
		e.prev.next = e.next
	}

	free(e)
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
	next: ^Entity = nil
	for e := entities; e != nil; e = next {
		next = e.next
		free_entity(e)
	}
}

entity_table := [Entity_Type]Entity_Table_Entry {
	.Player = {tick = player_tick, render = player_render, z_index = 2},
	.Enemy_Plane = {tick = enemy_plane_tick, render = enemy_plane_render, z_index = 2},
	.Bullet = {tick = bullet_tick, render = bullet_render, z_index = 2},
	.Bomb = {tick = bomb_tick, render = bomb_render, z_index = 2},
	.Island = {tick = island_tick, render = island_render, z_index = 0},
	.Explosion = {tick = explosion_tick, render = explosion_render, z_index = 2},
	.Ship = {tick = ship_tick, render = ship_render, z_index = 1},
	.Healer = {tick = healer_tick, render = healer_render, z_index = 2},
}
