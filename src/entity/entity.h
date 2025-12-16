#pragma once

#include <SDL2/SDL_rect.h>

#include "../util/util.h"

// Shared player properties
constexpr auto MaxPlayerHealth = 100;
constexpr auto PlayerMaxBombTickTime = 50;

// Entity types
typedef enum : u8 {
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
  EntityType ownerType; // entity type that spawned this bullet

  // EntityPlayer
  u64 score;
  u64 distance;
  bool hasShot;
  bool hasBombed;
  i32 deathTime;
  i32 bombTickTime;
} Entity;

// Entity polymorphic dispatch table entry
typedef struct {
  void (*Tick)(Entity *e);
  void (*Render)(Entity *e);
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
void FreeEntity(Entity *e);
void HurtEntity(Entity *e, i32 damage);
void RenderEntitySprite(const Entity *e);
void RemoveAllEntities();

// Player entity methods
Entity *NewPlayer();
void PlayerDoDie(Entity *player);
void HealPlayer(i32 healPoints);
void PlayerTick(Entity *e);
void PlayerRender(Entity *e);

// Bomb entity methods
Entity *NewBomb();
void BombTick(Entity *e);
void BombRender(Entity *e);

// Bullet entity methods
Entity *NewBullet(EntityType type,
                  i32 x,
                  i32 y,
                  i32 xa,
                  i32 ya,
                  EntityType ownertype,
                  i32 damage,
                  u32 bulletframeno);
void BulletTick(Entity *e);
void BulletRender(Entity *e);

// Enemy plane entity methods
Entity *NewEnemyPlane();
void EnemyPlaneTick(Entity *e);
void EnemyPlaneRender(Entity *e);

// Explosion entity methods
Entity *NewExplosion(i32 x, i32 y);
void ExplosionTick(Entity *e);
void ExplosionRender(Entity *e);

// Healer entity methods
Entity *NewHealer();
void HealerTick(Entity *e);
void HealerRender(Entity *e);

// Island entity methods
Entity *NewIsland();
void IslandTick(Entity *e);
void IslandRender(Entity *e);

// Ship entity methods
Entity *NewShip();
void ShipTick(Entity *e);
void ShipRender(Entity *e);
