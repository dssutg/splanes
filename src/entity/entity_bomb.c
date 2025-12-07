#include "entity.h"
#include "entity_explosion.h"

#include "../renderer/renderer.h"

static const SDL_Rect bombFrames[] = {
  {265, 265, 9, 21},
};

Entity *NewBomb(void) {
  Entity *bomb = NewEntity(EntityBomb);

  bomb->texture = 0;
  bomb->data = 1;

  bomb->pos.x = 0;
  bomb->pos.y = 0;
  bomb->pos.w = bombFrames[0].w * 2;
  bomb->pos.h = bombFrames[0].h * 2;

  bomb->xa = 0;
  bomb->ya = 0;

  bomb->damage = 1000;
  bomb->tickTime = 1;

  return bomb;
}

static void BombCalcSize(Entity *bomb, i32 *w, i32 *h) {
  const f64 scale = pow(0.90, bomb->tickTime);

  *w = (i32)(bomb->pos.w * scale);
  *h = (i32)(bomb->pos.h * scale);

  if (*w == 0 || *h == 0) {
    bomb->data = 0;
    *w = 1;
    *h = 1;
  }
}

void BombTick(Entity *entity) {
  entity->tickTime += entity->data;

  entity->pos.x += entity->xa * 2;
  entity->pos.y += entity->ya * 2;

  if (entity->data != 0) {
    return;
  }

  SDL_Rect scaledrect;
  scaledrect.x = entity->pos.x;
  scaledrect.y = entity->pos.y;
  BombCalcSize(entity, &scaledrect.w, &scaledrect.h);

  Entity *nextEntity = NULL;
  for (Entity *it = entities; it != NULL; it = nextEntity) {
    nextEntity = it->next;

    if (it->type == EntityShip && SDL_HasIntersection(&it->pos, &scaledrect)) {
      HurtEntity(it, entity->damage);
      NewExplosion(scaledrect.x, scaledrect.y);

      if (player->score + 1 > player->score) {
        player->score++;
      }

      break;
    }
  }

  entity->removed = true;
}

void BombRender(Entity *entity) {
  const SDL_Rect *frame = &bombFrames[0];

  i32 width, height;
  BombCalcSize(entity, &width, &height);

  RenderSprite(entity->texture,
               entity->pos.x,
               entity->pos.y,
               width,
               height,
               frame->x,
               frame->y,
               frame->w,
               frame->h);
}
