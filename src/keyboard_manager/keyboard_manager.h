#pragma once

#include "../util/util.h"

#include <SDL2/SDL_keycode.h>

// Keys
typedef enum Key {
  KeyDown,
  KeyEnter,
  KeyLeft,
  KeyMusicVolumeDown,
  KeyMusicVolumeUp,
  KeyPause,
  KeyRight,
  KeyUp,
  KeyShoot,
  KeyBomb,
  KeyCount,
} Key;

extern bool keys[KeyCount];

void UpdateKey(SDL_KeyCode keyCode, bool down);
void InitKeyboardManager(void);
