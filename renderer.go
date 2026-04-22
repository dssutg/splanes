package main

import (
	"log"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	WindowDelayMilliseconds = 50

	WindowScale = 1

	WindowW  = 800 * WindowScale
	WindowH = 600 * WindowScale

	TileSize = 32
)

var WindowRect = sdl.Rect{X: 0, Y: 0, W: WindowW, H: WindowH}

var (
	renderer *sdl.Renderer
	textures [16]*sdl.Texture
	window   *sdl.Window
)

const windowTitle = "Splanes"

type TextureID int

const (
	TextureMain TextureID = iota
)

func RenderSprite(texture TextureID, dest, src sdl.Rect) {
	renderer.Copy(textures[texture], &src, &dest)
}

func LoadTexture(filename string) *sdl.Texture {
	bitmap, err := img.Load(filename)
	if err != nil {
		log.Fatalf("can't load %v: %v", filename, err)
	}
	defer bitmap.Free()

	texture, err := renderer.CreateTextureFromSurface(bitmap)
	if err != nil {
		log.Fatalf("can't create texture from %v: %v", filename, err)
	}

	return texture
}

func LoadTextures() {
	textures[TextureMain] = LoadTexture("assets/sprites/sprites.png")
}

func RenderFilledRect(rect sdl.Rect, color sdl.Color) {
	renderer.SetDrawColor(color.R, color.G, color.B, color.A)
	renderer.FillRect(&rect)
}

func RenderStrokedRect(rect sdl.Rect, color sdl.Color, lineWidth int) {
	renderer.SetDrawColor(color.R, color.G, color.B, color.A)

	w := int32(lineWidth)

	// Render top border.
	renderer.FillRect(&sdl.Rect{X: rect.X, Y: rect.Y, W: rect.W, H: w})

	// Render bottom border.
	renderer.FillRect(&sdl.Rect{X: rect.X, Y: rect.Y + rect.H - w - 1, W: rect.W, H: w})

	// Render left border.
	renderer.FillRect(&sdl.Rect{X: rect.X, Y: rect.Y, W: w, H: rect.H})

	// Render right border.
	renderer.FillRect(&sdl.Rect{X: rect.X + rect.W - w - 1, Y: rect.Y, W: w, H: rect.H})
}
