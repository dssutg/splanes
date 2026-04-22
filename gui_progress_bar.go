package main

import "github.com/veandco/go-sdl2/sdl"

var (
	barColorGood   = sdl.Color{R: 0x00, G: 0xFF, B: 0x00, A: 0xFF}
	barColorNormal = sdl.Color{R: 0xFF, G: 0xFF, B: 0x00, A: 0xFF}
	barColorBad    = sdl.Color{R: 0xFF, G: 0x00, B: 0x00, A: 0xFF}
	barBorderColor = sdl.Color{R: 0x00, G: 0x00, B: 0x00, A: 0x00}
)

func RenderProgressBar(rect sdl.Rect, borderSize, progress int) {
	// Render progress bar outline/borders.
	RenderStrokedRect(rect, barBorderColor, borderSize)

	// Determine bar color based on progress.
	var barColor sdl.Color
	switch {
	case progress > 75:
		barColor = barColorGood
	case progress > 30:
		barColor = barColorNormal
	default:
		barColor = barColorBad
	}

	// Determine the full bar rectangle inside the borders.
	inner := sdl.Rect{
		X: rect.X + int32(borderSize),   // plus left border
		Y: rect.Y + int32(borderSize),   // plus top border
		W: rect.W - int32(borderSize*2), // minus left and right borders
		H: rect.H - int32(borderSize*2), // minus top and bottom borders
	}

	// Determine bar width based on progress.
	barW := inner.W * int32(progress) / 100

	// Render the bar rectangle.
	barRect := sdl.Rect{X: inner.X, Y: inner.Y, W: barW, H: inner.H}
	RenderFilledRect(barRect, barColor)
}
