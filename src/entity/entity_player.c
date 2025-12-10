#include "entity.h"

#include "../renderer/renderer.h"
#include "../keyboard_manager/keyboard_manager.h"
#include "../menu/menu.h"

Entity *NewPlayer(void) {
  auto p = NewEntity(EntityPlayer);

  p->texture = 0;

  p->xa = 0;
  p->ya = 0;

  p->crop.x = 299;
  p->crop.y = 101;
  p->crop.w = 61;
  p->crop.h = 49;

  p->pos.w = p->crop.w * 2;
  p->pos.h = p->crop.h * 2;
  p->pos.x = (WindowWidth - p->pos.w) / 2;
  p->pos.y = WindowHeight - p->pos.h - 40;

  p->hasShot = false;
  p->hasBombed = false;

  p->health = MaxPlayerHealth;

  p->score = 0;
  p->distance = 0;
  p->deathTime = 0;
  p->bombTickTime = PlayerMaxBombTickTime;

  return p;
}

void PlayerDoDie(Entity *p) {
  p->deathTime++;

  if (p->deathTime == 1) {
    constexpr i32 explosionCount = 10;
    for (i32 i = 0; i < explosionCount; i++) {
      const auto x = p->pos.x + rand() % p->pos.w;
      const auto y = p->pos.y + rand() % (p->pos.h / 2);
      NewExplosion(x, y);
    }
    return;
  }

  if (p->deathTime > 10) {
    menuID = MenuLose;
  }
}

void HealPlayer(i32 healPoints) {
  player->health += healPoints;
  if (player->health > MaxPlayerHealth) {
    player->health = MaxPlayerHealth;
  }
}

void PlayerTick(Entity *p) {
  if (p->health <= 0) {
    PlayerDoDie(p);
    return;
  }

  p->xa = 0;

  if (keys[KeyLeft]) {
    p->xa = -20;
  }

  if (keys[KeyRight]) {
    p->xa = 20;
  }

  if (keys[KeyBomb] && !p->hasBombed) {
    p->hasBombed = 1;
    p->bombTickTime = 0;

    auto bomb = NewBomb();
    bomb->pos.x = p->pos.x + (p->pos.w - bomb->pos.w) / 2 - 10;
    bomb->pos.y = p->pos.y - bomb->pos.h;
  } else {
    if (p->hasBombed) {
      p->bombTickTime++;

      if (p->bombTickTime >= PlayerMaxBombTickTime) {
        p->hasBombed = 0;
        p->bombTickTime = PlayerMaxBombTickTime;
      }
    }
  }

  if (keys[KeyShoot] && !p->hasShot) {
    p->hasShot = 1;

    constexpr i32 damage = 50;

    auto bullet = NewBullet(EntityBullet, 0, 0, 0, -1, p->type, damage, 0);
    bullet->pos.x = p->pos.x + (p->pos.w - bullet->pos.w) / 2 - 10;
    bullet->pos.y = p->pos.y - bullet->pos.h;
  } else if (p->hasShot) {
    p->tickTime++;

    if (p->tickTime > 3) {
      p->hasShot = 0;
      p->tickTime = 0;
    }
  }

  auto xn = p->pos.x + p->xa;
  auto yn = p->pos.y + p->ya;

  if (xn + p->pos.w >= WindowWidth + 1) {
    xn = WindowWidth - p->pos.w;
  }

  if (xn < 0) {
    xn = 0;
  }

  p->pos.x = xn;
  p->pos.y = yn;

  p->distance++;
}

void PlayerRender(Entity *p) {
  if (p->deathTime < 5) {
    RenderEntitySprite(p);
  }
}
