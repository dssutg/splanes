#include "entity.h"

#include "../renderer/renderer.h"
#include "../util/util.h"

static constexpr SDL_Rect shipFrames[] = {
  {505, 298, 41, 197},
  {463, 298, 41, 197},
};

Entity *NewShip(void) {
  auto ship = NewEntity(EntityShip);

  ship->texture = 0;

  const auto frame = &shipFrames[0];

  ship->pos.w = frame->w * 1;
  ship->pos.h = frame->h * 1;
  ship->pos.x = rand() % WindowWidth,
  ship->pos.y = -(rand() % WindowHeight) - ship->pos.h;

  ship->xa = 0;
  ship->ya = 1;

  ship->hasShot = false;
  ship->hasBombed = false;
  ship->health = 100;

  return ship;
}

void ShipTick(Entity *entity) {
  entity->tickTime++;
  if (entity->tickTime > 10) {
    entity->tickTime = 0;
  }

  if (entity->health <= 0) {
    NewExplosion(entity->pos.x, entity->pos.y);
    entity->removed = true;
    return;
  }

  entity->pos.x += entity->xa * 11;
  entity->pos.y += entity->ya * 11;

  if (entity->pos.y >= WindowHeight) {
    entity->removed = true;
  }
}

void ShipRender(Entity *entity) {
  const auto frame =
    &shipFrames[entity->tickTime / 5 % ArrayLength(shipFrames)];

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
