# AGENTS.md - Splanes

A WW2-style scroll shooter game written in Go using go-sdl2.

## Run

```bash
go run .
# or:
make run
```

## Build

```bash
make build    # builds ./splanes binary
```

## Code Quality

```bash
make lint     # golangci-lint run ./...
make fmt      # gofumpt -w .
```

**Linter config**: `.golangci.yml` - gofumpt formatting, staticcheck, gocyclo (min complexity 15), misspell, unconvert.
Excludes: G404 (math/rand for games), G115 (int conversions), exitAfterDefer, errcheck.

## Testing

No tests exist. `go test ./...` will pass with no test files.

## Dependencies

- `github.com/veandco/go-sdl2 v0.4.40` (SDL2 bindings)
- Requires SDL2, SDL2_mixer, and SDL2_ttf system libraries

## Project Structure

- `main.go` - Entry point, SDL init, game loop (fixed 20 ticks/sec)
- `entity.go` - Entity struct, entity pool (100 fixed slots), dispatch table
- `entity_*.go` - Game objects (player, enemies, bullets, items)
- `menu.go` - Menu system, dispatch table
- `menu_*.go` - Menu screens (main, lose, about, exit)
- `gui_*.go` - UI (health bar, progress bar, text, fonts)
- `renderer.go`, `sound.go` - Rendering and audio
- `keyboard.go` - Input handling, key mapping
- `util.go` - Math helpers (Clamp, Rotate, RandIntRange)
- `assets/` - Sprites, fonts, sfx, music

## Architecture

### Entity Pool System
- Fixed 100-slot pool (`EntityPool [100]Entity`)
- Entities allocated with `NewEntity(etype)` - finds first free slot
- Removed entities zeroed (Kind = EntityTypeNone)
- Dispatch via `entityTable` map[EntityType]EntityTableEntry

### Game Loop
- Fixed 20 ticks/sec independent of render framerate
- Uses SDL high-resolution timer for precise timing
- Render runs at display refresh rate (typically 60 FPS)
- Two-layer parallax water scrolling

### Menu System
- `menuID` tracks current menu state
- ESC toggles between game and main menu
- Dispatch via `menuTable` map[MenuType]MenuTableEntry

### Entity Types
| Type | Z | Behavior |
|------|---|---------|
| Player | 2 | Controlled by arrow keys, shoots bullets, drops bombs |
| EnemyPlane | 2 | Flies down, shoots at player |
| Bullet | 2 | Travels in direction, damages on collision |
| Bomb | 2 | Falls, shrinks, damages ships/submarines |
| Ship | 1 | Slow, moves down, no attack |
| Submarine | 1 | State machine: rising/idle/diving |
| Healer | 2 | Restores 20 health on pickup |
| Island | 0 | Decorative, scrolls slowly |
| Explosion | 2 | Animation effect, removed when done |

## Controls

| Key | Action |
|-----|-------|
| Arrow Keys / WASD | Move |
| Space | Shoot |
| X | Drop bomb |
| Q/E | Rotate |
| ESC | Pause / Menu |
| Enter | Confirm |

## Key Constants

- Window: 800x600
- Player max health: 100
- Bomb reload: 50 ticks
- Game tick rate: 20/sec

## Code Conventions

- Uses `math/rand/v2` (Go 1.22+)
- Single package, flat file structure (no subdirs)
- Generic math/rand is fine - game code doesn't need crypto/rand
- int->int32 conversions are intentional

## Binary

- Output: `./splanes` (in .gitignore)
- Window title: "Splanes"