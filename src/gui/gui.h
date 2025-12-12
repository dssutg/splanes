#pragma once

#include "../lib/std.h"

void RenderProgressBar(const SDL_Rect &boundingBox,
                       int32_t strokeSize,
                       int32_t progress);

void RenderString(int32_t x,
                  int32_t y,
                  int32_t size,
                  const SDL_Color &color,
                  bool centerRelativeToWindow,
                  int32_t lineNumber,
                  const char *const format,
                  ...);

void RenderHealthBar(int32_t health);
void RenderSmallLogo();
void FreeFontCache();
