package main

import SDL "vendor:sdl2"

main_menu: Menu

main_nuttons :: []string{"RESUME", "ABOUT", "EXIT"}

menu_main_tick :: proc() {
	handle_up_down_selection(&main_menu, len(main_nuttons))

	if single_key_press(.Enter) {
		switch main_menu.selected_index {
		case 0:
			// Resume
			menu_ID = .None
		case 1:
			// About
			menu_ID = .About
		case 2:
			// Exit
			menu_ID = .Exit
			prev_menu_ID = .Main
		}
	}
}

menu_main_render :: proc() {
	size: i32 : 40

	for button, i in main_nuttons {
		if main_menu.selected_index == i {
			render_string(
				0,
				0,
				size,
				{160, 160, 0, 255},
				true,
				i32(i - len(main_nuttons) + 1),
				"> %v <",
				button,
			)
		} else {
			render_string(
				0,
				0,
				size,
				{255, 255, 0, 255},
				true,
				i32(i - len(main_nuttons) + 1),
				"%v",
				button,
			)
		}
	}

	scale: i32 : 1

	crop := SDL.Rect{99, 573, 278, 141}

	dest: SDL.Rect
	dest.w = crop.w * scale
	dest.h = crop.h * scale
	dest.x = (Window_Width - dest.w) / 2
	dest.y = (Window_Height + (size + 50) * (-1 - i32(len(main_nuttons)) + 1)) / 2 - dest.h

	render_sprite(0, dest, crop)
}
