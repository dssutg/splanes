#pragma once

#include "../lib/std.h"

#include "../level/level.h"

inline bool running = true;

inline std::unique_ptr<Level> level;

void Reset();
void Restart();
void RunGame();
