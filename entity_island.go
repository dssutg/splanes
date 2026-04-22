package main

import (
	"math/rand/v2"

	"github.com/veandco/go-sdl2/sdl"
)

var islandFrames = []sdl.Rect{
	{X: 100, Y: 496, W: 64, H: 65},
	{X: 165, Y: 496, W: 64, H: 65},
	{X: 230, Y: 496, W: 64, H: 65},
}

func NewIsland() *Entity {
	e := NewEntity(EntityTypeIsland)

	e.InitialFrameNo = rand.Int32N(int32(len(islandFrames)))

	e.Pos.W = islandFrames[e.InitialFrameNo].W * 3
	e.Pos.H = islandFrames[e.InitialFrameNo].H * 3
	e.Pos.X = rand.Int32N(WindowW)
	e.Pos.Y = -rand.Int32N(WindowH) - e.Pos.H

	e.VelY = 1

	return e
}

func IslandTick(e *Entity) {
	e.Pos.X += e.VelX * 10
	e.Pos.Y += e.VelY * 10
	if e.Pos.Y >= WindowH {
		e.Remove()
	}
}

func IslandRender(e *Entity) {
	RenderSprite(e.Texture, e.Pos, islandFrames[e.InitialFrameNo])
}
