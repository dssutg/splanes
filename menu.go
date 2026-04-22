package main

type MenuType uint8

const (
	MenuTypeNone MenuType = iota
	MenuTypeMain
	MenuTypeExit
	MenuTypeAbout
	MenuTypeLose
)

type Menu struct {
	SelectedIndex int
}

type MenuTableEntry struct {
	Tick   func()
	Render func()
}

var prevMenuID = MenuTypeNone
var menuID = MenuTypeNone

func tickMenu() {
	menuTable[menuID].Tick()
}

func renderMenu() {
	menuTable[menuID].Render()
}

func menuNoneTick() {
	// Nothing should be here.
}

func menuNoneRender() {
	// Nothing should be here.
}

func handleUpDownSelection(menu *Menu, length int) {
	if singleKeyPress(KeyUp) {
		menu.SelectedIndex--
	}

	if singleKeyPress(KeyDown) {
		menu.SelectedIndex++
	}

	if menu.SelectedIndex >= length {
		menu.SelectedIndex = 0
	}

	if menu.SelectedIndex < 0 {
		menu.SelectedIndex = length - 1
	}
}

var menuTable = map[MenuType]MenuTableEntry{
	MenuTypeNone: {Tick: menuNoneTick, Render: menuNoneRender},
	MenuTypeMain: {Tick: menuMainTick, Render: menuMainRender},
	MenuTypeExit: {Tick: menuExitTick, Render: menuExitRender},
	MenuTypeAbout: {Tick: menuAboutTick, Render: menuAboutRender},
	MenuTypeLose: {Tick: menuLoseTick, Render: menuLoseRender},
}