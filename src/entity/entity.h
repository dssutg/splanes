#pragma once

#include "../lib/std.h"

// Entity base class
class Entity {
  public:
  bool removed = false;

  SDL_Rect pos = {};

  virtual ~Entity();

  virtual int32_t getZIndex() const;
  virtual void hurt(int32_t damageToTake);
  virtual int32_t getHealth() const;
  virtual bool collidesWith(const SDL_Rect &rect) const;
  virtual void tick();
  virtual void render();

  protected:
  int32_t xa = 0;
  int32_t ya = 0;

  SDL_Rect crop = {};
  int32_t texture = 0;

  int32_t health = 0;

  uint32_t tickTime = 0;
};
