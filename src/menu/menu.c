#include "menu.h"

#include "menu_table.h"

Menu aboutMenu;
Menu exitMenu;
Menu loseMenu;
Menu mainMenu;

MenuType prevMenuID;
MenuType menuID;

void TickMenu(void) {
  menuTable[menuID].Tick();
}

void RenderMenu(void) {
  menuTable[menuID].Render();
}
