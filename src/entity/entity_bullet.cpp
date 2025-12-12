#include "../lib/std.h"

#include "entity_bullet.h"
#include "entity_explosion.h"
#include "entity_enemy_plane.h"

#include "../level/level.h"
#include "../game_loop/game_loop.h"
#include "../renderer/renderer.h"

static constexpr auto frames = std::to_array<SDL_Rect>({
  {1, 166, 32, 32},
  {34, 199, 32, 32},
});

Bullet::Bullet(int32_t x,
               int32_t y,
               int32_t bulletXa,
               int32_t bulletYa,
               bool bulletOwnedByPlayer,
               int32_t bulletDamage,
               uint32_t bulletFrameNo) {
  frameNo = bulletFrameNo;

  pos.x = x;
  pos.y = y;
  pos.w = frames[frameNo].w * 2;
  pos.h = frames[frameNo].h * 2;

  xa = bulletXa;
  ya = bulletYa;

  ownedByPlayer = bulletOwnedByPlayer;
  damage = bulletDamage;
}

int32_t Bullet::getZIndex() const {
  return 2;
}

void Bullet::tick() {
  pos.x += xa * 20;
  pos.y += ya * 20;

  auto dead = false;

  for (auto &it : level->entities) {
    const auto canCollide = (ownedByPlayer && dynamic_cast<EnemyPlane *>(it)) ||
                            (!ownedByPlayer && dynamic_cast<Player *>(it));
    if (!canCollide) {
      continue;
    }

    if (!SDL_HasIntersection(&it->pos, &pos)) {
      continue;
    }

    dead = true;

    it->hurt(damage);

    if (dynamic_cast<Player *>(it)) {
      level->addEntity(*new Explosion(pos.x, pos.y));
      continue;
    }

    if (ownedByPlayer && level->player) {
      (*level->player)->addScore(1);
    }
  }

  const SDL_Rect windowRect = {
    .x = 0,
    .y = 0,
    .w = WindowWidth,
    .h = WindowHeight,
  };

  if (dead || !SDL_HasIntersection(&pos, &windowRect)) {
    removed = true;
  }
}

void Bullet::render() {
  const auto f = &frames[frameNo];
  RenderSprite(texture,
               {.x = pos.x, .y = pos.y, .w = pos.w, .h = pos.h},
               {.x = f->x, .y = f->y, .w = f->w, .h = f->h});
}
