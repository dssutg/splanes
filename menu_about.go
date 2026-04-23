package main

// aboutMenu is the about screen navigation state (for future use).
var aboutMenu Menu

// menuAboutTick handles input for the about screen.
// Returns to the previous menu when Enter is pressed.
func menuAboutTick() {
	if SingleKeyPress(KeyEnter) {
		menuID = MenuTypeMain
	}
}

// menuAboutRender draws the credits/about screen.
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

	// Render each line.
	for i, line := range lines {
		opts := RenderStringOptions{
			Size:                   menuFontSize,
			Color:                  menuNormalTextColor,
			RelativeToWindowCenter: true,
			LineNo:                 i - len(lines) + 1,
		}

		if i == len(lines)-1 {
			opts.Color = menuHoverTextColor
		}

		RenderStringf(opts, "%s", line)
	}
}
