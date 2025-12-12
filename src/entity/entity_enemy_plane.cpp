#include "../lib/std.h"

#include "entity_enemy_plane.h"
#include "entity_explosion.h"
#include "entity_bullet.h"

#include "../game_loop/game_loop.h"
#include "../renderer/renderer.h"

static constexpr auto frames = std::to_array<SDL_Rect>({
  {1, 1, 32, 32},
  {1, 34, 32, 32},
  {1, 67, 32, 32},
  {1, 100, 32, 32},
  {1, 133, 32, 32},
});

EnemyPlane::EnemyPlane() {
  crop = frames[static_cast<size_t>(rand()) % frames.size()];

  pos.w = crop.w * 2;
  pos.h = crop.h * 2;

  pos.x = rand() % WindowWidth;
  pos.y = -(rand() % WindowHeight) - pos.h;

  ya = 1;

  health = 100;
}

int32_t EnemyPlane::getZIndex() const {
  return 2;
}

void EnemyPlane::tick() {
  if (health <= 0) {
    level->addEntity(*new Explosion(pos.x, pos.y));
    removed = true;
    return;
  }

  if (!hasShot) {
    hasShot = true;
    const auto bulletX = pos.x + (pos.w - 128) / 2;
    const auto bulletY = pos.y + pos.w;
    level->addEntity(*new Bullet(bulletX, bulletY, 0, 2, false, 2, 1));
  } else {
    if (hasShot) {
      tickTime++;
      if (tickTime > 20) {
        hasShot = 0;
        tickTime = 0;
      }
    }
  }

  pos.x += xa * 20;
  pos.y += ya * 20;

  if (pos.y >= WindowHeight) {
    removed = true;
  }

  if (level->player && SDL_HasIntersection(&pos, &(*level->player)->pos)) {
    (*level->player)->die();
  }
}

void EnemyPlane::render() {
  Entity::render();
}
