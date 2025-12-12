#pragma once

#include "entity.h"

class Bomb : public Entity {
  public:
  Bomb();

  void tick();
  void render();

  private:
  int32_t damage = 0;

  uint32_t tickTimeDelta = 1;

  SDL_Rect calcSize();
};
