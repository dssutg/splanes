#include "entity.h"

#include "../sound_manager/sound_manager.h"
#include "../util/util.h"
#include "../renderer/renderer.h"

Entity *entities;
Entity *player;

Entity *NewEntity(EntityType type) {
  Entity *e = Emalloc(sizeof(*e));

  e->type = type;
  e->tickTime = 0;
  e->removed = false;

  // insert to the head of entity list
  e->prev = nullptr;
  e->next = entities;
  if (entities != nullptr) {
    entities->prev = e; // not first element?
  }

  entities = e;

  return e;
}

void FreeEntity(Entity *e) {
  if (e == nullptr || entities == nullptr) {
    return;
  }

  if (e == entities) {
    entities = e->next;
  }

  if (e->next != nullptr) {
    e->next->prev = e->prev;
  }
  if (e->prev != nullptr) {
    e->prev->next = e->next;
  }

  free(e);
}

void HurtEntity(Entity *e, i32 damage) {
  if (e->type != EntityPlayer && e->type != EntityEnemyPlane &&
      e->type != EntityShip) {
    return;
  }

  if (e->health < damage) {
    e->health = 0;
  } else {
    e->health -= damage;
  }

  PlaySound(SoundHurt, 100);
}

void RenderEntitySprite(const Entity *e) {
  RenderSprite(e->texture,
               e->pos.x,
               e->pos.y,
               e->pos.w,
               e->pos.h,
               e->crop.x,
               e->crop.y,
               e->crop.w,
               e->crop.h);
}

void RemoveAllEntities() {
  Entity *next = nullptr;
  for (auto entity = entities; entity != nullptr; entity = next) {
    next = entity->next;
    FreeEntity(entity);
  }
}
