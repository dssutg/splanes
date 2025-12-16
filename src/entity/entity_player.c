#include "entity.h"

#include "../renderer/renderer.h"
#include "../keyboard_manager/keyboard_manager.h"
#include "../menu/menu.h"

Entity *NewPlayer() {
  auto e = NewEntity(EntityPlayer);

  e->texture = 0;

  e->xa = 0;
  e->ya = 0;

  e->crop.x = 299;
  e->crop.y = 101;
  e->crop.w = 61;
  e->crop.h = 49;

  e->pos.w = e->crop.w * 2;
  e->pos.h = e->crop.h * 2;
  e->pos.x = (WindowWidth - e->pos.w) / 2;
  e->pos.y = WindowHeight - e->pos.h - 40;

  e->hasShot = false;
  e->hasBombed = false;

  e->health = MaxPlayerHealth;

  e->score = 0;
  e->distance = 0;
  e->deathTime = 0;
  e->bombTickTime = PlayerMaxBombTickTime;

  return e;
}

void PlayerDoDie(Entity *e) {
  e->deathTime++;

  if (e->deathTime == 1) {
    constexpr i32 explosionCount = 10;
    for (i32 i = 0; i < explosionCount; i++) {
      const auto x = e->pos.x + rand() % e->pos.w;
      const auto y = e->pos.y + rand() % (e->pos.h / 2);
      NewExplosion(x, y);
    }
    return;
  }

  if (e->deathTime > 10) {
    menuID = MenuLose;
  }
}

void HealPlayer(i32 healPoints) {
  player->health += healPoints;
  if (player->health > MaxPlayerHealth) {
    player->health = MaxPlayerHealth;
  }
}

void PlayerTick(Entity *e) {
  if (e->health <= 0) {
    PlayerDoDie(e);
    return;
  }

  e->xa = 0;

  if (keys[KeyLeft]) {
    e->xa = -20;
  }

  if (keys[KeyRight]) {
    e->xa = 20;
  }

  if (keys[KeyBomb] && !e->hasBombed) {
    e->hasBombed = 1;
    e->bombTickTime = 0;

    auto bomb = NewBomb();
    bomb->pos.x = e->pos.x + (e->pos.w - bomb->pos.w) / 2 - 10;
    bomb->pos.y = e->pos.y - bomb->pos.h;
  } else {
    if (e->hasBombed) {
      e->bombTickTime++;
      if (e->bombTickTime >= PlayerMaxBombTickTime) {
        e->hasBombed = 0;
        e->bombTickTime = PlayerMaxBombTickTime;
      }
    }
  }

  if (keys[KeyShoot] && !e->hasShot) {
    e->hasShot = 1;

    constexpr i32 damage = 50;

    auto bullet = NewBullet(EntityBullet, 0, 0, 0, -1, e->type, damage, 0);
    bullet->pos.x = e->pos.x + (e->pos.w - bullet->pos.w) / 2 - 10;
    bullet->pos.y = e->pos.y - bullet->pos.h;
  } else if (e->hasShot) {
    e->tickTime++;

    if (e->tickTime > 3) {
      e->hasShot = 0;
      e->tickTime = 0;
    }
  }

  auto xn = e->pos.x + e->xa;
  auto yn = e->pos.y + e->ya;

  if (xn + e->pos.w >= WindowWidth + 1) {
    xn = WindowWidth - e->pos.w;
  }

  if (xn < 0) {
    xn = 0;
  }

  e->pos.x = xn;
  e->pos.y = yn;

  e->distance++;
}

void PlayerRender(Entity *e) {
  if (e->deathTime < 5) {
    RenderEntitySprite(e);
  }
}
