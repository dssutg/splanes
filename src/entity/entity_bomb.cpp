#include "../lib/std.h"

#include "entity_bomb.h"
#include "entity_explosion.h"
#include "entity_ship.h"

#include "../game_loop/game_loop.h"
#include "../renderer/renderer.h"
#include "../level/level.h"

static constexpr auto frames = std::to_array<SDL_Rect>({
  {265, 265, 9, 21},
});

Bomb::Bomb() {
  pos.w = frames[0].w * 2;
  pos.h = frames[0].h * 2;

  pos.x = pos.x + (pos.w - pos.w) / 2 - 10;
  pos.y = pos.y - pos.h;

  tickTime = 1;
  tickTimeDelta = 1;

  damage = 1000;
}

SDL_Rect Bomb::calcSize() {
  const auto scale = pow(0.90, tickTime);

  auto w = static_cast<int32_t>(pos.w * scale);
  auto h = static_cast<int32_t>(pos.h * scale);

  if (w == 0 || h == 0) {
    tickTimeDelta = 0;
    w = 1;
    h = 1;
  }

  return {.x = pos.x, .y = pos.y, .w = w, .h = h};
}

void Bomb::tick() {
  tickTime += tickTimeDelta;

  pos.x += xa * 2;
  pos.y += ya * 2;

  if (tickTimeDelta != 0) {
    return;
  }

  const auto scaledrect = calcSize();

  for (auto &it : level->entities) {
    if (dynamic_cast<Ship *>(it) && it->collidesWith(scaledrect)) {
      it->hurt(damage);
      level->addEntity(*new Explosion(scaledrect.x, scaledrect.y));
      if (level->player) {
        (*level->player)->addScore(1);
      }
      break;
    }
  }

  removed = true;
}

void Bomb::render() {
  const auto f = &frames[0];
  const auto size = calcSize();
  RenderSprite(texture,
               {.x = pos.x, .y = pos.y, .w = size.w, .h = size.h},
               {.x = f->x, .y = f->y, .w = f->w, .h = f->h});
}
