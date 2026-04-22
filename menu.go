package main

import "github.com/veandco/go-sdl2/sdl"

type MenuType uint8

const (
	MenuTypeNone MenuType = iota
	MenuTypeMain
	MenuTypeExit
	MenuTypeAbout
	MenuTypeLose
)

var (
	menuNormalTextColor = sdl.Color{R: 255, G: 255, B: 0, A: 255}
	menuHoverTextColor  = sdl.Color{R: 160, G: 160, B: 0, A: 255}
)

const menuFontSize = 40

type Menu struct {
	SelectedIndex int
}

type MenuTableEntry struct {
	Tick   func()
	Render func()
}

var (
	prevMenuID = MenuTypeNone
	menuID     = MenuTypeNone
)

func tickMenu() {
	menuTable[menuID].Tick()
}

func renderMenu() {
	menuTable[menuID].Render()
}

func menuNoneCb() {
	// Nothing should be here.
}

func handleUpDownSelection(menu *Menu, length int) {
	if SingleKeyPress(KeyUp) {
		menu.SelectedIndex--
	}

	if SingleKeyPress(KeyDown) {
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
	MenuTypeNone:  {Tick: menuNoneCb, Render: menuNoneCb},
	MenuTypeMain:  {Tick: menuMainTick, Render: menuMainRender},
	MenuTypeExit:  {Tick: menuExitTick, Render: menuExitRender},
	MenuTypeAbout: {Tick: menuAboutTick, Render: menuAboutRender},
	MenuTypeLose:  {Tick: menuLoseTick, Render: menuLoseRender},
}
