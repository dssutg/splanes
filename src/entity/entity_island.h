#pragma once

#include "entity.h"

class Island : public Entity {
  public:
  Island();

  void tick();
  void render();

  private:
  uint32_t frameNo = 0;
};
