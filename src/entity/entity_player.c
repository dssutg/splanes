#include "entity.h"

#include "../renderer/renderer.h"
#include "../keyboard_manager/keyboard_manager.h"
#include "../menu/menu.h"

Entity *NewPlayer(void) {
  auto player = NewEntity(EntityPlayer);

  player->texture = 0;

  player->xa = 0;
  player->ya = 0;

  player->crop.x = 299;
  player->crop.y = 101;
  player->crop.w = 61;
  player->crop.h = 49;

  player->pos.w = player->crop.w * 2;
  player->pos.h = player->crop.h * 2;
  player->pos.x = (WindowWidth - player->pos.w) / 2;
  player->pos.y = WindowHeight - player->pos.h - 40;

  player->hasShot = false;
  player->hasBombed = false;

  player->health = MaxPlayerHealth;

  player->score = 0;
  player->distance = 0;
  player->deathTime = 0;
  player->bombTickTime = PlayerMaxBombTickTime;

  return player;
}

void PlayerDoDie(Entity *player) {
  player->deathTime++;

  if (player->deathTime == 1) {
    constexpr i32 explosionCount = 10;
    for (i32 i = 0; i < explosionCount; i++) {
      const auto x = player->pos.x + rand() % player->pos.w;
      const auto y = player->pos.y + rand() % (player->pos.h / 2);
      NewExplosion(x, y);
    }
    return;
  }

  if (player->deathTime > 10) {
    menuID = MenuLose;
  }
}

void HealPlayer(i32 healPoints) {
  player->health += healPoints;
  if (player->health > MaxPlayerHealth) {
    player->health = MaxPlayerHealth;
  }
}

void PlayerTick(Entity *entity) {
  if (entity->health <= 0) {
    PlayerDoDie(entity);
    return;
  }

  entity->xa = 0;

  if (keys[KeyLeft]) {
    entity->xa = -20;
  }

  if (keys[KeyRight]) {
    entity->xa = 20;
  }

  if (keys[KeyBomb] && !entity->hasBombed) {
    entity->hasBombed = 1;
    entity->bombTickTime = 0;

    auto bomb = NewBomb();
    bomb->pos.x = entity->pos.x + (entity->pos.w - bomb->pos.w) / 2 - 10;
    bomb->pos.y = entity->pos.y - bomb->pos.h;
  } else {
    if (entity->hasBombed) {
      entity->bombTickTime++;

      if (entity->bombTickTime >= PlayerMaxBombTickTime) {
        entity->hasBombed = 0;
        entity->bombTickTime = PlayerMaxBombTickTime;
      }
    }
  }

  if (keys[KeyShoot] && !entity->hasShot) {
    entity->hasShot = 1;

    constexpr i32 damage = 50;

    auto bullet = NewBullet(EntityBullet, 0, 0, 0, -1, entity->type, damage, 0);
    bullet->pos.x = entity->pos.x + (entity->pos.w - bullet->pos.w) / 2 - 10;
    bullet->pos.y = entity->pos.y - bullet->pos.h;
  } else if (entity->hasShot) {
    entity->tickTime++;

    if (entity->tickTime > 3) {
      entity->hasShot = 0;
      entity->tickTime = 0;
    }
  }

  auto xn = entity->pos.x + entity->xa;
  auto yn = entity->pos.y + entity->ya;

  if (xn + entity->pos.w - 1 >= WindowWidth) {
    xn = WindowWidth - entity->pos.w;
  }

  if (xn < 0) {
    xn = 0;
  }

  entity->pos.x = xn;
  entity->pos.y = yn;

  entity->distance++;
}

void PlayerRender(Entity *entity) {
  if (entity->deathTime < 5) {
    RenderEntitySprite(entity);
  }
}
