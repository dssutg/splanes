#pragma once

#include <SDL2/SDL_rect.h>

#include "../util/util.h"

// Entity types
typedef enum EntityType : u8 {
  EntityPlayer,
  EntityEnemyPlane,
  EntityBullet,
  EntityBomb,
  EntityIsland,
  EntityExplosion,
  EntityShip,
  EntityHealer,
} EntityType;

typedef struct Entity {
  // Common
  struct Entity *next;
  struct Entity *prev;
  SDL_Rect pos;
  SDL_Rect crop;
  i32 xa;
  i32 ya;
  i32 texture;
  EntityType type;
  bool removed;
  i32 health;
  i32 damage;
  i32 data;
  u32 tickTime;

  // EntityBullet
  i32 ownerType; // entity type that spawned this bullet

  // EntityPlayer
  u64 score;
  u64 distance;
  bool hasShot;
  bool hasBombed;
  i32 deathTime;
  i32 bombTickTime;
} Entity;

extern Entity *entities;
extern Entity *player;

Entity *NewEntity(EntityType type);
void FreeEntity(Entity *entity);
void HurtEntity(Entity *entity, i32 damage);
void RenderEntitySprite(const Entity *entity);
void RemoveAllEntities(void);
