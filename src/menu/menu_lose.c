#include "menu.h"

#include "../game_loop/game_loop.h"
#include "../gui/gui.h"
#include "../keyboard_manager/keyboard_manager.h"
#include "../util/util.h"

static const char *buttons[] = {
  "RESTART GAME",
  "EXIT",
};

constexpr i32 length = ArrayLength(buttons);

void MenuLoseTick() {
  HandleUpDownSelection(&loseMenu, length);

  if (SingleKeyPress(KeyEnter)) {
    switch (loseMenu.selectedIndex) {
    case 0: // Restart game
      Restart();
      break;

    case 1: // Exit
      menuID = MenuExit;
      prevMenuID = MenuLose;
      break;
    }
  }
}

void MenuLoseRender() {
  RenderString(0, 0, 40, 0xFF, 0xFF, 0x00, 0xFF, 1, -3, "YOU LOSE!");
  RenderString(0, 0, 40, 0xFF, 0xFF, 0x00, 0xFF, 1, -2, "TRY AGAIN?");

  for (i32 i = 0; i < length; i++) {
    if (loseMenu.selectedIndex == i) {
      RenderString(
        0, 0, 40, 0xA0, 0xA0, 0x00, 0xFF, 1, i, "> %s <", buttons[i]);
    } else {
      RenderString(0, 0, 40, 0xFF, 0xFF, 0x00, 0xFF, 1, i, "%s", buttons[i]);
    }
  }
}
