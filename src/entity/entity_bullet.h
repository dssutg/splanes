#pragma once

#include "entity.h"

Entity *NewBullet(i32 type,
                  i32 x,
                  i32 y,
                  i32 xa,
                  i32 ya,
                  i32 ownertype,
                  i32 damage,
                  i32 bulletframeno);
void BulletTick(Entity *entity);
void BulletRender(Entity *entity);
