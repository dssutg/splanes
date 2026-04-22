package main

var aboutMenu Menu

func menuAboutTick() {
	if SingleKeyPress(KeyEnter) {
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
		opts := RenderStringOptions{
			Size:                   menuFontSize,
			Color:                  menuNormalTextColor,
			RelativeToWindowCenter: true,
			LineNo:                 i - len(lines) + 1,
		}

		if i == len(lines)-1 {
			opts.Color = menuHoverTextColor
		}

		RenderString(opts, "%s", line)
	}
}
