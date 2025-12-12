#include "../lib/std.h"

#include "keyboard_manager.h"

// Default key map entries.
static const std::unordered_map<SDL_Keycode, Key> keyMap = {
  {SDLK_ESCAPE, Key::Pause},
  {SDLK_UP, Key::Up},
  {SDLK_DOWN, Key::Down},
  {SDLK_LEFT, Key::Left},
  {SDLK_RIGHT, Key::Right},
  {SDLK_k, Key::Up},
  {SDLK_j, Key::Down},
  {SDLK_h, Key::Left},
  {SDLK_l, Key::Right},
  {SDLK_w, Key::Up},
  {SDLK_s, Key::Down},
  {SDLK_a, Key::Left},
  {SDLK_d, Key::Right},
  {SDLK_F1, Key::MusicVolumeUp},
  {SDLK_F2, Key::MusicVolumeDown},
  {SDLK_RETURN, Key::Enter},
  {SDLK_SPACE, Key::Shoot},
  {SDLK_x, Key::Bomb},
};

void UpdateKey(SDL_Keycode keyCode, bool down) {
  const auto it = keyMap.find(keyCode);
  if (it != keyMap.end()) {
    keys[it->second] = down;
  }
}

bool SingleKeyPress(Key key) {
  const auto pressed = keys[key];
  keys[key] = false;
  return pressed;
}
