#include "../lib/std.h"

#include "gui.h"

#include "../renderer/renderer.h"

void RenderProgressBar(const SDL_Rect &boundingBox,
                       int32_t strokeSize,
                       int32_t progress) {
  const auto [x, y, w, h] = boundingBox;

  constexpr SDL_Color borderColor = {.r = 0, .g = 0, .b = 0, .a = 0};

  // Render top border
  RenderRect({.x = x, .y = y, .w = w, .h = strokeSize}, borderColor);

  // Render bottom border
  RenderRect({.x = x, .y = y + h - strokeSize - 1, .w = w, .h = strokeSize},
             borderColor);

  // Render left border
  RenderRect({.x = x, .y = y, .w = strokeSize, .h = h}, borderColor);

  // Render right border
  RenderRect({.x = x + w - strokeSize - 1, .y = y, .w = strokeSize, .h = h},
             borderColor);

  // Render bar
  const auto fullBarWidth = w - strokeSize * 2;

  const SDL_Rect barRect = {
    .x = x + strokeSize,
    .y = y + strokeSize,
    .w = fullBarWidth * progress / 100,
    .h = h - strokeSize * 2,
  };

  SDL_Color barColor = {.r = 0xFF, .g = 0x00, .b = 0x00, .a = 0xFF};
  if (progress > 75) {
    barColor = {.r = 0x00, .g = 0xFF, .b = 0x00, .a = 0xFF};
  } else if (progress > 30) {
    barColor = {.r = 0xFF, .g = 0xFF, .b = 0x00, .a = 0xFF};
  }

  RenderRect(barRect, barColor);
}
