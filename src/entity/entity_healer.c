#include "entity.h"

#include "../renderer/renderer.h"

static constexpr SDL_Rect healerFrames[] = {
  {166, 265, 29, 15},
};

Entity *NewHealer(void) {
  auto healer = NewEntity(EntityHealer);

  healer->texture = 0;
  healer->data = 0;

  healer->pos.w = healerFrames[healer->data].w * 2;
  healer->pos.h = healerFrames[healer->data].h * 2;
  healer->pos.x = rand() % WindowWidth;
  healer->pos.y = -(rand() % WindowHeight) - healer->pos.h;

  healer->xa = 0;
  healer->ya = 1;

  return healer;
}

void HealerTick(Entity *entity) {
  entity->pos.x += entity->xa * 10;
  entity->pos.y += entity->ya * 10;

  if (entity->pos.y >= WindowHeight) {
    entity->removed = true;
  }

  if (SDL_HasIntersection(&entity->pos, &player->pos)) {
    HealPlayer(20);
    entity->removed = true;
  }
}

void HealerRender(Entity *entity) {
  const auto frame = &healerFrames[entity->data];

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
