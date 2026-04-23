package main

import "github.com/veandco/go-sdl2/sdl"

// Health bar sprite crop rectangles from the sprite sheet.
var (
	healthBarBorderCrop  = sdl.Rect{X: 364, Y: 199, W: 130, H: 14}
	healthBarSegmentCrop = sdl.Rect{X: 494, Y: 199, W: 7, H: 14}
)

const maxHealthBarSegments = 18 // Number of segments in the health bar

// RenderHealthBar renders the player health bar at (x, y) with the given scale.
// health should be in the range [0, 100].
func RenderHealthBar(x, y, scale int32, health int) {
	// Render health bar background/border.
	RenderSprite(
		TextureMain,
		sdl.Rect{
			X: x,
			Y: y, W: healthBarBorderCrop.W * scale,
			H: healthBarBorderCrop.H * scale,
		},
		healthBarBorderCrop,
	)

	// Fast path for empty health.
	if health <= 0 {
		return
	}

	// Calculate the number of bar segments to render.
	// The bar scales the health percent to a fixed number of segments.
	segmentCount := min(maxHealthBarSegments, health*maxHealthBarSegments/100)

	// Render the filled portion of the health bar.
	for i := range segmentCount {
		segmentOffset := int32(i) * healthBarSegmentCrop.W * scale

		dest := sdl.Rect{
			X: x + segmentOffset,
			Y: y,
			W: healthBarSegmentCrop.W * scale,
			H: healthBarSegmentCrop.H * scale,
		}

		RenderSprite(TextureMain, dest, healthBarSegmentCrop)
	}
}
