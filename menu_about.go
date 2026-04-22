package main

import "github.com/veandco/go-sdl2/sdl"

var aboutMenu Menu

func menuAboutTick() {
	if singleKeyPress(KeyEnter) {
		menuID = MenuTypeMain
	}
}

func menuAboutRender() {
	lines := []string{
		"Splanes.",
		"",
		"Created by",
		"  Daniil Stepanov",
		"  in November, 2019.",
		"",
		"> BACK <",
	}

	for i, line := range lines {
		color := sdl.Color{R: 255, G: 255, B: 0, A: 255}
		if i == len(lines)-1 {
			color = sdl.Color{R: 160, G: 160, B: 0, A: 255}
		}
		renderString(0, 0, 40, color, true, i-len(lines)+1, "%s", line)
	}
}
