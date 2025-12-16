#include "menu.h"

#include "../keyboard_manager/keyboard_manager.h"

Menu aboutMenu;
Menu exitMenu;
Menu loseMenu;
Menu mainMenu;

MenuType prevMenuID = MenuNone;
MenuType menuID = MenuNone;

void TickMenu() {
  menuTable[menuID].Tick();
}

void RenderMenu() {
  menuTable[menuID].Render();
}

void MenuNoneTick() {
  // Nothing should be here.
}

void MenuNoneRender() {
  // Nothing should be here.
}

void HandleUpDownSelection(Menu *menu, i32 length) {
  if (SingleKeyPress(KeyUp)) {
    menu->selectedIndex--;
  }

  if (SingleKeyPress(KeyDown)) {
    menu->selectedIndex++;
  }

  if (exitMenu.selectedIndex >= length) {
    menu->selectedIndex = 0;
  }

  if (exitMenu.selectedIndex < 0) {
    menu->selectedIndex = length - 1;
  }
}
