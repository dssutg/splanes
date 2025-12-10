#pragma once

#include <SDL2/SDL_rect.h>

#include "../util/util.h"

// Shared player properties
constexpr auto MaxPlayerHealth = 100;
constexpr auto PlayerMaxBombTickTime = 50;

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

// Entity fat struct
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
  u32 data;
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

// Entity polymorphic dispatch table entry
typedef struct EntityTableEntry {
  void (*Tick)(Entity *entity);
  void (*Render)(Entity *entity);
  const i32 zIndex;
} EntityTableEntry;

// Entity polymorphic dispatch table
extern const EntityTableEntry entityTable[];

// All entities in game
extern Entity *entities;

// Reference to the player in entity list
extern Entity *player;

// Base entity method
Entity *NewEntity(EntityType type);
void FreeEntity(Entity *entity);
void HurtEntity(Entity *entity, i32 damage);
void RenderEntitySprite(const Entity *const entity);
void RemoveAllEntities(void);

// Player entity methods
Entity *NewPlayer(void);
void PlayerDoDie(Entity *player);
void HealPlayer(i32 healPoints);
void PlayerTick(Entity *entity);
void PlayerRender(Entity *entity);

// Bomb entity methods
Entity *NewBomb(void);
void BombTick(Entity *entity);
void BombRender(Entity *entity);

// Bullet entity methods
Entity *NewBullet(i32 type,
                  i32 x,
                  i32 y,
                  i32 xa,
                  i32 ya,
                  i32 ownertype,
                  i32 damage,
                  u32 bulletframeno);
void BulletTick(Entity *entity);
void BulletRender(Entity *entity);

// Enemy plane entity methods
Entity *NewEnemyPlane(void);
void EnemyPlaneTick(Entity *entity);
void EnemyPlaneRender(Entity *entity);

// Explosion entity methods
Entity *NewExplosion(i32 x, i32 y);
void ExplosionTick(Entity *entity);
void ExplosionRender(Entity *entity);

// Healer entity methods
Entity *NewHealer(void);
void HealerTick(Entity *entity);
void HealerRender(Entity *entity);

// Island entity methods
Entity *NewIsland(void);
void IslandTick(Entity *entity);
void IslandRender(Entity *entity);

// Ship entity methods
Entity *NewShip(void);
void ShipTick(Entity *entity);
void ShipRender(Entity *entity);
