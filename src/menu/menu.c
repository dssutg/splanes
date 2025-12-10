#include "menu.h"

Menu aboutMenu;
Menu exitMenu;
Menu loseMenu;
Menu mainMenu;

MenuType prevMenuID = MenuNone;
MenuType menuID = MenuNone;

void TickMenu(void) {
  menuTable[menuID].Tick();
}

void RenderMenu(void) {
  menuTable[menuID].Render();
}
