#include "../lib/std.h"

#include "menu.h"

#include "../keyboard_manager/keyboard_manager.h"
#include "../game_loop/game_loop.h"
#include "../gui/gui.h"

enum class Button {
  Yes,
  No,
};

static constexpr auto buttons = std::to_array({
  "YES",
  "NO",
});

constexpr int32_t length = static_cast<int32_t>(buttons.size());

void ExitMenu::tick() {
  handleUpAndDownSelection(length);

  if (SingleKeyPress(Key::Enter)) {
    switch (static_cast<Button>(selectedIndex)) {
    case Button::Yes:
      running = false;
      break;

    case Button::No:
      close();
      break;
    }
  }
}

void ExitMenu::render() {
  const char *title = "Are you sure you want to exit?";

  const SDL_Color titleColor = {.r = 0xFF, .g = 0xFF, .b = 0x00, .a = 0xFF};

  RenderString(0, 0, 40, titleColor, 1, (-2) - length + 1, title, buttons[0]);

  for (int32_t i = 0; i < length; i++) {
    const auto &button = buttons[static_cast<size_t>(i)];

    if (selectedIndex == i) {
      constexpr SDL_Color color = {.r = 0xA0, .g = 0xA0, .b = 0x00, .a = 0xFF};
      RenderString(0, 0, 40, color, 1, i - length + 1, "> %s <", button);
    } else {
      constexpr SDL_Color color = {.r = 0xFF, .g = 0xFF, .b = 0x00, .a = 0xFF};
      RenderString(0, 0, 40, color, 1, i - length + 1, "%s", button);
    }
  }
}
