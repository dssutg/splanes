package game_loop

import "core:math/rand"

import SDL "vendor:sdl2"
import SDL_Mixer "vendor:sdl2/mixer"
import SDL_TTF "vendor:sdl2/ttf"

import "../entity"
import "../game_state"
import "../gfx"
import "../kbd"
import "../menu"
import "../snd"
import "../util"

reset :: proc() {
	entity.player = entity.new_player()

	game_state.layer1 = -gfx.Window_Height
	game_state.layer2 = 0

	menu.menu_ID = .Main

	menu.main_menu.selected_index = 0
	menu.exit_menu.selected_index = 0
	menu.about_menu.selected_index = 0
	menu.lose_menu.selected_index = 0

	snd.play_music(snd.Music_ID.Background0, 70)
}

restart :: proc() {
	entity.remove_all_entities()
	reset()
}

tick :: proc() {
	if kbd.keys[.Pause] {
		kbd.keys[.Pause] = false

		#partial switch menu.menu_ID {
		case .None:
			menu.menu_ID = .Main
		case .Main:
			menu.menu_ID = .None
		}
	}

	if menu.menu_ID != .None {
		menu.tick_menu()
		return
	}

	if rand.int_max(20) == 0 {
		entity.new_enemy_plane()
	}
	if rand.int_max(80) == 0 {
		entity.new_ship()
	}
	if rand.int_max(30) == 0 {
		entity.new_island()
	}
	if rand.int_max(100) == 0 {
		entity.new_healer()
	}

	game_state.layer1 += 10
	game_state.layer2 += 10
	if game_state.layer2 >= gfx.Window_Height {
		game_state.layer2 = game_state.layer1 - gfx.Window_Height
		game_state.layer1, game_state.layer2 = game_state.layer2, game_state.layer1
	}

	for e := entity.entities; e != nil; e = e.next {
		entity.entity_table[e.type].tick(e)
	}

	// delete all dead entities
	nextEntity: ^entity.Entity
	for e := entity.entities; e != nil; e = nextEntity {
		nextEntity = e.next
		if e.removed {
			entity.free_entity(e)
		}
	}
}

render_layer :: proc(offset_y: i32) {
	tile_width :: (gfx.Window_Width + gfx.Tile_Size - 1) / gfx.Tile_Size
	tile_height :: (gfx.Window_Height + gfx.Tile_Size - 1) / gfx.Tile_Size

	for tile_y in i32(0) ..< tile_height {
		for tile_x in i32(0) ..< tile_width {
			gfx.render_sprite(
				0,
				{
					tile_x * gfx.Tile_Size,
					tile_y * gfx.Tile_Size + offset_y,
					gfx.Tile_Size,
					gfx.Tile_Size,
				},
				{265, 364, 32, 32},
			)
		}
	}
}

render :: proc() {
	render_layer(game_state.layer1)
	render_layer(game_state.layer2)

	// Render entitites
	for zIndex in i32(0) ..= 2 {
		for e := entity.entities; e != nil; e = e.next {
			entry := &entity.entity_table[e.type]
			if entry.z_index == zIndex {
				entry.render(e)
			}
		}
	}

	// Render GUI
	gfx.render_health_bar(entity.player.health)

	gfx.render_progress_bar(
		{580, 20, 100, 25},
		5,
		entity.player.bomb_tick_time * 100 / entity.Player_Max_Bomb_Tick_Time,
	)

	gfx.render_small_logo()

	gfx.render_string(
		300,
		20,
		20,
		{255, 202, 65, 255},
		false,
		0,
		"SCORE: %v, DISTANCE: %v",
		entity.player.score,
		entity.player.distance,
	)

	menu.render_menu()
}

do_game_loop :: proc() {
	for game_state.running {
		event: SDL.Event

		for SDL.PollEvent(&event) {
			#partial switch event.type {
			case .QUIT:
				game_state.running = false
			case .KEYDOWN:
				kbd.update_key(event.key.keysym.sym, true)
			case .KEYUP:
				kbd.update_key(event.key.keysym.sym, false)
			}
		}

		tick()

		SDL.SetRenderDrawColor(gfx.renderer, 0, 0, 0, 0)
		SDL.RenderClear(gfx.renderer)

		render()

		SDL.RenderPresent(gfx.renderer)
		SDL.Delay(gfx.Window_Delay_Milliseconds)
	}
}

runGame :: proc() {
	if SDL.Init(SDL.INIT_EVERYTHING) != 0 ||
	   SDL_TTF.Init() != 0 ||
	   SDL.Init(SDL.INIT_AUDIO) == -1 ||
	   SDL_Mixer.OpenAudio(44100, SDL_Mixer.DEFAULT_FORMAT, 2, 4096) == -1 {
		util.fatalf("can't init SDL: %v", SDL.GetError())
	}
	defer SDL.Quit()

	gfx.window = SDL.CreateWindow(
		cstring(raw_data(gfx.window_title)),
		SDL.WINDOWPOS_CENTERED,
		SDL.WINDOWPOS_CENTERED,
		gfx.Window_Width,
		gfx.Window_Height,
		SDL.WindowFlags{},
	)
	if gfx.window == nil {
		util.fatalf("can't create window: %v", SDL.GetError())
	}
	defer SDL.DestroyWindow(gfx.window)

	gfx.renderer = SDL.CreateRenderer(
		gfx.window,
		-1,
		SDL.RENDERER_ACCELERATED | SDL.RENDERER_TARGETTEXTURE,
	)
	if gfx.renderer == nil {
		util.fatalf("can't create renderer: %v", SDL.GetError())
	}

	gfx.load_textures()
	snd.init_sound_manager()

	menu.restart = restart

	reset()

	do_game_loop()
}
