package main

import "github.com/veandco/go-sdl2/sdl"

var (
	healthBarBorderCrop = sdl.Rect{X: 364, Y: 199, W: 130, H: 14}
	healthBarElemCrop   = sdl.Rect{X: 494, Y: 199, W: 7, H: 14}
)

const maxHealthBarElements = 18

func RenderHealthBar(x, y, scale int32, health int) {
	// Render health bar borders.
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

	// Determine the number of bar elements to render.
	// The bar scales the health percent to a fixed number of elements.
	//
	// 1. Convert the health from percent to normalized value [0..1],
	// 2. Interpolate it for max elements.
	// 3. If bar can't represent more elements (extrapolation), only show max elements.
	elementCount := min(maxHealthBarElements, health*maxHealthBarElements/100)

	// Render the required number of bar elements.
	for i := range elementCount {
		elementOffset := int32(i) * healthBarElemCrop.W * scale

		dest := sdl.Rect{
			X: x + elementOffset,
			Y: y,
			W: healthBarElemCrop.W * scale,
			H: healthBarElemCrop.H * scale,
		}

		RenderSprite(TextureMain, dest, healthBarElemCrop)
	}
}
