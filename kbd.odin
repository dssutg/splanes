#+feature dynamic-literals
package main

import SDL "vendor:sdl2"

// Keys
Key :: enum u8 {
	Down,
	Enter,
	Left,
	Music_Volume_Down,
	Music_Volume_Up,
	Pause,
	Right,
	Up,
	Shoot,
	Bomb,
}

keys: [Key]bool

// Default key map entries.
key_map := map[SDL.Keycode]Key {
	SDL.Keycode.ESCAPE = .Pause,
	SDL.Keycode.UP     = .Up,
	SDL.Keycode.DOWN   = .Down,
	SDL.Keycode.LEFT   = .Left,
	SDL.Keycode.RIGHT  = .Right,
	SDL.Keycode.k      = .Up,
	SDL.Keycode.j      = .Down,
	SDL.Keycode.h      = .Left,
	SDL.Keycode.l      = .Right,
	SDL.Keycode.w      = .Up,
	SDL.Keycode.s      = .Down,
	SDL.Keycode.a      = .Left,
	SDL.Keycode.d      = .Right,
	SDL.Keycode.F1     = .Music_Volume_Up,
	SDL.Keycode.F2     = .Music_Volume_Down,
	SDL.Keycode.RETURN = .Enter,
	SDL.Keycode.SPACE  = .Shoot,
	SDL.Keycode.x      = .Bomb,
}

update_key :: proc(key_code: SDL.Keycode, down: bool) {
	if key, ok := key_map[key_code]; ok {
		keys[key] = down
	}
}

single_key_press :: proc(key: Key) -> bool {
	pressed := keys[key]
	keys[key] = false
	return pressed
}
