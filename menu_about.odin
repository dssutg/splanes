package main

import SDL "vendor:sdl2"

about_menu: Menu

menu_about_tick :: proc() {
	if single_key_press(.Enter) {
		menu_ID = .Main
	}
}

menu_about_render :: proc() {
	lines :: []string {
		"Splanes.",
		"",
		"Created by",
		"  Daniil Stepanov",
		"  in November, 2019.",
		"",
		"> BACK <",
	}

	for line, i in lines {
		color := SDL.Color{255, 255, 0, 255}
		if i == len(lines) - 1 {
			color = {160, 160, 0, 255}
		}
		render_string(0, 0, 40, color, true, i32(i - len(lines) + 1), "%v", line)
	}
}
