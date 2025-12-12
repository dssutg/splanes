#pragma once

#include "../lib/std.h"

constexpr auto WindowDelayMilliseconds = 50;
constexpr auto WindowScale = 1;
constexpr auto WindowWidth = 800 * WindowScale;
constexpr auto WindowHeight = 600 * WindowScale;

constexpr std::string windowTitle = "Splanes";

constexpr auto TextureCount = 16;

constexpr auto TileSize = 32;

inline SDL_Renderer *renderer;
inline SDL_Texture *textures[TextureCount];
inline SDL_Window *window;

void RenderSprite(int32_t texture, const SDL_Rect &dest, const SDL_Rect &src);

SDL_Texture *LoadTexture(const std::string &filename);

void LoadTextures();

void RenderRect(const SDL_Rect &rect, const SDL_Color &color);
