#include "entity.h"
#include "entity_explosion.h"

#include "../renderer/renderer.h"

static const SDL_Rect bulletFrames[] = {
  {1, 166, 32, 32},
  {34, 199, 32, 32},
};

Entity *NewBullet(i32 type,
                  i32 x,
                  i32 y,
                  i32 xa,
                  i32 ya,
                  i32 ownertype,
                  i32 damage,
                  i32 bulletframeno) {
  Entity *bullet = NewEntity(type);

  bullet->texture = 0;
  bullet->data = bulletframeno;

  bullet->pos.x = x;
  bullet->pos.y = y;
  bullet->pos.w = bulletFrames[bulletframeno].w * 2;
  bullet->pos.h = bulletFrames[bulletframeno].h * 2;

  bullet->xa = xa;
  bullet->ya = ya;

  bullet->ownerType = ownertype;
  bullet->damage = damage;

  return bullet;
}

void BulletTick(Entity *entity) {
  entity->pos.x += entity->xa * 20;
  entity->pos.y += entity->ya * 20;

  bool dead = false;

  Entity *nextEntity = NULL;
  for (Entity *it = entities; it != NULL; it = nextEntity) {
    nextEntity = it->next;

    if ((entity->ownerType == EntityPlayer && it->type == EntityEnemyPlane) ||
        (entity->ownerType == EntityEnemyPlane && it->type == EntityPlayer)) {
      if (SDL_HasIntersection(&it->pos, &entity->pos)) {
        dead = true;

        HurtEntity(it, entity->damage);

        if (it->type == EntityPlayer) {
          NewExplosion(entity->pos.x, entity->pos.y);
        } else if (entity->ownerType == EntityPlayer) {
          if (player->score + 1 > player->score) {
            player->score++;
          }
        }
      }
    }
  }

  SDL_Rect windowRect = {
    .x = 0,
    .y = 0,
    .w = WindowWidth,
    .h = WindowHeight,
  };

  if (dead || !SDL_HasIntersection(&entity->pos, &windowRect)) {
    entity->removed = true;
  }
}

void BulletRender(Entity *entity) {
  const SDL_Rect *frame = &bulletFrames[entity->data];

  RenderSprite(entity->texture,
               entity->pos.x,
               entity->pos.y,
               entity->pos.w,
               entity->pos.h,
               frame->x,
               frame->y,
               frame->w,
               frame->h);
}
