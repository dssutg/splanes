#include <stdio.h>
#include <stdlib.h>
#include <stdarg.h>

#include "util.h"

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
  void *newData = realloc(data, newByteCount);

  if (newData == NULL) {
    Fatalf("out of memory");
  }

  return newData;
}

void *Emalloc(i64 byteCount) {
  return Erealloc(NULL, byteCount);
}
