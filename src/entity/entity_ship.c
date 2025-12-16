#include "entity.h"

#include "../renderer/renderer.h"
#include "../util/util.h"

static constexpr SDL_Rect frames[] = {
  {505, 298, 41, 197},
  {463, 298, 41, 197},
};

Entity *NewShip() {
  auto e = NewEntity(EntityShip);

  e->texture = 0;

  const auto frame = &frames[0];

  e->pos.w = frame->w * 1;
  e->pos.h = frame->h * 1;
  e->pos.x = rand() % WindowWidth,
  e->pos.y = -(rand() % WindowHeight) - e->pos.h;

  e->xa = 0;
  e->ya = 1;

  e->hasShot = false;
  e->hasBombed = false;
  e->health = 100;

  return e;
}

void ShipTick(Entity *e) {
  e->tickTime++;
  if (e->tickTime > 10) {
    e->tickTime = 0;
  }

  if (e->health <= 0) {
    NewExplosion(e->pos.x, e->pos.y);
    e->removed = true;
    return;
  }

  e->pos.x += e->xa * 11;
  e->pos.y += e->ya * 11;

  if (e->pos.y >= WindowHeight) {
    e->removed = true;
  }
}

void ShipRender(Entity *e) {
  const auto f = &frames[e->tickTime / 5 % ArrayLength(frames)];

  RenderSprite(
    e->texture, e->pos.x, e->pos.y, e->pos.w, e->pos.h, f->x, f->y, f->w, f->h);
}
