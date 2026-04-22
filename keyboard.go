package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

// Keys
type Key uint8

const (
	KeyDown Key = iota
	KeyEnter
	KeyLeft
	KeyMusicVolumeDown
	KeyMusicVolumeUp
	KeyPause
	KeyRight
	KeyUp
	KeyShoot
	KeyBomb
)

var Keys = make(map[Key]bool)

var KeyMap = map[sdl.Keycode]Key{
	sdl.K_ESCAPE: KeyPause,
	sdl.K_UP:     KeyUp,
	sdl.K_DOWN:   KeyDown,
	sdl.K_LEFT:   KeyLeft,
	sdl.K_RIGHT:  KeyRight,
	sdl.K_k:      KeyUp,
	sdl.K_j:      KeyDown,
	sdl.K_h:      KeyLeft,
	sdl.K_l:      KeyRight,
	sdl.K_w:      KeyUp,
	sdl.K_s:      KeyDown,
	sdl.K_a:      KeyLeft,
	sdl.K_d:      KeyRight,
	sdl.K_F1:     KeyMusicVolumeUp,
	sdl.K_F2:     KeyMusicVolumeDown,
	sdl.K_RETURN: KeyEnter,
	sdl.K_SPACE:  KeyShoot,
	sdl.K_x:      KeyBomb,
}

func UpdateKey(keyCode sdl.Keycode, down bool) {
	if key, ok := KeyMap[keyCode]; ok {
		Keys[key] = down
	}
}

func SingleKeyPress(key Key) bool {
	pressed := Keys[key]
	Keys[key] = false
	return pressed
}
