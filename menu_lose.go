package main

var loseMenu Menu

var loseButtons = []string{"RESTART GAME", "EXIT"}

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

func menuLoseRender() {
	RenderString(
		RenderStringOptions{
			Size:                   menuFontSize,
			Color:                  menuNormalTextColor,
			RelativeToWindowCenter: true,
			LineNo:                 -3,
		},
		"YOU LOSE!",
	)

	RenderString(
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

		RenderString(opts, format, button)
	}
}
