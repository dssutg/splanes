#pragma once

#include <SDL2/SDL_render.h>

#include "../util/util.h"

enum {
  WindowDelayMilliseconds = 50,
  WindowScale = 1,
  WindowWidth = 800 * WindowScale,
  WindowHeight = 600 * WindowScale,
  TileSize = 32,
  TextureCount = 16,
};

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

SDL_Texture *LoadTexture(const char *filename);

void LoadTextures(void);

void RenderRect(i32 x,
                i32 y,
                i32 width,
                i32 height,
                i32 red,
                i32 green,
                i32 blue);
