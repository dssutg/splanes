#pragma once

#include "entity.h"

class Player : public Entity {
  public:
  static constexpr auto MaxHealth = 100;
  static constexpr auto MaxBombTickTime = 50;

  Player();

  int32_t getZIndex() const;
  void heal(int32_t healPoints);
  void die();
  void tick();
  void render();
  void addScore(uint64_t scoreToAdd);
  uint64_t getScore() const;
  uint64_t getDistance() const;
  int32_t getBombTickTime() const;

  private:
  uint64_t distance = 0;
  uint64_t score = 0;

  bool hasShot = false;
  bool hasBombed = false;

  int32_t deathTime = 0;
  int32_t bombTickTime = 0;

  void doDie();
};
