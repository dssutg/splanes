#pragma once

#include "entity.h"

class Explosion : public Entity {
  public:
  Explosion(int32_t x, int32_t y);

  int32_t getZIndex() const;
  void tick();
  void render();
};
