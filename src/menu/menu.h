#pragma once

#include "../util/util.h"

// Menus
typedef enum MenuType {
  MenuNone,
  MenuMain,
  MenuExit,
  MenuAbout,
  MenuLose,
} MenuType;

typedef struct Menu {
  i32 selectedIndex;
} Menu;

extern Menu aboutMenu;
extern Menu exitMenu;
extern Menu loseMenu;
extern Menu mainMenu;

extern MenuType prevMenuID;
extern MenuType menuID;

void TickMenu(void);
void RenderMenu(void);
