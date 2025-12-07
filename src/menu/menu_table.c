#include "menu_table.h"

#include "menu_about.h"
#include "menu_exit.h"
#include "menu_lose.h"
#include "menu_main.h"
#include "menu_none.h"

const MenuTableEntry menuTable[] = {
  [MenuNone] =
    {
      .Tick = MenuNoneTick,
      .Render = MenuNoneRender,
    },

  [MenuMain] =
    {
      .Tick = MenuMainTick,
      .Render = MenuMainRender,
    },

  [MenuExit] =
    {
      .Tick = MenuExitTick,
      .Render = MenuExitRender,
    },

  [MenuAbout] =
    {
      .Tick = MenuAboutTick,
      .Render = MenuAboutRender,
    },

  [MenuLose] =
    {
      .Tick = MenuLoseTick,
      .Render = MenuLoseRender,
    },
};
