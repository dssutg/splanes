#include "../lib/std.h"

#include "renderer.h"

#include "../util/util.h"

void RenderSprite(int32_t texture, const SDL_Rect &dest, const SDL_Rect &src) {
  SDL_RenderCopy(renderer, textures[texture], &src, &dest);
}

SDL_Texture *LoadTexture(const std::string &filename) {
  auto bitmap = IMG_Load(filename.c_str());
  if (bitmap == nullptr) {
    Fatal(IMG_GetError());
  }

  auto texture = SDL_CreateTextureFromSurface(renderer, bitmap);
  if (texture == nullptr) {
    Fatal(std::format(
      "can't create texture from {}: {}", filename, SDL_GetError()));
  }

  SDL_FreeSurface(bitmap);

  return texture;
}

void LoadTextures() {
  textures[0] = LoadTexture("assets/sprites/sprites.png");
}

void RenderRect(const SDL_Rect &rect, const SDL_Color &color) {
  SDL_SetRenderDrawColor(renderer, color.r, color.g, color.b, color.a);
  SDL_RenderFillRect(renderer, &rect);
}
