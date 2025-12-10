#include "menu.h"

#include "../gui/gui.h"
#include "../keyboard_manager/keyboard_manager.h"
#include "../renderer/renderer.h"
#include "../util/util.h"

void MenuMainTick(void) {
  if (keys[KeyUp]) {
    mainMenu.selectedIndex--;
    keys[KeyUp] = false;
  }

  if (keys[KeyDown]) {
    mainMenu.selectedIndex++;
    keys[KeyDown] = false;
  }

  auto pressed = false;

  if (keys[KeyEnter]) {
    pressed = true;
    keys[KeyEnter] = false;
  }

  constexpr i32 length = 3;

  if (mainMenu.selectedIndex >= length) {
    mainMenu.selectedIndex = 0;
  }

  if (mainMenu.selectedIndex < 0) {
    mainMenu.selectedIndex = length - 1;
  }

  if (pressed) {
    switch (mainMenu.selectedIndex) {
    case 0: // Resume
      menuID = MenuNone;
      break;

    case 1: // About
      menuID = MenuAbout;
      break;

    case 2: // Exit
      menuID = MenuExit;
      prevMenuID = MenuMain;
      break;
    }
  }
}

void MenuMainRender(void) {
  static const char *const buttons[] = {
    "RESUME",
    "ABOUT",
    "EXIT",
  };

  constexpr i32 size = 40;

  constexpr i32 length = ArrayLength(buttons);

  for (i32 i = 0; i < length; i++) {
    if (mainMenu.selectedIndex == i) {
      RenderString(0,
                   0,
                   size,
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
        0, 0, size, 0xFF, 0xFF, 0x00, 0xFF, 1, i - length + 1, "%s", buttons[i]);
    }
  }

  constexpr i32 cropX = 99;
  constexpr i32 cropY = 573;
  constexpr i32 cropWidth = 278;
  constexpr i32 cropHeight = 141;

  constexpr i32 scale = 1;

  constexpr i32 width = cropWidth * scale;
  constexpr i32 height = cropHeight * scale;
  constexpr i32 x = (WindowWidth - width) / 2;
  constexpr i32 y =
    (WindowHeight + (size + 50) * ((-1) - length + 1)) / 2 - height;

  RenderSprite(0, x, y, width, height, cropX, cropY, cropWidth, cropHeight);
}
