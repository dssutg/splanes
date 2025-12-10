#pragma once

#include <stddef.h>
#include <stdint.h>

#define ArrayLength(array) (sizeof(array) / sizeof((array)[0]))

extern const char *programName;

typedef uint8_t u8;
typedef uint16_t u16;
typedef uint32_t u32;
typedef uint64_t u64;
typedef int8_t i8;
typedef int16_t i16;
typedef int32_t i32;
typedef int64_t i64;
typedef float f32;
typedef double f64;

void Fatalf(const char *const format, ...);
void *Erealloc(void *data, size_t newByteCount);
void *Emalloc(size_t byteCount);
