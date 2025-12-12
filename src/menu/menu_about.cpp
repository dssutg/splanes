#include "../lib/std.h"

#include "menu.h"

#include "../keyboard_manager/keyboard_manager.h"
#include "../gui/gui.h"

static constexpr auto lines = std::to_array({
  "Splanes.",
  "",
  "Created by",
  "  Daniil Stepanov",
  "  in November, 2019.",
  "",
  "> BACK <",
});

constexpr int32_t length = static_cast<int32_t>(lines.size());

void AboutMenu::tick() {
  if (SingleKeyPress(Key::Enter)) {
    close();
  }
}

void AboutMenu::render() {
  constexpr SDL_Color mainColor = {.r = 0xFF, .g = 0xFF, .b = 0x00, .a = 0xFF};
  constexpr SDL_Color lastColor = {.r = 0xA0, .g = 0xA0, .b = 0x00, .a = 0xFF};

  for (int32_t i = 0; i < length; i++) {
    const auto &line = lines[static_cast<size_t>(i)];
    const auto &color = i == length - 1 ? lastColor : mainColor;
    RenderString(0, 0, 40, color, 1, i - length + 1, "%s", line);
  }
}
