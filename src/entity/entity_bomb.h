#pragma once

#include "entity.h"

Entity *NewBomb(void);
void BombTick(Entity *entity);
void BombRender(Entity *entity);
