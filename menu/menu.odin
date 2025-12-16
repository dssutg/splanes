#+feature dynamic-literals
package menu

import "../kbd"

// Menus
Menu_Type :: enum u8 {
	None,
	Main,
	Exit,
	About,
	Lose,
}

Menu :: struct {
	selected_index: int,
}

Menu_Table_Entry :: struct {
	tick:   proc(),
	render: proc(),
}

menu_table := map[Menu_Type]Menu_Table_Entry {
	.None = {tick = menu_none_tick, render = menu_none_render},
	.Main = {tick = menu_main_tick, render = menu_main_render},
	.Exit = {tick = menu_exit_tick, render = menu_exit_render},
	.About = {tick = menu_about_tick, render = menu_about_render},
	.Lose = {tick = menu_lose_tick, render = menu_lose_render},
}

prev_menu_ID := Menu_Type.None
menu_ID := Menu_Type.None

tick_menu :: proc() {
	menu_table[menu_ID].tick()
}

render_menu :: proc() {
	menu_table[menu_ID].render()
}

menu_none_tick :: proc() {
	// Nothing should be here.
}

menu_none_render :: proc() {
	// Nothing should be here.
}

handle_up_down_selection :: proc(menu: ^Menu, length: int) {
	if kbd.single_key_press(.Up) {
		menu.selected_index -= 1
	}

	if kbd.single_key_press(.Down) {
		menu.selected_index += 1
	}

	if menu.selected_index >= length {
		menu.selected_index = 0
	}

	if menu.selected_index < 0 {
		menu.selected_index = length - 1
	}
}
