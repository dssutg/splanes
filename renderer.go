package main

import (
	"log"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

// Window configuration constants.
const (
	WindowDelayMilliseconds = 50 // Delay between window updates (ms)

	WindowScale = 1 // Window scaling factor

	WindowW = 800 * WindowScale // Window width in pixels
	WindowH = 600 * WindowScale // Window height in pixels

	TileSize = 32 // Size of background tiles
)

// WindowRect is the window bounds for rendering calculations.
var WindowRect = sdl.Rect{X: 0, Y: 0, W: WindowW, H: WindowH}

// Renderer and texture globals.
var (
	renderer *sdl.Renderer
	textures [16]*sdl.Texture // Loaded texture atlases
	window   *sdl.Window
)

const windowTitle = "Splanes"

// TextureID identifies loaded texture atlases.
type TextureID int

// Texture atlas IDs.
const (
	TextureMain TextureID = iota // Main sprite sheet
)

// RenderSprite renders a sprite from a texture atlas to the screen.
// texture: texture atlas ID
// dest: destination rectangle (position and size on screen)
// src: source rectangle (position and size in texture atlas)
func RenderSprite(texture TextureID, dest, src sdl.Rect) {
	_ = renderer.Copy(textures[texture], &src, &dest)
}

// RenderSpriteEx renders a sprite with rotation and flipping.
// Same as RenderSprite but adds rotation angle, center point, and flip mode.
func RenderSpriteEx(
	texture TextureID,
	dest, src sdl.Rect,
	angle float64,
	center *sdl.Point,
	flip sdl.RendererFlip,
) {
	_ = renderer.CopyEx(textures[texture], &src, &dest, angle, center, flip)
}

// LoadTexture loads an image and creates a GPU texture from it.
// Returns the created texture, or logs fatal error on failure.
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

// LoadTextures loads all game texture atlases into GPU memory.
func LoadTextures() {
	textures[TextureMain] = LoadTexture("assets/sprites/sprites.png")
}

// RenderFilledRect renders a filled rectangle with the given color.
func RenderFilledRect(rect sdl.Rect, color sdl.Color) {
	_ = renderer.SetDrawColor(color.R, color.G, color.B, color.A)
	_ = renderer.FillRect(&rect)
}

// RenderStrokedRect renders the outline of a rectangle with the given color and line width.
func RenderStrokedRect(rect sdl.Rect, color sdl.Color, lineWidth int) {
	_ = renderer.SetDrawColor(color.R, color.G, color.B, color.A)

	w := int32(lineWidth)

	// Render top border.
	_ = renderer.FillRect(&sdl.Rect{
		X: rect.X,
		Y: rect.Y,
		W: rect.W,
		H: w,
	})

	// Render bottom border.
	_ = renderer.FillRect(&sdl.Rect{
		X: rect.X,
		Y: rect.Y + rect.H - w - 1,
		W: rect.W,
		H: w,
	})

	// Render left border.
	_ = renderer.FillRect(&sdl.Rect{
		X: rect.X,
		Y: rect.Y,
		W: w,
		H: rect.H,
	})

	// Render right border.
	_ = renderer.FillRect(&sdl.Rect{
		X: rect.X + rect.W - w - 1,
		Y: rect.Y,
		W: w,
		H: rect.H,
	})
}
