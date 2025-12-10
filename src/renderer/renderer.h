#pragma once

#include <SDL2/SDL_render.h>

#include "../util/util.h"

constexpr auto WindowDelayMilliseconds = 50;
constexpr auto WindowScale = 1;
constexpr auto WindowWidth = 800 * WindowScale;
constexpr auto WindowHeight = 600 * WindowScale;
constexpr auto TileSize = 32;
constexpr auto TextureCount = 16;

extern SDL_Renderer *renderer;
extern SDL_Texture *textures[TextureCount];
extern SDL_Window *window;
extern const char *windowTitle;

void RenderSprite(i32 texture,
                  i32 x,
                  i32 y,
                  i32 width,
                  i32 height,
                  i32 cropX,
                  i32 cropY,
                  i32 cropWidth,
                  i32 cropHeight);

SDL_Texture *LoadTexture(const char *const filename);

void LoadTextures(void);

void RenderRect(i32 x, i32 y, i32 width, i32 height, u8 red, u8 green, u8 blue);
