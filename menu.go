package main

import "github.com/veandco/go-sdl2/sdl"

// MenuType identifies the current menu screen.
type MenuType uint8

// Menu type constants.
const (
	MenuTypeNone  MenuType = iota // No menu (gameplay)
	MenuTypeMain                  // Main menu
	MenuTypeExit                  // Exit confirmation
	MenuTypeAbout                 // About/credits
	MenuTypeLose                  // Game over
)

// Menu text colors.
var (
	menuNormalTextColor = sdl.Color{R: 255, G: 255, B: 0, A: 255}
	menuHoverTextColor  = sdl.Color{R: 160, G: 160, B: 0, A: 255}
)

// menuFontSize is the font size for menu text.
const menuFontSize = 40

// Menu holds the state for menu navigation.
type Menu struct {
	SelectedIndex int // Currently selected menu item
}

// MenuTableEntry defines tick and render functions for a menu type.
type MenuTableEntry struct {
	Tick   func() // Menu logic update
	Render func() // Menu rendering
}

// Menu state variables.
var (
	prevMenuID = MenuTypeNone // Previous menu (for back navigation)
	menuID     = MenuTypeNone // Current menu
)

// tickMenu delegates to the current menu's tick function.
func tickMenu() {
	menuTable[menuID].Tick()
}

// renderMenu delegates to the current menu's render function.
func renderMenu() {
	menuTable[menuID].Render()
}

// menuNoneCb is a no-op for menus that don't need updates.
func menuNoneCb() {
	// Nothing should be here.
}

// handleUpDownSelection handles arrow key navigation in menus.
// It wraps around at the top and bottom of the menu.
func handleUpDownSelection(menu *Menu, length int) {
	// Update the index based on user input.
	if SingleKeyPress(KeyUp) {
		menu.SelectedIndex--
	}
	if SingleKeyPress(KeyDown) {
		menu.SelectedIndex++
	}

	// Wrap the selection index.
	if menu.SelectedIndex >= length {
		menu.SelectedIndex = 0
	}
	if menu.SelectedIndex < 0 {
		menu.SelectedIndex = length - 1
	}
}

// menuTable is the dispatch table mapping menu types to their
// tick/render functions.
var menuTable = map[MenuType]MenuTableEntry{
	MenuTypeNone:  {Tick: menuNoneCb, Render: menuNoneCb},
	MenuTypeMain:  {Tick: menuMainTick, Render: menuMainRender},
	MenuTypeExit:  {Tick: menuExitTick, Render: menuExitRender},
	MenuTypeAbout: {Tick: menuAboutTick, Render: menuAboutRender},
	MenuTypeLose:  {Tick: menuLoseTick, Render: menuLoseRender},
}
