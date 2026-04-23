package main

import (
	"log"
	"math/rand/v2"

	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// running controls the main game loop. Set to false when the user closes the window.
var running = true

// waterLayer1 and waterLayer2 are the two scrolling water background layers.
// They scroll independently to create a seamless tiling effect.
var (
	waterLayer1 int32
	waterLayer2 int32
)

// reset initializes or resets the game state for a new game.
// It creates a fresh player, resets water layer positions,
// returns to the main menu, and starts background music.
func reset() {
	player = NewPlayer()

	waterLayer1 = -WindowH
	waterLayer2 = 0

	menuID = MenuTypeMain

	mainMenu.SelectedIndex = 0
	exitMenu.SelectedIndex = 0
	aboutMenu.SelectedIndex = 0
	loseMenu.SelectedIndex = 0

	PlayMusic(musicBackground0, 70)
}

// restart removes all entities and resets the game state,
// used when restarting after game over or player death.
func restart() {
	RemoveAllEntities()
	reset()
}

// renderWaterLayer renders a tiled water background at the given Y offset.
// It tiles the 32x32 water sprite to fill the entire window.
func renderWaterLayer(offsetY int32) {
	tileW := (WindowW + TileSize - 1) / TileSize
	tileH := (WindowH + TileSize - 1) / TileSize

	for tileY := int32(0); tileY < int32(tileH); tileY++ {
		for tileX := int32(0); tileX < int32(tileW); tileX++ {
			dest := sdl.Rect{
				X: tileX * TileSize,
				Y: tileY*TileSize + offsetY,
				W: TileSize,
				H: TileSize,
			}

			src := sdl.Rect{X: 265, Y: 364, W: 32, H: 32}

			RenderSprite(TextureMain, dest, src)
		}
	}
}

// doGameLoop is the main game loop.
// It uses a high-resolution timer for precise timing control,
// running game ticks at fixed intervals independent of frame rate.
// The loop processes events, updates game logic, and renders at 60 FPS.
func doGameLoop() {
	// High-resolution timer for precise timing.
	perFreq := float64(sdl.GetPerformanceFrequency())
	lastTime := float64(sdl.GetPerformanceCounter())

	// Accumulated "extra" time since last tick.
	// When this reaches timePerTick, we run a game tick.
	var unprocessed float64

	// How many counter ticks per one game tick (1/N sec).
	// At 60 FPS this would be ~1/3 of a frame.
	const ticksPerSec = 20
	timePerTick := perFreq / ticksPerSec

	frames := 0
	ticks := 0

	// For measuring actual FPS once per second.
	lastTimer := sdl.GetTicks64()

	for running {
		// Poll all pending SDL events: window close, keyboard, etc.
		// Called every frame for responsive input handling.
		pollEvents()

		// Calculate how much real time passed since last frame.
		now := float64(sdl.GetPerformanceCounter())
		unprocessed += (now - lastTime) / timePerTick
		lastTime = now

		// Flag to know if we should render this frame.
		// Only render when game advanced at least one tick.
		shouldRender := false

		// This ensures consistent ticks/sec regardless of
		// actual framerate. Time accumulates if we're slow.
		for unprocessed >= 1 {
			tick()
			ticks++
			unprocessed--
			shouldRender = true
		}

		// Only render if game advanced.
		if shouldRender {
			frames++
			render()
		}

		// Report actual ticks and FPS once per second.
		const milliPerSec = 1000
		if sdl.GetTicks64()-lastTimer > milliPerSec {
			lastTimer += milliPerSec
			ticks = 0
			frames = 0
		}

		// Small sleep to prevent busy-spinning.
		sdl.Delay(2)
	}
}

// pollEvents processes all pending SDL events: window close, keyboard, etc.
// Called every frame for responsive input handling.
func pollEvents() {
	for {
		event := sdl.PollEvent()
		if event == nil {
			break
		}
		switch e := event.(type) {
		case *sdl.QuitEvent:
			running = false
		case *sdl.KeyboardEvent:
			switch e.Type {
			case sdl.KEYDOWN:
				UpdateKey(e.Keysym.Sym, true)
			case sdl.KEYUP:
				UpdateKey(e.Keysym.Sym, false)
			}
		}
	}
}

// tick runs once per game update (20 times per second).
// It handles pause menu activation and delegates to either
// menu or game logic depending on the current menu state.
func tick() {
	if Keys[KeyPause] {
		Keys[KeyPause] = false
		switch menuID {
		case MenuTypeNone:
			menuID = MenuTypeMain
		case MenuTypeMain:
			menuID = MenuTypeNone
		}
	}

	pause := false

	if menuID != MenuTypeNone {
		tickMenu()
		pause = true
	}

	if !pause {
		tickGame()
	}
}

// tickGame runs the game logic: spawning enemies, updating entities,
// and scrolling the water background. Called every game tick (20 times/second).
func tickGame() {
	// Spawn entities.
	if rand.IntN(20) == 0 {
		NewEnemyPlane()
	}
	if rand.IntN(80) == 0 {
		if rand.IntN(10) <= 3 {
			NewSubmarine()
		} else {
			NewShip()
		}
	}
	if rand.IntN(30) == 0 {
		NewIsland()
	}
	if rand.IntN(100) == 0 {
		NewHealer()
	}

	// Move water layers.
	waterLayer1 += 10
	waterLayer2 += 10
	if waterLayer2 >= WindowH {
		waterLayer2 = waterLayer1 - WindowH
		// Swap the water layers to create the infinite scroll illusion.
		waterLayer1, waterLayer2 = waterLayer2, waterLayer1
	}

	// Update all entities.
	for i := range EntityPool {
		e := &EntityPool[i]

		if e.Kind == EntityTypeNone {
			continue
		}

		entry := entityTable[e.Kind]
		entry.Tick(e)
	}
}

// render clears the screen and draws all visible game elements:
// water layers, entities (sorted by z-index), UI elements (health bar,
// bomb cooldown, score display), and the current menu if active.
func render() {
	// Clear the screen.
	_ = renderer.SetDrawColor(0, 0, 0, 0)
	_ = renderer.Clear()

	// Render water layers.
	renderWaterLayer(waterLayer1)
	renderWaterLayer(waterLayer2)

	// Render all entities.
	for zIndex := 0; zIndex <= 2; zIndex++ {
		for i := range EntityPool {
			e := &EntityPool[i]
			entry := entityTable[e.Kind]
			if entry.ZIndex == int32(zIndex) {
				entry.Render(e)
			}
		}
	}

	// Render player health bar.
	RenderHealthBar(20, 20, 2, int(player.Health))

	// Render bomb cooldown.
	RenderProgressBar(
		sdl.Rect{X: 580, Y: 20, W: 100, H: 25},
		5,
		int(player.BombTicks*100/PlayerMaxBombTickTime),
	)

	RenderSmallLogo()

	// Render stats.
	RenderStringf(
		RenderStringOptions{
			X:     300,
			Y:     20,
			Size:  20,
			Color: sdl.Color{R: 255, G: 202, B: 65, A: 255},
		},
		"SCORE: %d, DISTANCE: %d",
		player.Score,
		player.Distance,
	)

	renderMenu()

	renderer.Present()
}

func main() {
	// Initialize SDL.
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		log.Fatal("can't init SDL:", err)
	}
	defer sdl.Quit()

	if err := ttf.Init(); err != nil {
		log.Fatal("can't init TTF:", err)
	}

	if err := sdl.InitSubSystem(sdl.INIT_AUDIO); err != nil {
		log.Fatal("can't init audio:", err)
	}

	if err := mix.OpenAudio(44100, mix.DEFAULT_FORMAT, 2, 4096); err != nil {
		log.Fatal("can't open audio:", err)
	}
	defer mix.CloseAudio()

	var err error
	window, err = sdl.CreateWindow(
		windowTitle,
		sdl.WINDOWPOS_CENTERED,
		sdl.WINDOWPOS_CENTERED,
		WindowW,
		WindowH,
		0,
	)
	if err != nil {
		log.Fatal("can't create window:", err)
	}
	defer window.Destroy()

	renderer, err = sdl.CreateRenderer(
		window,
		-1,
		sdl.RENDERER_ACCELERATED|sdl.RENDERER_TARGETTEXTURE,
	)
	if err != nil {
		log.Fatal("can't create renderer:", err)
	}

	LoadTextures()
	InitSoundManager()

	reset()

	doGameLoop()
}
