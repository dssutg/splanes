#pragma once

#include "entity.h"

Entity *NewHealer(void);
void HealerTick(Entity *entity);
void HealerRender(Entity *entity);
