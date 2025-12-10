#include "menu.h"

#include "../keyboard_manager/keyboard_manager.h"
#include "../util/util.h"
#include "../gui/gui.h"

void MenuAboutTick(void) {
  auto pressed = false;

  if (keys[KeyEnter]) {
    keys[KeyEnter] = 0;
    pressed = true;
  }

  if (pressed) {
    menuID = MenuMain;
  }
}

void MenuAboutRender(void) {
  static const char *const lines[] = {
    "Splanes.",
    "",
    "Created by",
    "  Daniil Stepanov",
    "  in November, 2019.",
    "",
    "> BACK <",
  };

  constexpr i32 length = ArrayLength(lines);

  for (i32 i = 0; i < length; i++) {
    u8 red = 0xFF;
    u8 green = 0xFF;
    u8 blue = 0x00;

    if (i == length - 1) {
      red = 0xA0;
      green = 0xA0;
      blue = 0x00;
    }

    RenderString(
      0, 0, 40, red, green, blue, 0xFF, 1, i - length + 1, "%s", lines[i]);
  }
}
