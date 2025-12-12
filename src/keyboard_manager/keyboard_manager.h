#pragma once

#include "../lib/std.h"

#include "../util/util.h"

// Keys
enum class Key : uint8_t {
  Down,
  Enter,
  Left,
  MusicVolumeDown,
  MusicVolumeUp,
  Pause,
  Right,
  Up,
  Shoot,
  Bomb,
};

inline std::unordered_map<Key, bool> keys;

bool SingleKeyPress(Key key);

void UpdateKey(SDL_Keycode keyCode, bool down);
