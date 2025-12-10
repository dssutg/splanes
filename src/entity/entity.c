#include "entity.h"

#include "../sound_manager/sound_manager.h"
#include "../util/util.h"
#include "../renderer/renderer.h"

Entity *entities;
Entity *player;

Entity *NewEntity(EntityType type) {
  Entity *entity = Emalloc(sizeof(*entity));

  entity->type = type;
  entity->tickTime = 0;
  entity->removed = false;

  // insert to the head of entity list
  entity->prev = nullptr;
  entity->next = entities;
  if (entities != nullptr) {
    entities->prev = entity; // not first element?
  }

  entities = entity;

  return entity;
}

void FreeEntity(Entity *entity) {
  if (entity == nullptr || entities == nullptr) {
    return;
  }

  if (entity == entities) {
    entities = entity->next;
  }

  if (entity->next != nullptr) {
    entity->next->prev = entity->prev;
  }
  if (entity->prev != nullptr) {
    entity->prev->next = entity->next;
  }

  free(entity);
}

void HurtEntity(Entity *entity, i32 damage) {
  if (entity->type != EntityPlayer && entity->type != EntityEnemyPlane &&
      entity->type != EntityShip) {
    return;
  }

  if (entity->health < damage) {
    entity->health = 0;
  } else {
    entity->health -= damage;
  }

  PlaySound(SoundHurt, 100);
}

void RenderEntitySprite(const Entity *const entity) {
  RenderSprite(entity->texture,
               entity->pos.x,
               entity->pos.y,
               entity->pos.w,
               entity->pos.h,
               entity->crop.x,
               entity->crop.y,
               entity->crop.w,
               entity->crop.h);
}

void RemoveAllEntities(void) {
  Entity *next = nullptr;
  for (auto entity = entities; entity != nullptr; entity = next) {
    next = entity->next;
    FreeEntity(entity);
  }
}
