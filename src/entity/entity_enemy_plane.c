#include "entity.h"

#include "../util/util.h"
#include "../renderer/renderer.h"

static constexpr SDL_Rect frames[] = {
  {1, 1, 32, 32},
  {1, 34, 32, 32},
  {1, 67, 32, 32},
  {1, 100, 32, 32},
  {1, 133, 32, 32},
};

Entity *NewEnemyPlane() {
  auto e = NewEntity(EntityEnemyPlane);

  e->texture = 0;

  const auto frame = &frames[rand() % (i32)ArrayLength(frames)];

  e->crop.x = frame->x;
  e->crop.y = frame->y;
  e->crop.w = frame->w;
  e->crop.h = frame->h;

  e->pos.w = e->crop.w * 2;
  e->pos.h = e->crop.h * 2;
  e->pos.x = rand() % WindowWidth;
  e->pos.y = -(rand() % WindowHeight) - e->pos.h;

  e->xa = 0;
  e->ya = 1;

  e->hasShot = false;
  e->hasBombed = false;
  e->health = 100;

  return e;
}

void EnemyPlaneTick(Entity *e) {
  if (e->health <= 0) {
    NewExplosion(e->pos.x, e->pos.y);
    e->removed = true;
    return;
  }

  if (!e->hasShot) {
    e->hasShot = true;
    const auto x = e->pos.x + (e->pos.w - 128) / 2;
    const auto y = e->pos.y + e->pos.w;
    NewBullet(EntityBullet, x, y, 0, 2, e->type, 2, 1);
  } else {
    if (e->hasShot) {
      e->tickTime++;
      if (e->tickTime > 20) {
        e->hasShot = 0;
        e->tickTime = 0;
      }
    }
  }

  e->pos.x += e->xa * 20;
  e->pos.y += e->ya * 20;

  if (e->pos.y >= WindowHeight) {
    e->removed = true;
  }

  if (SDL_HasIntersection(&e->pos, &player->pos)) {
    player->health = 0;
  }
}

void EnemyPlaneRender(Entity *e) {
  RenderEntitySprite(e);
}
