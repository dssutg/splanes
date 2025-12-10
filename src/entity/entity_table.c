#include "entity.h"

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
