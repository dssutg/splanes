package main

// exitMenu is the exit confirmation menu navigation state.
var exitMenu Menu

// exitButtons are the confirmation options.
var exitButtons = []string{"YES", "NO"}

// menuExitTick handles input for the exit confirmation menu.
func menuExitTick() {
	handleUpDownSelection(&exitMenu, len(exitButtons))

	if SingleKeyPress(KeyEnter) {
		switch exitMenu.SelectedIndex {
		case 0: // Yes
			running = false
		case 1: // No
			menuID = prevMenuID
		}
	}
}

// menuExitRender draws the exit confirmation prompt.
func menuExitRender() {
	RenderStringf(
		RenderStringOptions{
			Size:                   menuFontSize,
			Color:                  menuNormalTextColor,
			RelativeToWindowCenter: true,
			LineNo:                 -2 - len(exitButtons) + 1,
		},
		"Are you sure you want to exit?",
	)

	for i, button := range exitButtons {
		color := menuNormalTextColor
		format := "%s"

		if exitMenu.SelectedIndex == i {
			color = menuHoverTextColor
			format = "> %s <"
		}

		RenderStringf(
			RenderStringOptions{
				Size:                   menuFontSize,
				Color:                  color,
				RelativeToWindowCenter: true,
				LineNo:                 i - len(exitButtons) + 1,
			},
			format,
			button,
		)
	}
}
