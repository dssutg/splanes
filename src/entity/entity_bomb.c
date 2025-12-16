#include "entity.h"

#include "../renderer/renderer.h"

static constexpr SDL_Rect frame[] = {
  {265, 265, 9, 21},
};

Entity *NewBomb() {
  auto e = NewEntity(EntityBomb);

  e->texture = 0;
  e->data = 1;

  e->pos.x = 0;
  e->pos.y = 0;
  e->pos.w = frame[0].w * 2;
  e->pos.h = frame[0].h * 2;

  e->xa = 0;
  e->ya = 0;

  e->damage = 1000;
  e->tickTime = 1;

  return e;
}

static void BombCalcSize(Entity *e, i32 *w, i32 *h) {
  const auto scale = pow(0.90, e->tickTime);

  *w = (i32)(e->pos.w * scale);
  *h = (i32)(e->pos.h * scale);

  if (*w == 0 || *h == 0) {
    e->data = 0;
    *w = 1;
    *h = 1;
  }
}

void BombTick(Entity *e) {
  e->tickTime += e->data;

  e->pos.x += e->xa * 2;
  e->pos.y += e->ya * 2;

  if (e->data != 0) {
    return;
  }

  SDL_Rect scaledrect;
  scaledrect.x = e->pos.x;
  scaledrect.y = e->pos.y;
  BombCalcSize(e, &scaledrect.w, &scaledrect.h);

  Entity *nextEntity = nullptr;
  for (auto it = entities; it != nullptr; it = nextEntity) {
    nextEntity = it->next;

    if (it->type == EntityShip && SDL_HasIntersection(&it->pos, &scaledrect)) {
      HurtEntity(it, e->damage);
      NewExplosion(scaledrect.x, scaledrect.y);

      if (player->score + 1 > player->score) {
        player->score++;
      }

      break;
    }
  }

  e->removed = true;
}

void BombRender(Entity *e) {
  const auto f = &frame[0];

  i32 width, height;
  BombCalcSize(e, &width, &height);

  RenderSprite(
    e->texture, e->pos.x, e->pos.y, width, height, f->x, f->y, f->w, f->h);
}
