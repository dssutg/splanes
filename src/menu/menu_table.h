#pragma once

#include "menu.h"

typedef struct MenuTableEntry {
  void (*Tick)(void);
  void (*Render)(void);
} MenuTableEntry;

extern const MenuTableEntry menuTable[];
