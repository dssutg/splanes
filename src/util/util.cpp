#include "../lib/std.h"

#include "util.h"

[[noreturn]] void Fatal(const std::string &msg) {
  std::cerr << programName << ": " << msg << '\n';
  std::exit(EXIT_FAILURE);
}
