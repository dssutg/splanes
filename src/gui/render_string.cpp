#include "../lib/std.h"

#include "gui.h"

#include "../renderer/renderer.h"
#include "../util/util.h"

class FontCacheEntry {
  public:
  TTF_Font *font;
  int32_t size;
};

static std::vector<FontCacheEntry> fontCache;

static TTF_Font *LoadFont(int32_t size) {
  const auto existing = std::find_if(
    fontCache.begin(),
    fontCache.end(),
    [size](const FontCacheEntry &entry) -> bool { return entry.size == size; });

  if (existing != fontCache.end()) {
    return existing->font;
  }

  // Open the font.
  auto font = TTF_OpenFont("assets/fonts/OpenSans-Bold.ttf", size);
  if (font == nullptr) {
    Fatal(SDL_GetError());
  }

  fontCache.emplace_back(FontCacheEntry{
    .font = font,
    .size = size,
  });

  return font;
}

void FreeFontCache() {
  for (const auto &entry : fontCache) {
    TTF_CloseFont(entry.font);
  }
  fontCache.clear();
}

void RenderString(int32_t x,
                  int32_t y,
                  int32_t size,
                  const SDL_Color &color,
                  bool centerRelativeToWindow,
                  int32_t lineNumber,
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

  const auto font = LoadFont(size);

  auto fontSurface = TTF_RenderText_Blended(font, text, color);
  if (fontSurface == nullptr) {
    Fatal(SDL_GetError());
  }

  auto texture = SDL_CreateTextureFromSurface(renderer, fontSurface);
  if (texture == nullptr) {
    Fatal(SDL_GetError());
  }

  if (centerRelativeToWindow) {
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
