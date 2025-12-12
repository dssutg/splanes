#include "entity.h"

#include "../sound_manager/sound_manager.h"
#include "../renderer/renderer.h"

Entity::~Entity() {
}

int32_t Entity::getZIndex() const {
  return 0;
}

void Entity::hurt(int32_t damageToTake) {
  health = health > damageToTake ? health - damageToTake : 0;
  PlaySound(SoundID::Hurt, 100);
}

bool Entity::collidesWith(const SDL_Rect &rect) const {
  return SDL_HasIntersection(&pos, &rect);
}

int32_t Entity::getHealth() const {
  return health;
}

void Entity::tick() {
}

void Entity::render() {
  RenderSprite(texture,
               {.x = pos.x, .y = pos.y, .w = pos.w, .h = pos.h},
               {.x = crop.x, .y = crop.y, .w = crop.w, .h = crop.h});
}
