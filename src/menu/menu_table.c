#include "menu.h"

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
