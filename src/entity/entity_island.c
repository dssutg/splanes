#include "entity.h"

#include "../renderer/renderer.h"

static constexpr SDL_Rect frames[] = {
  {100, 496, 64, 65},
  {165, 496, 64, 65},
  {230, 496, 64, 65},
};

Entity *NewIsland() {
  auto e = NewEntity(EntityIsland);

  e->texture = 0;
  e->data = (u32)(rand()) % 3;

  e->pos.w = frames[e->data].w * 3;
  e->pos.h = frames[e->data].h * 3;
  e->pos.x = rand() % WindowWidth;
  e->pos.y = -(rand() % WindowHeight) - e->pos.h;

  e->xa = 0;
  e->ya = 1;

  return e;
}

void IslandTick(Entity *e) {
  e->pos.x += e->xa * 10;
  e->pos.y += e->ya * 10;

  if (e->pos.y >= WindowHeight) {
    e->removed = true;
  }
}

void IslandRender(Entity *e) {
  const auto f = &frames[e->data];

  RenderSprite(
    e->texture, e->pos.x, e->pos.y, e->pos.w, e->pos.h, f->x, f->y, f->w, f->h);
}
