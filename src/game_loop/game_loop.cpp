#include "../lib/std.h"

#include "game_loop.h"

#include "../level/level.h"
#include "../gui/gui.h"
#include "../menu/menu.h"
#include "../keyboard_manager/keyboard_manager.h"
#include "../renderer/renderer.h"
#include "../sound_manager/sound_manager.h"
#include "../util/util.h"

void Restart() {
  level = std::make_unique<Level>(WindowWidth, WindowHeight);

  ResetToMenu(new MainMenu());
}

static void Tick() {
  if (keys[Key::Pause]) {
    keys[Key::Pause] = false;

    if (!HasMenus()) {
      PushMenu(new MainMenu());
    } else if (dynamic_cast<MainMenu *>(*TopMenu())) {
      PopMenu();
    }
  }

  if (TopMenu()) {
    (*TopMenu())->tick();
    return;
  }

  level->tick();
}

static void renderGUI() {
  const auto player = level->player;

  RenderHealthBar(player ? (*player)->getHealth() : 0);

  RenderProgressBar(
    {.x = 580, .y = 20, .w = 100, .h = 25},
    5,
    player ? (*player)->getBombTickTime() * 100 / Player::MaxBombTickTime : 0);

  RenderSmallLogo();

  RenderString(300,
               20,
               20,
               {
                 .r = 0xFF,
                 .g = 0xCA,
                 .b = 0x41,
                 .a = 0xFF,
               },
               0,
               0,
               "SCORE: %lu, DISTANCE: %lu",
               player ? (*player)->getScore() : 0,
               player ? (*player)->getDistance() : 0);

  if (TopMenu()) {
    (*TopMenu())->render();
  }
}

static void Render() {
  level->render();
  renderGUI();
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
    Fatal(std::format("can't init SDL: {}", SDL_GetError()));
  }

  window = SDL_CreateWindow(windowTitle.c_str(),
                            SDL_WINDOWPOS_CENTERED,
                            SDL_WINDOWPOS_CENTERED,
                            WindowWidth,
                            WindowHeight,
                            0);
  if (window == nullptr) {
    Fatal(std::format("can't create window: {}", SDL_GetError()));
  }

  renderer = SDL_CreateRenderer(
    window, -1, SDL_RENDERER_ACCELERATED | SDL_RENDERER_TARGETTEXTURE);
  if (renderer == nullptr) {
    Fatal(std::format("can't create renderer: {}", SDL_GetError()));
  }

  LoadTextures();
  InitSoundManager();

  Restart();

  DoGameLoop();

#ifndef NDEBUG
  FreeFontCache();
  FreeSoundManager();
  CloseAllMenus();
#endif

  SDL_DestroyWindow(window);
  SDL_Quit();
}
