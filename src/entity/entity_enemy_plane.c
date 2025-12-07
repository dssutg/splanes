#include "entity.h"
#include "entity_bullet.h"
#include "entity_explosion.h"

#include "../util/util.h"
#include "../renderer/renderer.h"

static const SDL_Rect enemyPlaneFrames[] = {
  {1, 1, 32, 32},
  {1, 34, 32, 32},
  {1, 67, 32, 32},
  {1, 100, 32, 32},
  {1, 133, 32, 32},
};

Entity *NewEnemyPlane(void) {
  Entity *enemyPlane = NewEntity(EntityEnemyPlane);

  enemyPlane->texture = 0;

  const SDL_Rect *frame =
    &enemyPlaneFrames[rand() % ArrayLength(enemyPlaneFrames)];

  enemyPlane->crop.x = frame->x;
  enemyPlane->crop.y = frame->y;
  enemyPlane->crop.w = frame->w;
  enemyPlane->crop.h = frame->h;

  enemyPlane->pos.w = enemyPlane->crop.w * 2;
  enemyPlane->pos.h = enemyPlane->crop.h * 2;
  enemyPlane->pos.x = rand() % WindowWidth;
  enemyPlane->pos.y = -(rand() % WindowHeight) - enemyPlane->pos.h;

  enemyPlane->xa = 0;
  enemyPlane->ya = 1;

  enemyPlane->hasShot = false;
  enemyPlane->hasBombed = false;
  enemyPlane->health = 100;

  return enemyPlane;
}

void EnemyPlaneTick(Entity *entity) {
  if (entity->health <= 0) {
    NewExplosion(entity->pos.x, entity->pos.y);
    entity->removed = true;
    return;
  }

  if (!entity->hasShot) {
    entity->hasShot = true;
    NewBullet(EntityBullet,
              entity->pos.x + (entity->pos.w - 128) / 2,
              entity->pos.y + entity->pos.w,
              0,
              2,
              entity->type,
              2,
              1);
  } else {
    if (entity->hasShot) {
      entity->tickTime++;

      if (entity->tickTime > 20) {
        entity->hasShot = 0;
        entity->tickTime = 0;
      }
    }
  }

  entity->pos.x += entity->xa * 20;
  entity->pos.y += entity->ya * 20;

  if (entity->pos.y >= WindowHeight) {
    entity->removed = true;
  }

  if (SDL_HasIntersection(&entity->pos, &player->pos)) {
    player->health = 0;
  }
}

void EnemyPlaneRender(Entity *entity) {
  RenderEntitySprite(entity);
}
