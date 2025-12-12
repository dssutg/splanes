#include "../lib/std.h"

#include "entity_island.h"

#include "../renderer/renderer.h"

static constexpr auto frames = std::to_array<SDL_Rect>({
  {100, 496, 64, 65},
  {165, 496, 64, 65},
  {230, 496, 64, 65},
});

Island::Island() {
  frameNo = static_cast<uint32_t>(rand()) % 3;

  pos.w = frames[frameNo].w * 3;
  pos.h = frames[frameNo].h * 3;

  pos.x = rand() % WindowWidth;
  pos.y = -(rand() % WindowHeight) - pos.h;

  ya = 1;
}

void Island::tick() {
  pos.x += xa * 10;
  pos.y += ya * 10;
  if (pos.y >= WindowHeight) {
    removed = true;
  }
}

void Island::render() {
  const auto f = &frames[frameNo];
  RenderSprite(texture,
               {.x = pos.x, .y = pos.y, .w = pos.w, .h = pos.h},
               {.x = f->x, .y = f->y, .w = f->w, .h = f->h});
}
