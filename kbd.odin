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
	.ESCAPE = .Pause,
	.UP     = .Up,
	.DOWN   = .Down,
	.LEFT   = .Left,
	.RIGHT  = .Right,
	.k      = .Up,
	.j      = .Down,
	.h      = .Left,
	.l      = .Right,
	.w      = .Up,
	.s      = .Down,
	.a      = .Left,
	.d      = .Right,
	.F1     = .Music_Volume_Up,
	.F2     = .Music_Volume_Down,
	.RETURN = .Enter,
	.SPACE  = .Shoot,
	.x      = .Bomb,
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
