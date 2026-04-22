package main

import (
	"fmt"
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

type RenderStringOptions struct {
	X                      int
	Y                      int
	Size                   int
	Color                  sdl.Color
	RelativeToWindowCenter bool
	LineNo                 int
}

func RenderString(opts RenderStringOptions, format string, args ...any) {
	text := fmt.Sprintf(format, args...)

	if len(text) == 0 {
		return
	}

	font := LoadFont(opts.Size)

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
