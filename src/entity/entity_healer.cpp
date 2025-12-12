#include "../lib/std.h"

#include "entity_healer.h"

#include "../renderer/renderer.h"
#include "../game_loop/game_loop.h"
#include "../level/level.h"

static constexpr auto frames = std::to_array<SDL_Rect>({
  {166, 265, 29, 15},
});

Healer::Healer() {
  pos.w = frames[frameNo].w * 2;
  pos.h = frames[frameNo].h * 2;
  pos.x = rand() % WindowWidth;
  pos.y = -(rand() % WindowHeight) - pos.h;
  ya = 1;
}

int32_t Healer::getZIndex() const {
  return 2;
}

void Healer::tick() {
  pos.x += xa * 10;
  pos.y += ya * 10;

  if (pos.y >= WindowHeight) {
    removed = true;
  }

  if (level->player && SDL_HasIntersection(&pos, &(*level->player)->pos)) {
    (*level->player)->heal(20);
    removed = true;
  }
}

void Healer::render() {
  const auto f = &frames[frameNo];
  RenderSprite(texture,
               {.x = pos.x, .y = pos.y, .w = pos.w, .h = pos.h},
               {.x = f->x, .y = f->y, .w = f->w, .h = f->h});
}
