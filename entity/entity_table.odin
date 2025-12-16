#+feature dynamic-literals
package entity

entity_table := map[Entity_Type]Entity_Table_Entry {
	.Player = {tick = player_tick, render = player_render, z_index = 2},
	.Enemy_Plane = {tick = enemy_plane_tick, render = enemy_plane_render, z_index = 2},
	.Bullet = {tick = bullet_tick, render = bullet_render, z_index = 2},
	.Bomb = {tick = bomb_tick, render = bomb_render, z_index = 2},
	.Island = {tick = island_tick, render = island_render, z_index = 0},
	.Explosion = {tick = explosion_tick, render = explosion_render, z_index = 2},
	.Ship = {tick = ship_tick, render = ship_render, z_index = 1},
	.Healer = {tick = healer_tick, render = healer_render, z_index = 2},
}
