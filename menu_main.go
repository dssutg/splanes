package main

import "github.com/veandco/go-sdl2/sdl"

var mainMenu Menu

var mainButtons = []string{"RESUME", "ABOUT", "EXIT"}

func menuMainTick() {
	handleUpDownSelection(&mainMenu, len(mainButtons))

	if SingleKeyPress(KeyEnter) {
		switch mainMenu.SelectedIndex {
		case 0: // Resume
			menuID = MenuTypeNone
		case 1: // About
			menuID = MenuTypeAbout
		case 2: // Exit
			menuID = MenuTypeExit
			prevMenuID = MenuTypeMain
		}
	}
}

func menuMainRender() {
	for i, button := range mainButtons {
		opts := RenderStringOptions{
			Size:                   menuFontSize,
			Color:                  menuNormalTextColor,
			RelativeToWindowCenter: true,
			LineNo:                 i - len(mainButtons) + 1,
		}

		format := "%s"

		if mainMenu.SelectedIndex == i {
			opts.Color = menuHoverTextColor
			format = "> %s <"
		}

		RenderString(opts, format, button)
	}

	crop := sdl.Rect{X: 99, Y: 573, W: 278, H: 141}

	dest := crop
	dest.X = (WindowW - dest.W) / 2
	dest.Y = (WindowH+(int32(menuFontSize)+50)*int32(-1-len(mainButtons)+1))/2 - dest.H

	RenderSprite(TextureMain, dest, crop)
}
