package main

import "github.com/veandco/go-sdl2/sdl"

var exitMenu Menu

var exitButtons = []string{"YES", "NO"}

func menuExitTick() {
	handleUpDownSelection(&exitMenu, len(exitButtons))

	if singleKeyPress(KeyEnter) {
		switch exitMenu.SelectedIndex {
		case 0:
			// Yes
			running = false
		case 1:
			// No
			menuID = prevMenuID
		}
	}
}

func menuExitRender() {
	title := "Are you sure you want to exit?"

	renderString(0, 0, 40, sdl.Color{R: 255, G: 255, B: 0, A: 255}, true, -2-len(exitButtons)+1, title)

	for i, button := range exitButtons {
		if exitMenu.SelectedIndex == i {
			renderString(
				0,
				0,
				40,
				sdl.Color{R: 160, G: 160, B: 0, A: 255},
				true,
				i-len(exitButtons)+1,
				"> %s <",
				button,
			)
		} else {
			renderString(
				0,
				0,
				40,
				sdl.Color{R: 255, G: 255, B: 0, A: 255},
				true,
				i-len(exitButtons)+1,
				"%s",
				button,
			)
		}
	}
}

