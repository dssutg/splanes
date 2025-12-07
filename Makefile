# Toolchain
CC ?= gcc
AR ?= ar
RM ?= rm -f
MKDIR_P ?= mkdir -p

# Project
TARGET := splanes
SRC_DIR := src
BUILD_DIR := build
BIN_DIR := bin

# Recursive source list
SRC := $(shell find $(SRC_DIR) -name '*.c')
OBJ := $(patsubst $(SRC_DIR)/%.c,$(BUILD_DIR)/%.o,$(SRC))
DEPS := $(OBJ:.o=.d)

# Common flags
COMMON_CFLAGS := -Wall -Werror -O2 -g -Iinclude -std=c2x -pedantic
COMMON_LDFLAGS := -lm

# dependency flags: generate .d files (exclude system headers) + phony targets
DEPFLAGS := -MMD -MP

# Platform selection (unchanged)
PLAT ?= auto

ifeq ($(PLAT),auto)
UNAME_S := $(shell uname -s 2>/dev/null || echo Unknown)
ifeq ($(UNAME_S),Linux)
PLAT := linux
endif
ifeq ($(UNAME_S),Darwin)
PLAT := macos
endif
ifeq ($(findstring MINGW,$(UNAME_S)),MINGW)
PLAT := windows
endif
ifeq ($(PLAT),auto)
PLAT := linux
endif
endif

ifeq ($(PLAT),linux)
CFLAGS := $(COMMON_CFLAGS) `sdl2-config --cflags`
LDFLAGS := `sdl2-config --libs` -lSDL2_ttf -lSDL2_mixer -lSDL2_image $(COMMON_LDFLAGS) -lGL -lGLU
PLATFORM_TAG := linux
endif

ifeq ($(PLAT),macos)
CFLAGS := $(COMMON_CFLAGS) `sdl2-config --cflags`
LDFLAGS := `sdl2-config --libs` -lSDL2_ttf -lSDL2_mixer -lSDL2_image $(COMMON_LDFLAGS) -framework OpenGL
PLATFORM_TAG := macos
endif

ifeq ($(PLAT),windows)
CFLAGS := $(COMMON_CFLAGS) `sdl2-config --cflags` -DWIN32
LDFLAGS := `sdl2-config --libs` -lSDL2_ttf -lSDL2_mixer -lSDL2_image $(COMMON_LDFLAGS) -mwindows -lmingw32 -lopengl32 -lglu32
PLATFORM_TAG := windows
endif

# Phony targets
.PHONY: all clean rebuild distclean dirs linux macos windows check format

all: dirs $(BIN_DIR)/$(PLATFORM_TAG)/$(TARGET)

run: all
	./$(BIN_DIR)/$(PLATFORM_TAG)/$(TARGET)

dirs:
	@$(MKDIR_P) $(dir $(OBJ))
	@$(MKDIR_P) $(BIN_DIR)/$(PLATFORM_TAG)

# Include dependency files (ignore missing on first build)
-include $(DEPS)

# Link
$(BIN_DIR)/$(PLATFORM_TAG)/$(TARGET): $(OBJ)
	$(CC) $(CFLAGS) -o $@ $^ $(LDFLAGS)

# Compile object files and generate deps. Make each .o also depend on its .d to avoid races with -j.
# The compiler invocation produces the .d file as a side-effect.
$(BUILD_DIR)/%.o: $(SRC_DIR)/%.c $(BUILD_DIR)/%.d
	@$(MKDIR_P) $(dir $@)
	$(CC) $(CFLAGS) $(DEPFLAGS) -c $< -o $@

# A pattern rule to create the .d file if make asks for it directly.
# This runs the compiler to produce .d (and a .o as side-effect if necessary).
$(BUILD_DIR)/%.d: ;

# Convenience per-platform targets
linux:
	$(MAKE) PLAT=linux all

macos:
	$(MAKE) PLAT=macos all

windows:
	$(MAKE) PLAT=windows all

# Clean rules
clean:
	$(RM) $(OBJ) $(DEPS)
	-@rmdir --ignore-fail-on-non-empty -p $(sort $(dir $(OBJ))) 2>/dev/null || true

distclean: clean
	$(RM) $(BIN_DIR)/$()/$(TARGET)

# Rebuild
rebuild: distclean all

check: format

format:
	find . -name '*.c' -or -name '*.h' | xargs clang-format -i
