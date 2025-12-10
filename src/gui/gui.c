#include <stdarg.h>

#include <SDL2/SDL_ttf.h>

#include "gui.h"

#include "../util/util.h"
#include "../renderer/renderer.h"

typedef struct FontCacheEntry {
  TTF_Font *font;
  i32 size;
} FontCacheEntry;

static FontCacheEntry *fontCache;
static i64 fontCacheLength;
static i64 fontCacheCapacity;

TTF_Font *LoadFont(i32 size) {
  // Find the font with the required size in the loaded font cache.
  for (i64 i = 0; i < fontCacheLength; i++) {
    const auto entry = &fontCache[i];
    if (entry->size == size) {
      return entry->font;
    }
  }

  // Font not found. Ensure there is enough capacity to add it to the cache.
  if (fontCacheLength >= fontCacheCapacity) {
    if (fontCacheCapacity == 0) {
      fontCacheCapacity = 8;
    } else {
      fontCacheCapacity *= 2;
    }
    fontCache = Erealloc(fontCache, fontCacheCapacity * sizeof(fontCache[0]));
  }

  // Reserve the new font entry.
  auto entry = &fontCache[fontCacheLength];
  fontCacheLength++;

  // Open the font.
  auto font = TTF_OpenFont("assets/fonts/OpenSans-Bold.ttf", size);
  if (font == nullptr) {
    Fatalf("%s", SDL_GetError());
  }

  // Add the new font to the font cache.
  entry->font = font;
  entry->size = size;

  return font;
}

void RenderProgressBar(i32 x,
                       i32 y,
                       i32 width,
                       i32 height,
                       i32 strokeSize,
                       i32 progress) {
  RenderRect(x, y, width, strokeSize, 0, 0, 0);
  RenderRect(x, y + height - strokeSize - 1, width, strokeSize, 0, 0, 0);
  RenderRect(x, y, strokeSize, height, 0, 0, 0);
  RenderRect(x + width - strokeSize - 1, y, strokeSize, height, 0, 0, 0);

  const auto cropX = x + strokeSize;
  const auto cropY = y + strokeSize;
  const auto cropWidth = width - strokeSize * 2;
  const auto cropHeight = height - strokeSize * 2;

  u8 red = 0xFF;
  u8 green = 0x00;
  u8 blue = 0x00;

  if (progress > 75) {
    red = 0x00;
    green = 0xFF;
    blue = 0x00;
  } else if (progress > 30) {
    red = 0xFF;
    green = 0xFF;
    blue = 0x00;
  }

  RenderRect(
    cropX, cropY, cropWidth * progress / 100, cropHeight, red, green, blue);
}

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
                  ...) {
  char text[4096];
  va_list args;
  va_start(args, format);
  vsnprintf(text, sizeof(text), format, args);
  va_end(args);

  if (text[0] == '\0') {
    return;
  }

  const SDL_Color color = {
    .r = red,
    .g = green,
    .b = blue,
    .a = alpha,
  };

  const auto font = LoadFont(size);

  auto fontSurface = TTF_RenderText_Blended(font, text, color);
  if (fontSurface == nullptr) {
    Fatalf("%s", SDL_GetError());
  }

  auto texture = SDL_CreateTextureFromSurface(renderer, fontSurface);
  if (flags == 1) {
    x = (WindowWidth - fontSurface->w) / 2;
    y = (WindowHeight + (fontSurface->h + 50) * lineNumber) / 2;
  }

  const SDL_Rect source = {
    .x = 0,
    .y = 0,
    .w = fontSurface->w,
    .h = fontSurface->h,
  };

  const SDL_Rect destination = {
    .x = x,
    .y = y,
    .w = fontSurface->w,
    .h = fontSurface->h,
  };

  SDL_RenderCopy(renderer, texture, &source, &destination);
  SDL_DestroyTexture(texture);
  SDL_FreeSurface(fontSurface);
}

void RenderHealthBar(i32 health) {
  constexpr i32 scale = 2;

  constexpr i32 x = 20;
  constexpr i32 y = 20;

  i32 cropX = 364;
  constexpr i32 cropY = 199;
  i32 cropWidth = 130;
  constexpr i32 cropHeight = 14;

  RenderSprite(0,
               x,
               y,
               cropWidth * scale,
               cropHeight * scale,
               cropX,
               cropY,
               cropWidth,
               cropHeight);

  cropX = 494;
  cropWidth = 7;

  if (health < 0) {
    return;
  }

  constexpr i32 maxElements = 18;

  auto elementCount = health * maxElements / 100;
  if (elementCount > maxElements) {
    elementCount = maxElements;
  }

  for (i32 i = 0; i < elementCount; i++) {
    RenderSprite(0,
                 x + i * cropWidth * scale,
                 y,
                 cropWidth * scale,
                 cropHeight * scale,
                 cropX,
                 cropY,
                 cropWidth,
                 cropHeight);
  }
}

void RenderSmallLogo(void) {
  constexpr i32 cropWidth = 115;
  constexpr i32 cropHeight = 58;

  constexpr i32 scale = 1;

  constexpr auto width = cropWidth * scale;
  constexpr auto height = cropHeight * scale;

  RenderSprite(
    0, WindowWidth - width, 0, width, height, 700, 401, cropWidth, cropHeight);
}
