#include "entity.h"

#include "../renderer/renderer.h"
#include "../sound_manager/sound_manager.h"
#include "../util/util.h"

static constexpr SDL_Rect frames[] = {
  {67, 166, 32, 32},
  {100, 166, 32, 32},
  {133, 166, 32, 32},
  {166, 166, 32, 32},
  {199, 166, 32, 32},
  {232, 166, 32, 32},
};

Entity *NewExplosion(i32 x, i32 y) {
  auto e = NewEntity(EntityExplosion);

  e->texture = 0;
  e->pos.x = x;
  e->pos.y = y;
  e->pos.w = frames[0].w * 2;
  e->pos.h = frames[0].h * 2;

  PlaySound(SoundExplosion1, 100);

  return e;
}

void ExplosionTick(Entity *e) {
  e->tickTime++;
  if (e->tickTime >= ArrayLength(frames)) {
    e->removed = true;
  }
}

void ExplosionRender(Entity *e) {
  const auto f = &frames[e->tickTime % ArrayLength(frames)];

  RenderSprite(
    e->texture, e->pos.x, e->pos.y, e->pos.w, e->pos.h, f->x, f->y, f->w, f->h);
}
