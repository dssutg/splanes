#include "../lib/std.h"

#include "gui.h"

#include "../renderer/renderer.h"

void RenderSmallLogo() {
  constexpr int32_t cropWidth = 115;
  constexpr int32_t cropHeight = 58;

  constexpr int32_t scale = 1;

  constexpr auto width = cropWidth * scale;
  constexpr auto height = cropHeight * scale;

  RenderSprite(0,
               {.x = WindowWidth - width, .y = 0, .w = width, .h = height},
               {.x = 700, .y = 401, .w = cropWidth, .h = cropHeight});
}
