package main

import "github.com/veandco/go-sdl2/sdl"

var loseMenu Menu

var loseButtons = []string{"RESTART GAME", "EXIT"}

func menuLoseTick() {
	handleUpDownSelection(&loseMenu, len(loseButtons))

	if singleKeyPress(KeyEnter) {
		switch loseMenu.SelectedIndex {
		case 0:
			// Restart game
			restart()
		case 1:
			// Exit
			menuID = MenuTypeExit
			prevMenuID = MenuTypeLose
		}
	}
}

func menuLoseRender() {
	renderString(0, 0, 40, sdl.Color{R: 255, G: 255, B: 0, A: 255}, true, -3, "YOU LOSE!")
	renderString(0, 0, 40, sdl.Color{R: 255, G: 255, B: 0, A: 255}, true, -2, "TRY AGAIN?")

	for i, button := range loseButtons {
		if loseMenu.SelectedIndex == i {
			renderString(0, 0, 40, sdl.Color{R: 160, G: 160, B: 0, A: 255}, true, i, "> %s <", button)
		} else {
			renderString(0, 0, 40, sdl.Color{R: 255, G: 255, B: 0, A: 255}, true, i, "%s", button)
		}
	}
}

