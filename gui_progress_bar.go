package main

import "github.com/veandco/go-sdl2/sdl"

// Progress bar colors (green for high, yellow for medium, red for low).
var (
	barColorGood   = sdl.Color{R: 0x00, G: 0xFF, B: 0x00, A: 0xFF}
	barColorNormal = sdl.Color{R: 0xFF, G: 0xFF, B: 0x00, A: 0xFF}
	barColorBad    = sdl.Color{R: 0xFF, G: 0x00, B: 0x00, A: 0xFF}
	barBorderColor = sdl.Color{R: 0x00, G: 0x00, B: 0x00, A: 0x00}
)

// RenderProgressBar renders a progress bar in the given rectangle.
// progress should be in the range [0, 100]. If outside, it is clampped.
func RenderProgressBar(rect sdl.Rect, borderSize, progress int) {
	// Ensure the progress is within the valid range.
	progress = Clamp(progress, 0, 100)

	// Render progress bar outline/border.
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

	// Determine the inner rectangle (inside borders).
	inner := sdl.Rect{
		X: rect.X + int32(borderSize),   // plus left border
		Y: rect.Y + int32(borderSize),   // plus top border
		W: rect.W - int32(borderSize*2), // minus left and right borders
		H: rect.H - int32(borderSize*2), // minus top and bottom borders
	}

	// Calculate bar width based on progress.
	barW := inner.W * int32(progress) / 100

	// Render the filled portion of the bar.
	barRect := sdl.Rect{X: inner.X, Y: inner.Y, W: barW, H: inner.H}
	RenderFilledRect(barRect, barColor)
}
