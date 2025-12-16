package menu

import "../game_state"
import "../gfx"
import "../kbd"

exit_menu: Menu

exit_buttons :: []string{"YES", "NO"}

menu_exit_tick :: proc() {
	handle_up_down_selection(&exit_menu, len(exit_buttons))

	if kbd.single_key_press(.Enter) {
		switch exit_menu.selected_index {
		case 0:
			// Yes
			game_state.running = false
		case 1:
			// No
			menu_ID = prev_menu_ID
		}
	}
}

menu_exit_render :: proc() {
	title :: "Are you sure you want to exit?"

	gfx.render_string(0, 0, 40, {255, 255, 0, 255}, true, i32(-2 - len(exit_buttons) + 1), title)

	for button, i in exit_buttons {
		if exit_menu.selected_index == i {
			gfx.render_string(
				0,
				0,
				40,
				{160, 160, 0, 255},
				true,
				i32(i - len(exit_buttons) + 1),
				"> %v <",
				button,
			)
		} else {
			gfx.render_string(
				0,
				0,
				40,
				{255, 255, 0, 255},
				true,
				i32(i - len(exit_buttons) + 1),
				"%v",
				button,
			)
		}
	}
}
