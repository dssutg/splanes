#include "entity.h"

#include "../renderer/renderer.h"
#include "../sound_manager/sound_manager.h"
#include "../util/util.h"

static const SDL_Rect explosionFrames[] = {
  {67, 166, 32, 32},
  {100, 166, 32, 32},
  {133, 166, 32, 32},
  {166, 166, 32, 32},
  {199, 166, 32, 32},
  {232, 166, 32, 32},
};

Entity *NewExplosion(i32 x, i32 y) {
  Entity *explosion = NewEntity(EntityExplosion);

  explosion->texture = 0;
  explosion->pos.x = x;
  explosion->pos.y = y;
  explosion->pos.w = explosionFrames[0].w * 2;
  explosion->pos.h = explosionFrames[0].h * 2;

  PlaySound(SoundExplosion1, 100, 0);

  return explosion;
}

void ExplosionTick(Entity *entity) {
  entity->tickTime++;
  if (entity->tickTime >= ArrayLength(explosionFrames)) {
    entity->removed = true;
  }
}

void ExplosionRender(Entity *entity) {
  const SDL_Rect *frame =
    &explosionFrames[entity->tickTime % ArrayLength(explosionFrames)];

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
