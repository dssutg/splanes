#include "menu.h"

#include "../keyboard_manager/keyboard_manager.h"
#include "../game_loop/game_loop.h"
#include "../util/util.h"
#include "../gui/gui.h"

static const char *buttons[] = {
  "YES",
  "NO",
};

constexpr i32 length = ArrayLength(buttons);

void MenuExitTick() {
  HandleUpDownSelection(&exitMenu, length);

  if (SingleKeyPress(KeyEnter)) {
    switch (exitMenu.selectedIndex) {
    case 0: // Yes
      running = false;
      break;

    case 1: // No
      menuID = prevMenuID;
      break;
    }
  }
}

void MenuExitRender() {
  const char *title = "Are you sure you want to exit?";

  RenderString(
    0, 0, 40, 0xFF, 0xFF, 0x00, 0xFF, 1, (-2) - length + 1, title, buttons[0]);

  for (i32 i = 0; i < length; i++) {
    if (exitMenu.selectedIndex == i) {
      RenderString(0,
                   0,
                   40,
                   0xA0,
                   0xA0,
                   0x00,
                   0xFF,
                   1,
                   i - length + 1,
                   "> %s <",
                   buttons[i]);
    } else {
      RenderString(
        0, 0, 40, 0xFF, 0xFF, 0x00, 0xFF, 1, i - length + 1, "%s", buttons[i]);
    }
  }
}
