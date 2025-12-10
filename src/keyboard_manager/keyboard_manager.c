#include <SDL2/SDL.h>

#include "keyboard_manager.h"

#include "../util/util.h"

bool keys[KeyCount];

typedef struct KeyMapEntry {
  const SDL_KeyCode keyCode;
  const Key key;
} KeyMapEntry;

// Default key map entries.
static KeyMapEntry keyMap[] = {
  {SDLK_ESCAPE, KeyPause},
  {SDLK_UP, KeyUp},
  {SDLK_DOWN, KeyDown},
  {SDLK_LEFT, KeyLeft},
  {SDLK_RIGHT, KeyRight},
  {SDLK_k, KeyUp},
  {SDLK_j, KeyDown},
  {SDLK_h, KeyLeft},
  {SDLK_l, KeyRight},
  {SDLK_w, KeyUp},
  {SDLK_s, KeyDown},
  {SDLK_a, KeyLeft},
  {SDLK_d, KeyRight},
  {SDLK_F1, KeyMusicVolumeUp},
  {SDLK_F2, KeyMusicVolumeDown},
  {SDLK_RETURN, KeyEnter},
  {SDLK_SPACE, KeyShoot},
  {SDLK_x, KeyBomb},
};

static int CompareKeyMapEntries(const void *a, const void *b) {
  const KeyMapEntry *const aEntry = a;
  const KeyMapEntry *const bEntry = b;

  if (aEntry->keyCode < bEntry->keyCode) {
    return -1;
  }

  if (aEntry->keyCode > bEntry->keyCode) {
    return 1;
  }

  return 0;
}

void InitKeyboardManager(void) {
  qsort(keyMap, ArrayLength(keyMap), sizeof(keyMap[0]), CompareKeyMapEntries);
}

void UpdateKey(SDL_KeyCode keyCode, bool down) {
  const KeyMapEntry entryKey = {.keyCode = keyCode};

  const KeyMapEntry *const entry = bsearch(&entryKey,
                                           keyMap,
                                           ArrayLength(keyMap),
                                           sizeof(keyMap[0]),
                                           CompareKeyMapEntries);
  if (entry == nullptr) {
    return;
  }

  keys[entry->key] = down;
}
