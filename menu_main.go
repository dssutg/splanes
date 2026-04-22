package main

import "github.com/veandco/go-sdl2/sdl"

var mainMenu Menu

var mainButtons = []string{"RESUME", "ABOUT", "EXIT"}

func menuMainTick() {
	handleUpDownSelection(&mainMenu, len(mainButtons))

	if singleKeyPress(KeyEnter) {
		switch mainMenu.SelectedIndex {
		case 0:
			// Resume
			menuID = MenuTypeNone
		case 1:
			// About
			menuID = MenuTypeAbout
		case 2:
			// Exit
			menuID = MenuTypeExit
			prevMenuID = MenuTypeMain
		}
	}
}

func menuMainRender() {
	const size = 40

	for i, button := range mainButtons {
		if mainMenu.SelectedIndex == i {
			renderString(
				0,
				0,
				size,
				sdl.Color{R: 160, G: 160, B: 0, A: 255},
				true,
				i-len(mainButtons)+1,
				"> %s <",
				button,
			)
		} else {
			renderString(
				0,
				0,
				size,
				sdl.Color{R: 255, G: 255, B: 0, A: 255},
				true,
				i-len(mainButtons)+1,
				"%s",
				button,
			)
		}
	}

	const scale = 1

	crop := sdl.Rect{X: 99, Y: 573, W: 278, H: 141}

	var dest sdl.Rect
	dest.W = crop.W * scale
	dest.H = crop.H * scale
	dest.X = (WindowWidth - dest.W) / 2
	dest.Y = (WindowHeight+(int32(size)+50)*int32(-1-len(mainButtons)+1))/2 - dest.H

	renderSprite(0, dest, crop)
}
