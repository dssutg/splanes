#pragma once

#include "entity.h"

Entity *NewEnemyPlane(void);
void EnemyPlaneTick(Entity *entity);
void EnemyPlaneRender(Entity *entity);
