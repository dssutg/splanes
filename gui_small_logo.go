package main

import "github.com/veandco/go-sdl2/sdl"

// smallLogoCrop is the sprite crop for the game logo.
var smallLogoCrop = sdl.Rect{X: 700, Y: 401, W: 115, H: 58}

// RenderSmallLogo renders the game logo in the top-right corner.
func RenderSmallLogo() {
	dest := sdl.Rect{
		X: WindowW - smallLogoCrop.W,
		Y: 0,
		W: smallLogoCrop.W,
		H: smallLogoCrop.H,
	}

	RenderSprite(TextureMain, dest, smallLogoCrop)
}
