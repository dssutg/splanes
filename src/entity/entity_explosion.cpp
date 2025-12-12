#include "../lib/std.h"

#include "entity_explosion.h"

#include "../renderer/renderer.h"
#include "../sound_manager/sound_manager.h"

static constexpr auto frames = std::to_array<SDL_Rect>({
  {67, 166, 32, 32},
  {100, 166, 32, 32},
  {133, 166, 32, 32},
  {166, 166, 32, 32},
  {199, 166, 32, 32},
  {232, 166, 32, 32},
});

Explosion::Explosion(int32_t x, int32_t y) {
  pos.x = x;
  pos.y = y;
  pos.w = frames[0].w * 2;
  pos.h = frames[0].h * 2;

  PlaySound(SoundID::Explosion1, 100);
}

int32_t Explosion::getZIndex() const {
  return 2;
}

void Explosion::tick() {
  tickTime++;
  if (tickTime >= frames.size()) {
    removed = true;
  }
}

void Explosion::render() {
  const auto f = &frames[tickTime % frames.size()];
  RenderSprite(texture,
               {.x = pos.x, .y = pos.y, .w = pos.w, .h = pos.h},
               {.x = f->x, .y = f->y, .w = f->w, .h = f->h});
}
