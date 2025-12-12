#pragma once

#include "entity.h"

class Bullet : public Entity {
  public:
  Bullet(int32_t x,
         int32_t y,
         int32_t xa,
         int32_t ya,
         bool ownedByPlayer,
         int32_t damage,
         uint32_t frameNo);

  int32_t getZIndex() const;
  void tick();
  void render();

  private:
  bool ownedByPlayer = false;
  int32_t damage = 0;
  uint32_t frameNo = 0;
};
