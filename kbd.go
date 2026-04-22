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

var keys = make(map[Key]bool)

var keyMap = map[sdl.Keycode]Key{
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

func updateKey(keyCode sdl.Keycode, down bool) {
	if key, ok := keyMap[keyCode]; ok {
		keys[key] = down
	}
}

func singleKeyPress(key Key) bool {
	pressed := keys[key]
	keys[key] = false
	return pressed
}
