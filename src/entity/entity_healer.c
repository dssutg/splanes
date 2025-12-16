#include "entity.h"

#include "../renderer/renderer.h"

static constexpr SDL_Rect frames[] = {
  {166, 265, 29, 15},
};

Entity *NewHealer() {
  auto e = NewEntity(EntityHealer);

  e->texture = 0;
  e->data = 0;

  e->pos.w = frames[e->data].w * 2;
  e->pos.h = frames[e->data].h * 2;
  e->pos.x = rand() % WindowWidth;
  e->pos.y = -(rand() % WindowHeight) - e->pos.h;

  e->xa = 0;
  e->ya = 1;

  return e;
}

void HealerTick(Entity *e) {
  e->pos.x += e->xa * 10;
  e->pos.y += e->ya * 10;

  if (e->pos.y >= WindowHeight) {
    e->removed = true;
  }

  if (SDL_HasIntersection(&e->pos, &player->pos)) {
    HealPlayer(20);
    e->removed = true;
  }
}

void HealerRender(Entity *e) {
  const auto f = &frames[e->data];

  RenderSprite(
    e->texture, e->pos.x, e->pos.y, e->pos.w, e->pos.h, f->x, f->y, f->w, f->h);
}
