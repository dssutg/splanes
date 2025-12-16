#include "entity.h"

#include "../renderer/renderer.h"

static constexpr SDL_Rect frames[] = {
  {1, 166, 32, 32},
  {34, 199, 32, 32},
};

Entity *NewBullet(EntityType type,
                  i32 x,
                  i32 y,
                  i32 xa,
                  i32 ya,
                  EntityType ownertype,
                  i32 damage,
                  u32 bulletframeno) {
  auto e = NewEntity(type);

  e->texture = 0;
  e->data = bulletframeno;

  e->pos.x = x;
  e->pos.y = y;
  e->pos.w = frames[bulletframeno].w * 2;
  e->pos.h = frames[bulletframeno].h * 2;

  e->xa = xa;
  e->ya = ya;

  e->ownerType = ownertype;
  e->damage = damage;

  return e;
}

void BulletTick(Entity *e) {
  e->pos.x += e->xa * 20;
  e->pos.y += e->ya * 20;

  auto dead = false;

  Entity *nextEntity = nullptr;
  for (auto it = entities; it != nullptr; it = nextEntity) {
    nextEntity = it->next;

    if ((e->ownerType == EntityPlayer && it->type == EntityEnemyPlane) ||
        (e->ownerType == EntityEnemyPlane && it->type == EntityPlayer)) {
      if (SDL_HasIntersection(&it->pos, &e->pos)) {
        dead = true;

        HurtEntity(it, e->damage);

        if (it->type == EntityPlayer) {
          NewExplosion(e->pos.x, e->pos.y);
        } else if (e->ownerType == EntityPlayer) {
          if (player->score + 1 > player->score) {
            player->score++;
          }
        }
      }
    }
  }

  const SDL_Rect windowRect = {
    .x = 0,
    .y = 0,
    .w = WindowWidth,
    .h = WindowHeight,
  };

  if (dead || !SDL_HasIntersection(&e->pos, &windowRect)) {
    e->removed = true;
  }
}

void BulletRender(Entity *e) {
  const auto f = &frames[e->data];

  RenderSprite(
    e->texture, e->pos.x, e->pos.y, e->pos.w, e->pos.h, f->x, f->y, f->w, f->h);
}
