#include "entity_bomb.h"
#include "entity_bullet.h"
#include "entity_enemy_plane.h"
#include "entity_explosion.h"
#include "entity_healer.h"
#include "entity_island.h"
#include "entity_player.h"
#include "entity_ship.h"
#include "entity_table.h"

const EntityTableEntry entityTable[] = {
  [EntityPlayer] =
    {
      .Tick = PlayerTick,
      .Render = PlayerRender,
      .zIndex = 2,
    },

  [EntityEnemyPlane] =
    {
      .Tick = EnemyPlaneTick,
      .Render = EnemyPlaneRender,
      .zIndex = 2,
    },

  [EntityBullet] =
    {
      .Tick = BulletTick,
      .Render = BulletRender,
      .zIndex = 2,
    },

  [EntityBomb] =
    {
      .Tick = BombTick,
      .Render = BombRender,
      .zIndex = 2,
    },

  [EntityIsland] =
    {
      .Tick = IslandTick,
      .Render = IslandRender,
      .zIndex = 0,
    },

  [EntityExplosion] =
    {
      .Tick = ExplosionTick,
      .Render = ExplosionRender,
      .zIndex = 2,
    },

  [EntityShip] =
    {
      .Tick = ShipTick,
      .Render = ShipRender,
      .zIndex = 1,
    },

  [EntityHealer] =
    {
      .Tick = HealerTick,
      .Render = HealerRender,
      .zIndex = 2,
    },
};
