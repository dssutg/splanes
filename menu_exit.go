package main

var exitMenu Menu

var exitButtons = []string{"YES", "NO"}

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

func menuExitRender() {
	title := "Are you sure you want to exit?"

	RenderString(
		RenderStringOptions{
			Size:                   menuFontSize,
			Color:                  menuNormalTextColor,
			RelativeToWindowCenter: true,
			LineNo:                 -2 - len(exitButtons) + 1,
		},
		title,
	)

	for i, button := range exitButtons {
		color := menuNormalTextColor
		format := "%s"

		if exitMenu.SelectedIndex == i {
			color = menuHoverTextColor
			format = "> %s <"
		}

		RenderString(
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
