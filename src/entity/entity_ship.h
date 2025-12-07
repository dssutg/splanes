#pragma once

#include "entity.h"

Entity *NewShip(void);
void ShipTick(Entity *entity);
void ShipRender(Entity *entity);
