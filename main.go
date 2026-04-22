package main

import (
	"log"
	"math/rand/v2"

	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

var running = true

var (
	waterLayer1 int32
	waterLayer2 int32
)

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

func restart() {
	RemoveAllEntities()
	reset()
}

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

func tickGame() {
	if rand.IntN(20) == 0 {
		NewEnemyPlane()
	}
	if rand.IntN(80) == 0 {
		NewShip()
	}
	if rand.IntN(30) == 0 {
		NewIsland()
	}
	if rand.IntN(100) == 0 {
		NewHealer()
	}

	waterLayer1 += 10
	waterLayer2 += 10
	if waterLayer2 >= WindowH {
		waterLayer2 = waterLayer1 - WindowH
		waterLayer1, waterLayer2 = waterLayer2, waterLayer1
	}

	for i := range EntityPool {
		e := &EntityPool[i]
		entry := entityTable[e.Kind]
		if entry.ZIndex >= 0 {
			entry.Tick(e)
		}
	}
}

func render() {
	renderer.SetDrawColor(0, 0, 0, 0)
	renderer.Clear()

	renderWaterLayer(waterLayer1)
	renderWaterLayer(waterLayer2)

	for zIndex := 0; zIndex <= 2; zIndex++ {
		for i := range EntityPool {
			e := &EntityPool[i]
			entry := entityTable[e.Kind]
			if entry.ZIndex == int32(zIndex) {
				entry.Render(e)
			}
		}
	}

	RenderHealthBar(20, 20, 2, int(player.Health))

	RenderProgressBar(
		sdl.Rect{X: 580, Y: 20, W: 100, H: 25},
		5,
		int(player.BombTicks*100/PlayerMaxBombTickTime),
	)

	RenderSmallLogo()

	RenderString(
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
