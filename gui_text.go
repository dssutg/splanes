package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/veandco/go-sdl2/sdl"
)

// textBuf is a reusable buffer for rendering text to reduce allocations.
var textBuf strings.Builder

// RenderStringOptions contains options for rendering text.
type RenderStringOptions struct {
	X                      int       // X position (ignored if RelativeToWindowCenter is set)
	Y                      int       // Y position
	Size                   int       // Font size in points
	Color                  sdl.Color // Text color
	RelativeToWindowCenter bool      // If true, position relative to center of window
	LineNo                 int       // Line offset for relative positioning
}

// RenderStringf renders formatted text to the screen.
// Uses a global buffer to avoid allocations.
func RenderStringf(opts RenderStringOptions, format string, args ...any) {
	// Optimistically use just the format string if no arguments.
	text := format

	// If arguments provided, fallback to slower buffer fill.
	if len(args) > 0 {
		textBuf.Reset()
		fmt.Fprintf(&textBuf, format, args...)
		text = textBuf.String()
	}

	// Fast path on empty text.
	if len(text) == 0 {
		return
	}

	// Load font from cache.
	font := LoadFont(opts.Size)

	// XXX Text texture creation needs optimization.
	fontSurface, err := font.RenderUTF8Blended(text, opts.Color)
	if err != nil {
		log.Fatal("can't get font surface:", err)
	}
	defer fontSurface.Free()

	texture, err := renderer.CreateTextureFromSurface(fontSurface)
	if err != nil {
		log.Fatal("can't create texture:", err)
	}
	defer texture.Destroy()

	src := sdl.Rect{
		X: 0,
		Y: 0,
		W: fontSurface.W,
		H: fontSurface.H,
	}

	dest := sdl.Rect{
		X: int32(opts.X),
		Y: int32(opts.Y),
		W: fontSurface.W,
		H: fontSurface.H,
	}

	if opts.RelativeToWindowCenter {
		dest.X = (WindowW - fontSurface.W) / 2
		dest.Y = (WindowH + (fontSurface.H+50)*int32(opts.LineNo)) / 2
	}

	renderer.Copy(texture, &src, &dest)
}
