package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

// Key represents a logical game action, independent of physical key bindings.
type Key uint8

// Logical key constants for game actions.
const (
	KeyDown            Key = iota // Move down
	KeyEnter                      // Confirm selection
	KeyLeft                       // Move left
	KeyMusicVolumeDown            // Decrease music volume
	KeyMusicVolumeUp              // Increase music volume
	KeyPause                      // Pause/unpause game
	KeyRight                      // Move right
	KeyUp                         // Move up
	KeyShoot                      // Fire bullet
	KeyBomb                       // Drop bomb
	KeyRotateLeft                 // Rotate left
	KeyRotateRight                // Rotate right
)

// Keys is the current state of all logical keys.
// true = pressed, false = released.
var Keys = make(map[Key]bool)

// KeyMap maps SDL keycodes to logical game keys.
// Allows multiple physical keys to map to the same action.
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
	sdl.K_q:      KeyRotateLeft,
	sdl.K_e:      KeyRotateRight,
}

// UpdateKey sets the state of a logical key based on SDL key events.
func UpdateKey(keyCode sdl.Keycode, down bool) {
	if key, ok := KeyMap[keyCode]; ok {
		Keys[key] = down
	}
}

// SingleKeyPress returns true if a key was pressed this frame,
// then automatically clears it (for one-shot key actions).
func SingleKeyPress(key Key) bool {
	pressed := Keys[key]
	Keys[key] = false
	return pressed
}
