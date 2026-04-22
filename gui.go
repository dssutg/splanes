package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type FontCacheEntry struct {
	font *ttf.Font
	size int
}

var fontCache []FontCacheEntry

func loadFont(size int) *ttf.Font {
	// Find the font with the required size in the loaded font cache
	for _, entry := range fontCache {
		if entry.size == size {
			return entry.font
		}
	}

	font, err := ttf.OpenFont("assets/fonts/OpenSans-Bold.ttf", size)
	if err != nil {
		fatalf("can't open font file: %v", sdl.GetError())
	}

	// Add the new font to the font cache
	fontCache = append(fontCache, FontCacheEntry{font: font, size: size})

	return font
}

func renderProgressBar(rect sdl.Rect, strokeSize, progress int) {
	borderColor := sdl.Color{R: 0, G: 0, B: 0, A: 0}

	renderRect(sdl.Rect{X: rect.X, Y: rect.Y, W: rect.W, H: int32(strokeSize)}, borderColor)
	renderRect(sdl.Rect{X: rect.X, Y: rect.Y + rect.H - int32(strokeSize) - 1, W: rect.W, H: int32(strokeSize)}, borderColor)
	renderRect(sdl.Rect{X: rect.X, Y: rect.Y, W: int32(strokeSize), H: rect.H}, borderColor)
	renderRect(sdl.Rect{X: rect.X + rect.W - int32(strokeSize) - 1, Y: rect.Y, W: int32(strokeSize), H: rect.H}, borderColor)

	barColor := sdl.Color{R: 255, G: 0, B: 0, A: 255}
	switch {
	case progress > 75:
		barColor = sdl.Color{R: 0, G: 255, B: 0, A: 255}
	case progress > 30:
		barColor = sdl.Color{R: 255, G: 255, B: 0, A: 255}
	}

	barRect := sdl.Rect{
		X: rect.X + int32(strokeSize),
		Y: rect.Y + int32(strokeSize),
		W: (rect.W - int32(strokeSize*2)) * int32(progress) / 100,
		H: rect.H - int32(strokeSize*2),
	}

	renderRect(barRect, barColor)
}

func renderString(
	x, y, size int,
	color sdl.Color,
	relativeToWindowCenter bool,
	lineNo int,
	format string,
	args ...any,
) {
	text := fmt.Sprintf(format, args...)

	if len(text) == 0 {
		return
	}

	font := loadFont(size)

	fontSurface, err := font.RenderUTF8Blended(text, color)
	if err != nil {
		fatalf("can't get font surface: %v", err)
	}
	defer fontSurface.Free()

	texture, err := renderer.CreateTextureFromSurface(fontSurface)
	if err != nil {
		fatalf("can't create texture: %v", err)
	}
	defer texture.Destroy()

	src := sdl.Rect{X: 0, Y: 0, W: fontSurface.W, H: fontSurface.H}

	dest := sdl.Rect{X: int32(x), Y: int32(y), W: fontSurface.W, H: fontSurface.H}
	if relativeToWindowCenter {
		dest.X = (WindowWidth - fontSurface.W) / 2
		dest.Y = (WindowHeight + (fontSurface.H+50)*int32(lineNo)) / 2
	}

	renderer.Copy(texture, &src, &dest)
}

func renderHealthBar(health int) {
	const scale = 2

	const (
		x = 20
		y = 20
	)

	crop := sdl.Rect{X: 364, Y: 199, W: 130, H: 14}

	renderSprite(0, sdl.Rect{X: x, Y: y, W: crop.W * scale, H: crop.H * scale}, crop)

	crop.X = 494
	crop.W = 7

	if health < 0 {
		return
	}

	const maxElems = 18
	elemNum := min(maxElems, health*maxElems/100)

	for i := range elemNum {
		renderSprite(0, sdl.Rect{X: x + int32(i)*crop.W*scale, Y: y, W: crop.W * scale, H: crop.H * scale}, crop)
	}
}

func renderSmallLogo() {
	const scale = 1

	const (
		cropW = 115
		cropH = 58
	)

	w := cropW * scale
	h := cropH * scale

	renderSprite(0, sdl.Rect{X: int32(WindowWidth - w), Y: 0, W: int32(w), H: int32(h)}, sdl.Rect{X: 700, Y: 401, W: cropW, H: cropH})
}
