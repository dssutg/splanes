#include "../lib/std.h"

#include "entity_bomb.h"
#include "entity_bullet.h"
#include "entity_explosion.h"
#include "entity_player.h"

#include "../menu/menu.h"
#include "../renderer/renderer.h"
#include "../game_loop/game_loop.h"
#include "../level/level.h"
#include "../keyboard_manager/keyboard_manager.h"

Player::Player() {
  crop = {.x = 299, .y = 101, .w = 61, .h = 49};

  pos.w = crop.w * 2;
  pos.h = crop.h * 2;

  pos.x = (WindowWidth - pos.w) / 2;
  pos.y = WindowHeight - pos.h - 40;

  health = Player::MaxHealth;

  bombTickTime = Player::MaxBombTickTime;
}

int32_t Player::getZIndex() const {
  return 2;
}

void Player::die() {
  health = 0;
}

void Player::doDie() {
  deathTime++;

  if (deathTime == 1) {
    constexpr int32_t explosionCount = 10;
    for (int32_t i = 0; i < explosionCount; i++) {
      const auto x = pos.x + rand() % pos.w;
      const auto y = pos.y + rand() % (pos.h / 2);
      level->addEntity(*new Explosion(x, y));
    }
    return;
  }

  if (deathTime > 10) {
    PushMenu(new LoseMenu());
  }
}

void Player::heal(int32_t healPoints) {
  health += healPoints;
  if (health > Player::MaxHealth) {
    health = Player::MaxHealth;
  }
}

void Player::addScore(uint64_t scoreToAdd) {
  score += scoreToAdd;
}

uint64_t Player::getScore() const {
  return score;
}
uint64_t Player::getDistance() const {
  return distance;
}

int32_t Player::getBombTickTime() const {
  return bombTickTime;
}

void Player::tick() {
  if (health <= 0) {
    doDie();
    return;
  }

  xa = 0;
  if (keys[Key::Left]) {
    xa -= 20;
  }
  if (keys[Key::Right]) {
    xa += 20;
  }

  if (keys[Key::Bomb] && !hasBombed) {
    hasBombed = true;
    bombTickTime = 0;

    level->addEntity(*new Bomb());
  } else {
    if (hasBombed) {
      bombTickTime++;
      if (bombTickTime >= Player::MaxBombTickTime) {
        hasBombed = 0;
        bombTickTime = Player::MaxBombTickTime;
      }
    }
  }

  if (keys[Key::Shoot] && !hasShot) {
    hasShot = true;

    constexpr int32_t bulletDamage = 50;

    auto bullet = new Bullet(0, 0, 0, -1, true, bulletDamage, 0);
    bullet->pos.x = pos.x + (pos.w - bullet->pos.w) / 2 - 10;
    bullet->pos.y = pos.y - bullet->pos.h;
    level->addEntity(*bullet);
  } else if (hasShot) {
    tickTime++;
    if (tickTime > 3) {
      hasShot = false;
      tickTime = 0;
    }
  }

  auto xn = pos.x + xa;
  auto yn = pos.y + ya;

  if (xn + pos.w >= WindowWidth + 1) {
    xn = WindowWidth - pos.w;
  }

  xn = std::max(0, xn);

  pos.x = xn;
  pos.y = yn;

  distance++;
}

void Player::render() {
  if (deathTime < 5) {
    Entity::render();
  }
}
