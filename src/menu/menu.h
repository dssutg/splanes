#pragma once

#include "../util/util.h"

// Menus
typedef enum : u8 {
  MenuNone,
  MenuMain,
  MenuExit,
  MenuAbout,
  MenuLose,
} MenuType;

typedef struct {
  i32 selectedIndex;
} Menu;

typedef struct {
  void (*Tick)();
  void (*Render)();
} MenuTableEntry;

extern const MenuTableEntry menuTable[];

extern Menu aboutMenu;
extern Menu exitMenu;
extern Menu loseMenu;
extern Menu mainMenu;

extern MenuType prevMenuID;
extern MenuType menuID;

// Base Menu methods
void TickMenu();
void RenderMenu();
void HandleUpDownSelection(Menu *menu, i32 length);

// About Menu methods
void MenuAboutTick();
void MenuAboutRender();

// Exit Menu methods
void MenuExitTick();
void MenuExitRender();

// Lose Menu methods
void MenuLoseTick();
void MenuLoseRender();

// Main Menu methods
void MenuMainTick();
void MenuMainRender();

// None Menu methods
void MenuNoneTick();
void MenuNoneRender();
