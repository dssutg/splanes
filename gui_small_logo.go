package main

import "github.com/veandco/go-sdl2/sdl"

var smallLogoCrop = sdl.Rect{X: 700, Y: 401, W: 115, H: 58}

func RenderSmallLogo() {
	dest := sdl.Rect{
		X: int32(WindowW - smallLogoCrop.W),
		Y: 0,
		W: int32(smallLogoCrop.W),
		H: int32(smallLogoCrop.H),
	}

	RenderSprite(TextureMain, dest, smallLogoCrop)
}
