package main

// loseMenu is the game over menu navigation state.
var loseMenu Menu

// loseButtons are the menu options after losing.
var loseButtons = []string{"RESTART GAME", "EXIT"}

// menuLoseTick handles input for the game over screen.
func menuLoseTick() {
	handleUpDownSelection(&loseMenu, len(loseButtons))

	if SingleKeyPress(KeyEnter) {
		switch loseMenu.SelectedIndex {
		case 0: // Restart game
			restart()
		case 1: // Exit
			menuID = MenuTypeExit
			prevMenuID = MenuTypeLose
		}
	}
}

// menuLoseRender draws the game over screen.
func menuLoseRender() {
	RenderStringf(
		RenderStringOptions{
			Size:                   menuFontSize,
			Color:                  menuNormalTextColor,
			RelativeToWindowCenter: true,
			LineNo:                 -3,
		},
		"YOU LOSE!",
	)

	RenderStringf(
		RenderStringOptions{
			Size:                   menuFontSize,
			Color:                  menuNormalTextColor,
			RelativeToWindowCenter: true,
			LineNo:                 -2,
		},
		"TRY AGAIN?",
	)

	for i, button := range loseButtons {
		opts := RenderStringOptions{
			Size:                   menuFontSize,
			Color:                  menuNormalTextColor,
			RelativeToWindowCenter: true,
			LineNo:                 i,
		}

		format := "%s"

		if loseMenu.SelectedIndex == i {
			opts.Color = menuHoverTextColor
			format = "> %s <"
		}

		RenderStringf(opts, format, button)
	}
}
