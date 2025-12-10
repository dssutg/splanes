#pragma once

#include "../util/util.h"

// Menus
typedef enum MenuType : u8 {
  MenuNone,
  MenuMain,
  MenuExit,
  MenuAbout,
  MenuLose,
} MenuType;

typedef struct Menu {
  i32 selectedIndex;
} Menu;

typedef struct MenuTableEntry {
  void (*Tick)(void);
  void (*Render)(void);
} MenuTableEntry;

extern const MenuTableEntry menuTable[];

extern Menu aboutMenu;
extern Menu exitMenu;
extern Menu loseMenu;
extern Menu mainMenu;

extern MenuType prevMenuID;
extern MenuType menuID;

// Base Menu methods
void TickMenu(void);
void RenderMenu(void);

// About Menu methods
void MenuAboutTick(void);
void MenuAboutRender(void);

// Exit Menu methods
void MenuExitTick(void);
void MenuExitRender(void);

// Lose Menu methods
void MenuLoseTick(void);
void MenuLoseRender(void);

// Main Menu methods
void MenuMainTick(void);
void MenuMainRender(void);

// None Menu methods
void MenuNoneTick(void);
void MenuNoneRender(void);
