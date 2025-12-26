package main

import "core:math/rand"

import SDL "vendor:sdl2"
import SDL_Mixer "vendor:sdl2/mixer"
import SDL_TTF "vendor:sdl2/ttf"

running := true

layer1: i32
layer2: i32

reset :: proc() {
	player = new_player()

	layer1 = -Window_Height
	layer2 = 0

	menu_ID = .Main

	main_menu.selected_index = 0
	exit_menu.selected_index = 0
	about_menu.selected_index = 0
	lose_menu.selected_index = 0

	play_music(MusicBackground0, 70)
}

restart :: proc() {
	remove_all_entities()
	reset()
}

render_layer :: proc(offset_y: i32) {
	tile_width :: (Window_Width + Tile_Size - 1) / Tile_Size
	tile_height :: (Window_Height + Tile_Size - 1) / Tile_Size

	for tile_y in i32(0) ..< tile_height {
		for tile_x in i32(0) ..< tile_width {
			render_sprite(
				0,
				{tile_x * Tile_Size, tile_y * Tile_Size + offset_y, Tile_Size, Tile_Size},
				{265, 364, 32, 32},
			)
		}
	}
}

do_game_loop :: proc() {
	for running {
		// Handle new events
		event: SDL.Event
		for SDL.PollEvent(&event) {
			#partial switch event.type {
			case .QUIT:
				running = false
			case .KEYDOWN:
				update_key(event.key.keysym.sym, true)
			case .KEYUP:
				update_key(event.key.keysym.sym, false)
			}
		}

		// Handle pause and menus
		if keys[.Pause] {
			keys[.Pause] = false
			#partial switch menu_ID {
			case .None:
				menu_ID = .Main
			case .Main:
				menu_ID = .None
			}
		}
		pause := false
		if menu_ID != .None {
			tick_menu()
			pause = true
		}

		// Clear screen
		SDL.SetRenderDrawColor(renderer, 0, 0, 0, 0)
		SDL.RenderClear(renderer)

		if !pause {
			// Spawn new entities
			if rand.int_max(20) == 0 {
				new_enemy_plane()
			}
			if rand.int_max(80) == 0 {
				new_ship()
			}
			if rand.int_max(30) == 0 {
				new_island()
			}
			if rand.int_max(100) == 0 {
				new_healer()
			}

			// Update water layers
			layer1 += 10
			layer2 += 10
			if layer2 >= Window_Height {
				layer2 = layer1 - Window_Height
				layer1, layer2 = layer2, layer1
			}
		}

		// Render water layers
		render_layer(layer1)
		render_layer(layer2)

		// Update and render entitites
		for z_index in 0 ..= 2 {
			for &e in entity_pool {
				entry := entity_table[e.type]
				if entry.z_index == i32(z_index) {
					if !pause {
						entry.tick(&e)
					}
					entry.render(&e)
				}
			}
		}

		// Render GUI
		render_health_bar(player.health)

		render_progress_bar(
			{580, 20, 100, 25},
			5,
			player.bomb_tick_time * 100 / Player_Max_Bomb_Tick_Time,
		)

		render_small_logo()

		render_string(
			300,
			20,
			20,
			{255, 202, 65, 255},
			false,
			0,
			"SCORE: %v, DISTANCE: %v",
			player.score,
			player.distance,
		)

		render_menu()

		SDL.RenderPresent(renderer)
		SDL.Delay(Window_Delay_Milliseconds)
	}
}

main :: proc() {
	if SDL.Init(SDL.INIT_EVERYTHING) != 0 ||
	   SDL_TTF.Init() != 0 ||
	   SDL.Init(SDL.INIT_AUDIO) == -1 ||
	   SDL_Mixer.OpenAudio(44100, SDL_Mixer.DEFAULT_FORMAT, 2, 4096) == -1 {
		fatalf("can't init SDL: %v", SDL.GetError())
	}
	defer SDL.Quit()

	window = SDL.CreateWindow(
		cstring(raw_data(window_title)),
		SDL.WINDOWPOS_CENTERED,
		SDL.WINDOWPOS_CENTERED,
		Window_Width,
		Window_Height,
		SDL.WindowFlags{},
	)
	if window == nil {
		fatalf("can't create window: %v", SDL.GetError())
	}
	defer SDL.DestroyWindow(window)

	renderer = SDL.CreateRenderer(
		window,
		-1,
		SDL.RENDERER_ACCELERATED | SDL.RENDERER_TARGETTEXTURE,
	)
	if renderer == nil {
		fatalf("can't create renderer: %v", SDL.GetError())
	}

	load_textures()
	init_sound_manager()

	reset()

	do_game_loop()
}
