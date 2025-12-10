#include <stdio.h>
#include <stdlib.h>
#include <stdarg.h>

#include "util.h"

const char *programName = "splanes";

void Fatalf(const char *const format, ...) {
  fprintf(stderr, "%s: ", programName);

  va_list args;
  va_start(args, format);
  vfprintf(stderr, format, args);
  va_end(args);

  fprintf(stderr, "\n");

  exit(EXIT_FAILURE);
}

void *Erealloc(void *data, size_t newByteCount) {
  auto newData = realloc(data, newByteCount);

  if (newData == nullptr) {
    Fatalf("out of memory");
  }

  return newData;
}

void *Emalloc(size_t byteCount) {
  return Erealloc(nullptr, byteCount);
}
