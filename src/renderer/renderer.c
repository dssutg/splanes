#include <SDL2/SDL_render.h>
#include <SDL2/SDL_video.h>

#ifdef __APPLE__
#include <OpenGL/gl.h>
#include <OpenGL/glu.h>
#else
#include <GL/gl.h>
#include <GL/glu.h>
#endif
#include <SDL2/SDL.h>
#include <SDL2/SDL_image.h>
#include <SDL2/SDL_mixer.h>
#include <SDL2/SDL_ttf.h>

#include "renderer.h"

#include "../util/util.h"

SDL_Renderer *renderer;
SDL_Texture *textures[TextureCount];
SDL_Window *window;
const char *windowTitle = "Splanes";

void RenderSprite(i32 texture,
                  i32 x,
                  i32 y,
                  i32 width,
                  i32 height,
                  i32 cropX,
                  i32 cropY,
                  i32 cropWidth,
                  i32 cropHeight) {
  const SDL_Rect src = {
    .x = cropX,
    .y = cropY,
    .w = cropWidth,
    .h = cropHeight,
  };

  const SDL_Rect dest = {
    .x = x,
    .y = y,
    .w = width,
    .h = height,
  };

  SDL_RenderCopy(renderer, textures[texture], &src, &dest);
}

SDL_Texture *LoadTexture(const char *filename) {
  SDL_Surface *bitmap = IMG_Load(filename);
  if (bitmap == NULL) {
    Fatalf("%s", IMG_GetError());
  }

  SDL_Texture *texture = SDL_CreateTextureFromSurface(renderer, bitmap);
  if (texture == NULL) {
    Fatalf("can't create texture from %s: %s", filename, SDL_GetError());
  }

  SDL_FreeSurface(bitmap);

  return texture;
}

void LoadTextures(void) {
  textures[0] = LoadTexture("assets/sprites/sprites.png");
}

void RenderRect(i32 x,
                i32 y,
                i32 width,
                i32 height,
                i32 red,
                i32 green,
                i32 blue) {
  const SDL_Rect rect = {
    .x = x,
    .y = y,
    .w = width,
    .h = height,
  };

  SDL_SetRenderDrawColor(renderer, red, green, blue, 255);
  SDL_RenderFillRect(renderer, &rect);
}
