#pragma once

#include "../lib/std.h"

class Menu {
  public:
  int32_t selectedIndex = 0;

  virtual ~Menu();

  void close();
  void openMenu(Menu *menu);
  void handleUpAndDownSelection(int32_t itemCount);

  virtual void tick() = 0;
  virtual void render() = 0;
};

class MainMenu : public Menu {
  void tick() override;
  void render() override;
};

class ExitMenu : public Menu {
  void tick() override;
  void render() override;
};

class AboutMenu : public Menu {
  void tick() override;
  void render() override;
};

class LoseMenu : public Menu {
  void tick() override;
  void render() override;
};

void PushMenu(Menu *menu);
void ResetToMenu(Menu *menu);
void PopMenu();
void CloseAllMenus();
bool HasMenus();
std::optional<Menu *> TopMenu();
