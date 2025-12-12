#include "../lib/std.h"

#include "menu.h"

#include "../game_loop/game_loop.h"
#include "../gui/gui.h"
#include "../keyboard_manager/keyboard_manager.h"

enum class Button {
  RestartGame,
  Exit,
};

static constexpr auto buttons = std::to_array({
  "RESTART GAME",
  "EXIT",
});

constexpr int32_t length = static_cast<int32_t>(buttons.size());

void LoseMenu::tick() {
  handleUpAndDownSelection(length);

  if (SingleKeyPress(Key::Enter)) {
    switch (static_cast<Button>(selectedIndex)) {
    case Button::RestartGame:
      close();
      Restart();
      break;

    case Button::Exit:
      openMenu(new ExitMenu());
      break;
    }
  }
}

void LoseMenu::render() {
  constexpr SDL_Color titleColor = {.r = 0xFF, .g = 0xFF, .b = 0x00, .a = 0xFF};

  RenderString(0, 0, 40, titleColor, 1, -3, "YOU LOSE!");
  RenderString(0, 0, 40, titleColor, 1, -2, "TRY AGAIN?");

  for (int32_t i = 0; i < length; i++) {
    const auto &button = buttons[static_cast<size_t>(i)];

    if (selectedIndex == i) {
      constexpr SDL_Color color = {.r = 0xA0, .g = 0xA0, .b = 0x00, .a = 0xFF};
      RenderString(0, 0, 40, color, 1, i, "> %s <", button);
    } else {
      constexpr SDL_Color color = {.r = 0xFF, .g = 0xFF, .b = 0x00, .a = 0xFF};
      RenderString(0, 0, 40, color, 1, i, "%s", button);
    }
  }
}
