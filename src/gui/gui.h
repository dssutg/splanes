#pragma once

#include "../util/util.h"

void RenderProgressBar(i32 x,
                       i32 y,
                       i32 width,
                       i32 height,
                       i32 strokeSize,
                       i32 progress);

void RenderString(i32 x,
                  i32 y,
                  i32 size,
                  u8 red,
                  u8 green,
                  u8 blue,
                  u8 alpha,
                  u32 flags,
                  i32 lineNumber,
                  const char *const format,
                  ...);

void RenderHealthBar(i32 health);
void RenderSmallLogo(void);
void FreeFontCache(void);
