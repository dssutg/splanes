#pragma once

#include "entity.h"

class Ship : public Entity {
  public:
  Ship();

  int32_t getZIndex() const;
  void tick();
  void render();

  private:
  bool hasShot = false;
  bool hasBombed = false;
};
