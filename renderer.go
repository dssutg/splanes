package main

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	WindowDelayMilliseconds = 50

	WindowScale = 1

	WindowWidth  = 800 * WindowScale
	WindowHeight = 600 * WindowScale

	TileSize = 32
)

var WindowRect = sdl.Rect{X: 0, Y: 0, W: WindowWidth, H: WindowHeight}

var (
	renderer *sdl.Renderer
	textures [16]*sdl.Texture
	window   *sdl.Window
)

const windowTitle = "Splanes"

func renderSprite(texture int, dest, src sdl.Rect) {
	renderer.Copy(textures[texture], &src, &dest)
}

func loadTexture(filename string) *sdl.Texture {
	bitmap, err := img.Load(filename)
	if err != nil {
		fatalf("can't load %v: %v", filename, err)
	}
	defer bitmap.Free()

	texture, err := renderer.CreateTextureFromSurface(bitmap)
	if err != nil {
		fatalf("can't create texture from %v: %v", filename, err)
	}

	return texture
}

func loadTextures() {
	textures[0] = loadTexture("assets/sprites/sprites.png")
}

func renderRect(rect sdl.Rect, color sdl.Color) {
	renderer.SetDrawColor(color.R, color.G, color.B, color.A)
	renderer.FillRect(&rect)
}
