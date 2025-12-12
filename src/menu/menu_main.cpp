#include "../lib/std.h"

#include "menu.h"

#include "../gui/gui.h"
#include "../keyboard_manager/keyboard_manager.h"
#include "../renderer/renderer.h"

enum class Button {
  Resume,
  About,
  Exit,
};

static constexpr auto buttons = std::to_array({
  "RESUME",
  "ABOUT",
  "EXIT",
});

constexpr int32_t length = static_cast<int32_t>(buttons.size());

void MainMenu::tick() {
  handleUpAndDownSelection(length);

  if (SingleKeyPress(Key::Enter)) {
    switch (static_cast<Button>(selectedIndex)) {
    case Button::Resume:
      close();
      break;

    case Button::About:
      openMenu(new AboutMenu());
      break;

    case Button::Exit:
      openMenu(new ExitMenu());
      break;
    }
  }
}

void MainMenu::render() {
  constexpr int32_t size = 40;

  for (int32_t i = 0; i < length; i++) {
    const auto &button = buttons[static_cast<size_t>(i)];

    if (selectedIndex == i) {
      constexpr SDL_Color color = {.r = 0xA0, .g = 0xA0, .b = 0x00, .a = 0xFF};
      RenderString(0, 0, size, color, 1, i - length + 1, "> %s <", button);
    } else {
      constexpr SDL_Color color = {.r = 0xFF, .g = 0xFF, .b = 0x00, .a = 0xFF};
      RenderString(0, 0, size, color, 1, i - length + 1, "%s", button);
    }
  }

  constexpr int32_t cropX = 99;
  constexpr int32_t cropY = 573;
  constexpr int32_t cropWidth = 278;
  constexpr int32_t cropHeight = 141;

  constexpr int32_t scale = 1;

  constexpr int32_t width = cropWidth * scale;
  constexpr int32_t height = cropHeight * scale;
  constexpr int32_t x = (WindowWidth - width) / 2;
  constexpr int32_t y =
    (WindowHeight + (size + 50) * ((-1) - length + 1)) / 2 - height;

  RenderSprite(0,
               {.x = x, .y = y, .w = width, .h = height},
               {.x = cropX, .y = cropY, .w = cropWidth, .h = cropHeight});
}
