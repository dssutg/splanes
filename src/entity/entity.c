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
  entity->prev = NULL;
  entity->next = entities;
  if (entities != NULL) {
    entities->prev = entity; // not first element?
  }

  entities = entity;

  return entity;
}

void FreeEntity(Entity *entity) {
  if (entity == NULL || entities == NULL) {
    return;
  }

  if (entity == entities) {
    entities = entity->next;
  }

  if (entity->next != NULL) {
    entity->next->prev = entity->prev;
  }
  if (entity->prev != NULL) {
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

  PlaySound(SoundHurt, 100, 0);
}

void RenderEntitySprite(const Entity *entity) {
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
  Entity *next = NULL;
  for (Entity *entity = entities; entity != NULL; entity = next) {
    next = entity->next;
    FreeEntity(entity);
  }
}
