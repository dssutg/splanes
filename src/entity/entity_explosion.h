#pragma once

#include "entity.h"

Entity *NewExplosion(i32 x, i32 y);
void ExplosionTick(Entity *entity);
void ExplosionRender(Entity *entity);
