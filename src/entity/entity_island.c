#include "entity.h"

#include "../renderer/renderer.h"

static const SDL_Rect islandFrames[] = {
  {100, 496, 64, 65},
  {165, 496, 64, 65},
  {230, 496, 64, 65},
};

Entity *NewIsland(void) {
  Entity *island = NewEntity(EntityIsland);

  island->texture = 0;
  island->data = rand() % 3;

  island->pos.w = islandFrames[island->data].w * 3;
  island->pos.h = islandFrames[island->data].h * 3;
  island->pos.x = rand() % WindowWidth;
  island->pos.y = -(rand() % WindowHeight) - island->pos.h;

  island->xa = 0;
  island->ya = 1;

  return island;
}

void IslandTick(Entity *entity) {
  entity->pos.x += entity->xa * 10;
  entity->pos.y += entity->ya * 10;

  if (entity->pos.y >= WindowHeight) {
    entity->removed = true;
  }
}

void IslandRender(Entity *entity) {
  const SDL_Rect *frame = &islandFrames[entity->data];

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
