#include "../lib/std.h"

#include "menu.h"

#include "../keyboard_manager/keyboard_manager.h"

static std::vector<Menu *> menuStack;

Menu::~Menu() {
}

void Menu::close() {
  PopMenu();
}

void Menu::openMenu(Menu *menu) {
  PushMenu(menu);
}

void Menu::handleUpAndDownSelection(int32_t itemCount) {
  if (SingleKeyPress(Key::Up)) {
    selectedIndex--;
  }

  if (SingleKeyPress(Key::Down)) {
    selectedIndex++;
  }

  if (selectedIndex >= itemCount) {
    selectedIndex = 0;
  }

  if (selectedIndex < 0) {
    selectedIndex = itemCount - 1;
  }
}

void PushMenu(Menu *menu) {
  menuStack.push_back(menu);
}

void PopMenu() {
  if (menuStack.empty()) {
    return;
  }

  delete menuStack.back();

  menuStack.pop_back();
}

bool HasMenus() {
  return !menuStack.empty();
}

std::optional<Menu *> TopMenu() {
  if (menuStack.empty()) {
    return std::nullopt;
  }
  return menuStack.back();
}

void CloseAllMenus() {
  for (auto menu : menuStack) {
    delete menu;
  }
  menuStack.clear();
}

void ResetToMenu(Menu *menu) {
  CloseAllMenus();
  PushMenu(menu);
}
