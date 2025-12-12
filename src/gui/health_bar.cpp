#include "gui.h"

#include "../renderer/renderer.h"

void RenderHealthBar(int32_t health) {
  constexpr int32_t scale = 2;

  constexpr int32_t x = 20;
  constexpr int32_t y = 20;

  int32_t cropX = 364;
  constexpr int32_t cropY = 199;
  int32_t cropWidth = 130;
  constexpr int32_t cropHeight = 14;

  RenderSprite(
    0,
    {.x = x, .y = y, .w = cropWidth * scale, .h = cropHeight * scale},
    {.x = cropX, .y = cropY, .w = cropWidth, .h = cropHeight});

  cropX = 494;
  cropWidth = 7;

  if (health < 0) {
    return;
  }

  constexpr int32_t maxElements = 18;

  auto elementCount = health * maxElements / 100;
  if (elementCount > maxElements) {
    elementCount = maxElements;
  }

  for (int32_t i = 0; i < elementCount; i++) {
    const SDL_Rect dest = {.x = x + i * cropWidth * scale,
                           .y = y,
                           .w = cropWidth * scale,
                           .h = cropHeight * scale};
    const SDL_Rect src = {
      .x = cropX, .y = cropY, .w = cropWidth, .h = cropHeight};
    RenderSprite(0, dest, src);
  }
}
