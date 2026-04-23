# Splanes

![Demo Screenshot](doc/images/demo-screenshot.png)

A WW2-style scroll shooter game built with Go and SDL2.

## Gameplay

- Control a fighter plane flying over the ocean
- Survive waves of enemy planes, submarines, and bombers
- Collect health pickups and avoid enemy bombs
- Score points by destroying enemies
- Press ESC at any time to pause and access the menu

## Controls

| Key | Action |
|-----|--------|
| Arrow Keys | Move plane |
| Space | Drop bomb (limited ammo) |
| Enter/Return | Confirm selection |
| ESC | Pause / Menu |

## Build & Run

```bash
make run      # Run the game
make build    # Build binary
make lint     # Run linter
make fmt      # Format code
```

## Requirements

- Go 1.26+
- SDL2, SDL2_mixer, SDL2_ttf system libraries

## Project History

Written by Daniil Stepanov in 2019. Originally a Java game (2015-2016), then ported to C. Ported to Go on April 22nd, 2026.

## Project Structure

- `main.go` - Game loop and SDL initialization
- `entity_*.go` - Game entities (player, enemies, bullets, explosions)
- `menu_*.go` - Menu screens (main, pause, game over, about)
- `gui_*.go` - UI components (health bar, progress bar, fonts)
- `renderer.go`, `sound.go` - Rendering and audio systems
- `assets/` - Sprites, fonts, sound effects, music
