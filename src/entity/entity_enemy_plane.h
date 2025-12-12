#pragma once

#include "entity.h"

class EnemyPlane : public Entity {
  public:
  EnemyPlane();

  int32_t getZIndex() const;
  void tick();
  void render();

  private:
  bool hasShot = false;
  bool hasBombed = false;
};
