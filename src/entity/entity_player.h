#pragma once

#include "entity.h"

enum {
  MaxPlayerHealth = 100,
  PlayerMaxBombTickTime = 50,
};

Entity *NewPlayer(void);
void PlayerDoDie(Entity *player);
void HealPlayer(i32 healPoints);
void PlayerTick(Entity *entity);
void PlayerRender(Entity *entity);
