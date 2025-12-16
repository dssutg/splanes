#include <asm-generic/errno.h>
#include <stdio.h>
#include <stdlib.h>
#include <stdarg.h>

#include "util.h"

typedef unsigned long Size;

const char *programName = "splanes";

void Fatalf(const char *format, ...) {
  fprintf(stderr, "%s: ", programName);

  va_list args;
  va_start(args, format);
  vfprintf(stderr, format, args);
  va_end(args);

  fprintf(stderr, "\n");

  exit(EXIT_FAILURE);
}

void *Erealloc(void *data, i64 newByteCount) {
  if (newByteCount < 0) {
    newByteCount = 0;
  }

  auto newData = realloc(data, (Size)newByteCount);

  if (newData == nullptr) {
    Fatalf("out of memory");
  }

  return newData;
}

void *Emalloc(i64 byteCount) {
  return Erealloc(nullptr, byteCount);
}
