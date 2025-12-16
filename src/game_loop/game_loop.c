#include <SDL2/SDL.h>
#include <SDL2/SDL_events.h>
#include <SDL2/SDL_mixer.h>
#include <SDL2/SDL_ttf.h>

#include "game_loop.h"

#include "../entity/entity.h"
#include "../gui/gui.h"
#include "../keyboard_manager/keyboard_manager.h"
#include "../menu/menu.h"
#include "../renderer/renderer.h"
#include "../sound_manager/sound_manager.h"
#include "../util/util.h"

bool running = true;

i32 layer1;
i32 layer2;

void Reset() {
  player = NewPlayer();

  layer1 = -WindowHeight;
  layer2 = 0;

  menuID = MenuMain;

  mainMenu.selectedIndex = 0;
  exitMenu.selectedIndex = 0;
  aboutMenu.selectedIndex = 0;
  aboutMenu.selectedIndex = 0;

  PlayMusic(MusicBackground0, 70);
}

void Restart() {
  RemoveAllEntities();
  Reset();
}

static void Tick() {
  if (keys[KeyPause]) {
    keys[KeyPause] = false;

    if (menuID == MenuNone) {
      menuID = MenuMain;
    } else if (menuID == MenuMain) {
      menuID = MenuNone;
    }
  }

  if (menuID != MenuNone) {
    TickMenu();
    return;
  }

  if (rand() % 20 == 0) {
    NewEnemyPlane();
  }
  if (rand() % 80 == 0) {
    NewShip();
  }
  if (rand() % 30 == 0) {
    NewIsland();
  }
  if (rand() % 100 == 0) {
    NewHealer();
  }

  layer1 += 10;
  layer2 += 10;
  if (layer2 >= WindowHeight) {
    layer2 = layer1 - WindowHeight;
    const auto tmp = layer1;
    layer1 = layer2;
    layer2 = tmp;
  }

  for (auto entity = entities; entity != nullptr; entity = entity->next) {
    entityTable[entity->type].Tick(entity);
  }

  // delete all dead entities
  Entity *nextEntity = nullptr;
  for (auto entity = entities; entity != nullptr; entity = nextEntity) {
    nextEntity = entity->next;
    if (entity->removed) {
      FreeEntity(entity);
    }
  }
}

static void RenderLayer(i32 offsetY) {
  constexpr auto tileWidth = (WindowWidth + TileSize - 1) / TileSize;
  constexpr auto tileHeight = (WindowHeight + TileSize - 1) / TileSize;

  for (i32 tileY = 0; tileY < tileHeight; tileY++) {
    for (i32 tileX = 0; tileX < tileWidth; tileX++) {
      RenderSprite(0,
                   tileX * TileSize,
                   tileY * TileSize + offsetY,
                   TileSize,
                   TileSize,
                   265,
                   364,
                   32,
                   32);
    }
  }
}

static void Render() {
  RenderLayer(layer1);
  RenderLayer(layer2);

  // Render entitites
  for (i32 zIndex = 0; zIndex <= 2; zIndex++) {
    for (auto entity = entities; entity != nullptr; entity = entity->next) {
      const auto entry = &entityTable[entity->type];
      if (entry->zIndex == zIndex) {
        entry->Render(entity);
      }
    }
  }

  // Render GUI
  RenderHealthBar(player->health);

  RenderProgressBar(
    580, 20, 100, 25, 5, player->bombTickTime * 100 / PlayerMaxBombTickTime);

  RenderSmallLogo();

  RenderString(300,
               20,
               20,
               0xFF,
               0xCA,
               0x41,
               0xFF,
               0,
               0,
               "SCORE: %lu, DISTANCE: %lu",
               player->score,
               player->distance);

  RenderMenu();
}

static void DoGameLoop() {
  while (running) {
    SDL_Event event;

    while (SDL_PollEvent(&event)) {
      switch (event.type) {
      case SDL_QUIT:
        running = false;
        break;

      case SDL_KEYDOWN:
        UpdateKey(event.key.keysym.sym, true);
        break;

      case SDL_KEYUP:
        UpdateKey(event.key.keysym.sym, false);
        break;
      }
    }

    Tick();

    SDL_SetRenderDrawColor(renderer, 0, 0, 0, 0);
    SDL_RenderClear(renderer);

    Render();

    SDL_RenderPresent(renderer);
    SDL_Delay(WindowDelayMilliseconds);
  }
}

void RunGame() {
  if (SDL_Init(SDL_INIT_EVERYTHING) != 0 || TTF_Init() != 0 ||
      SDL_Init(SDL_INIT_AUDIO) == -1 ||
      Mix_OpenAudio(44100, MIX_DEFAULT_FORMAT, 2, 4096) == -1) {
    Fatalf("can't init SDL: %s", SDL_GetError());
  }

  window = SDL_CreateWindow(windowTitle,
                            SDL_WINDOWPOS_CENTERED,
                            SDL_WINDOWPOS_CENTERED,
                            WindowWidth,
                            WindowHeight,
                            0);
  if (window == nullptr) {
    Fatalf("can't create window: %s", SDL_GetError());
  }

  renderer = SDL_CreateRenderer(
    window, -1, SDL_RENDERER_ACCELERATED | SDL_RENDERER_TARGETTEXTURE);
  if (renderer == nullptr) {
    Fatalf("can't create renderer: %s", SDL_GetError());
  }

  LoadTextures();
  InitSoundManager();

  InitKeyboardManager();

  Reset();

  DoGameLoop();

  SDL_DestroyWindow(window);
  SDL_Quit();
}
