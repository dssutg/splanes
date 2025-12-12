#pragma once

#include "../lib/std.h"

inline std::string programName = "splanes";

[[noreturn]] void Fatal(const std::string &msg);
