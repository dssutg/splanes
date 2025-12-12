#pragma once

#include "entity.h"

class Healer : public Entity {
  public:
  Healer();

  int32_t getZIndex() const;
  void tick();
  void render();

  private:
  uint32_t frameNo = 0;
};
