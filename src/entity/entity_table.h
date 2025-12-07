#pragma once

#include "entity.h"

typedef struct EntityTableEntry {
  void (*Tick)(Entity *entity);
  void (*Render)(Entity *entity);
  const i32 zIndex;
} EntityTableEntry;

extern const EntityTableEntry entityTable[];
