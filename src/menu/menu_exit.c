#include "menu.h"

#include "../keyboard_manager/keyboard_manager.h"
#include "../game.h"
#include "../util/util.h"
#include "../gui/gui.h"

void MenuExitTick(void) {
  bool pressed = false;

  if (keys[KeyUp]) {
    exitMenu.selectedIndex--;
    keys[KeyUp] = false;
  }

  if (keys[KeyDown]) {
    exitMenu.selectedIndex++;
    keys[KeyDown] = false;
  }

  if (keys[KeyEnter]) {
    pressed = true;
    keys[KeyEnter] = false;
  }

  const i32 length = 2;

  if (exitMenu.selectedIndex >= length) {
    exitMenu.selectedIndex = 0;
  }

  if (exitMenu.selectedIndex < 0) {
    exitMenu.selectedIndex = length - 1;
  }

  if (pressed) {
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

void MenuExitRender(void) {
  static const char *buttons[] = {
    "YES",
    "NO",
  };

  const i32 length = ArrayLength(buttons);

  RenderString(0,
               0,
               40,
               0xFF,
               0xFF,
               0x00,
               0xFF,
               1,
               (-2) - length + 1,
               "Are you sure you want to exit?",
               buttons[0]);

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
