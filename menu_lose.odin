package main

lose_menu: Menu

lose_buttons :: []string{"RESTART GAME", "EXIT"}

menu_lose_tick :: proc() {
	handle_up_down_selection(&lose_menu, len(lose_buttons))

	if single_key_press(.Enter) {
		switch lose_menu.selected_index {
		case 0:
			// Restart game
			restart()
		case 1:
			// Exit
			menu_ID = .Exit
			prev_menu_ID = .Lose
		}
	}
}

menu_lose_render :: proc() {
	render_string(0, 0, 40, {255, 255, 0, 255}, true, -3, "YOU LOSE!")
	render_string(0, 0, 40, {255, 255, 0, 255}, true, -2, "TRY AGAIN?")

	for button, i in lose_buttons {
		if lose_menu.selected_index == i {
			render_string(0, 0, 40, {160, 160, 0, 255}, true, i32(i), "> %v <", button)
		} else {
			render_string(0, 0, 40, {255, 255, 0, 255}, true, i32(i), "%v", button)
		}
	}
}
