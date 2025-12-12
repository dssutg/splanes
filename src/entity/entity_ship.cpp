#include "../lib/std.h"

#include "entity_ship.h"
#include "entity_explosion.h"

#include "../game_loop/game_loop.h"
#include "../renderer/renderer.h"
#include "../level/level.h"

static constexpr auto frames = std::to_array<SDL_Rect>({
  {505, 298, 41, 197},
  {463, 298, 41, 197},
});

Ship::Ship() {
  const auto frame = &frames[0];

  pos.w = frame->w * 1;
  pos.h = frame->h * 1;
  pos.x = rand() % WindowWidth;
  pos.y = -(rand() % WindowHeight) - pos.h;

  ya = 1;

  health = 100;
}

int32_t Ship::getZIndex() const {
  return 1;
}

void Ship::tick() {
  tickTime++;
  if (tickTime > 10) {
    tickTime = 0;
  }

  if (health <= 0) {
    level->addEntity(*new Explosion(pos.x, pos.y));
    removed = true;
    return;
  }

  pos.x += xa * 11;
  pos.y += ya * 11;

  if (pos.y >= WindowHeight) {
    removed = true;
  }
}

void Ship::render() {
  const auto f = &frames[tickTime / 5 % frames.size()];
  RenderSprite(texture,
               {.x = pos.x, .y = pos.y, .w = pos.w, .h = pos.h},
               {.x = f->x, .y = f->y, .w = f->w, .h = f->h});
}
