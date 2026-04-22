package main

import (
	"math/rand/v2"

	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

var running = true

var (
	layer1 int32
	layer2 int32
)

func reset() {
	player = newPlayer()

	layer1 = -WindowHeight
	layer2 = 0

	menuID = MenuTypeMain

	mainMenu.SelectedIndex = 0
	exitMenu.SelectedIndex = 0
	aboutMenu.SelectedIndex = 0
	loseMenu.SelectedIndex = 0

	playMusic(musicBackground0, 70)
}

func restart() {
	removeAllEntities()
	reset()
}

func renderLayer(offsetY int32) {
	tileWidth := (WindowWidth + TileSize - 1) / TileSize
	tileHeight := (WindowHeight + TileSize - 1) / TileSize

	for tileY := int32(0); tileY < int32(tileHeight); tileY++ {
		for tileX := int32(0); tileX < int32(tileWidth); tileX++ {
			renderSprite(
				0,
				sdl.Rect{X: tileX * TileSize, Y: tileY*TileSize + offsetY, W: TileSize, H: TileSize},
				sdl.Rect{X: 265, Y: 364, W: 32, H: 32},
			)
		}
	}
}

func doGameLoop() {
	for running {
		// Handle new events
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
					updateKey(e.Keysym.Sym, true)
				case sdl.KEYUP:
					updateKey(e.Keysym.Sym, false)
				}
			}
		}

		// Handle pause and menus
		if keys[KeyPause] {
			keys[KeyPause] = false
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

		// Clear screen
		renderer.SetDrawColor(0, 0, 0, 0)
		renderer.Clear()

		if !pause {
			// Spawn new entities
			if rand.IntN(20) == 0 {
				newEnemyPlane()
			}
			if rand.IntN(80) == 0 {
				newShip()
			}
			if rand.IntN(30) == 0 {
				newIsland()
			}
			if rand.IntN(100) == 0 {
				newHealer()
			}

			// Update water layers
			layer1 += 10
			layer2 += 10
			if layer2 >= WindowHeight {
				layer2 = layer1 - WindowHeight
				layer1, layer2 = layer2, layer1
			}
		}

		// Render water layers
		renderLayer(layer1)
		renderLayer(layer2)

		// Update and render entitites
		for zIndex := 0; zIndex <= 2; zIndex++ {
			for i := range entityPool {
				e := &entityPool[i]
				entry := entityTable[e.etype]
				if entry.ZIndex == int32(zIndex) {
					if !pause {
						entry.Tick(e)
					}
					entry.Render(e)
				}
			}
		}

		// Render GUI
		renderHealthBar(int(player.health))

		renderProgressBar(
			sdl.Rect{X: 580, Y: 20, W: 100, H: 25},
			5,
			int(player.bombTickTime*100/PlayerMaxBombTickTime),
		)

		renderSmallLogo()

		renderString(
			300,
			20,
			20,
			sdl.Color{R: 255, G: 202, B: 65, A: 255},
			false,
			0,
			"SCORE: %d, DISTANCE: %d",
			player.score,
			player.distance,
		)

		renderMenu()

		renderer.Present()
		sdl.Delay(WindowDelayMilliseconds)
	}
}

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fatalf("can't init SDL: %v", err)
	}
	defer sdl.Quit()

	if err := ttf.Init(); err != nil {
		fatalf("can't init TTF: %v", err)
	}

	if err := sdl.InitSubSystem(sdl.INIT_AUDIO); err != nil {
		fatalf("can't init audio: %v", err)
	}

	if err := mix.OpenAudio(44100, mix.DEFAULT_FORMAT, 2, 4096); err != nil {
		fatalf("can't open audio: %v", err)
	}
	defer mix.CloseAudio()

	var err error
	window, err = sdl.CreateWindow(
		windowTitle,
		sdl.WINDOWPOS_CENTERED,
		sdl.WINDOWPOS_CENTERED,
		WindowWidth,
		WindowHeight,
		0,
	)
	if err != nil {
		fatalf("can't create window: %v", err)
	}
	defer window.Destroy()

	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_TARGETTEXTURE)
	if err != nil {
		fatalf("can't create renderer: %v", err)
	}

	loadTextures()
	initSoundManager()

	reset()

	doGameLoop()
}
