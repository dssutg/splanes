package main

import (
	"math/rand/v2"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	SubmarineStateRising = iota
	SubmarineStateIdle
	SubmarineStateDiving
)

const (
	maxSubmarineRisingTicks    = 60 * 3
	maxSubmarineIdleTicks      = maxSubmarineRisingTicks + 40
	submarineTicksBeforeDiving = maxSubmarineRisingTicks / 3
)

var submarineFrames = []sdl.Rect{
	{X: 529, Y: 100, W: 32, H: 98},
	{X: 496, Y: 100, W: 32, H: 98},
	{X: 463, Y: 100, W: 32, H: 98},
	{X: 430, Y: 100, W: 32, H: 98},
	{X: 397, Y: 100, W: 32, H: 98},
	{X: 364, Y: 100, W: 32, H: 98},
}

func NewSubmarine() *Entity {
	e := NewEntity(EntityTypeSubmarine)

	frame := &submarineFrames[0]

	e.Pos.W = frame.W * 1
	e.Pos.H = frame.H * 1
	e.Pos.X = rand.Int32N(WindowW)
	e.Pos.Y = -rand.Int32N(WindowH) - e.Pos.H

	e.VelY = 11

	e.Health = 100

	switch rand.IntN(4) {
	case 0:
		// Rising: 1/4 probability.
		e.State = SubmarineStateRising
		e.Ticks = rand.Int32N(maxSubmarineRisingTicks)
	case 1:
		// Idle: 1/4 probability.
		e.State = SubmarineStateIdle
		e.Ticks = RandInt32Range(maxSubmarineRisingTicks, maxSubmarineIdleTicks)
	default:
		// Diving: 2/4 probability.
		e.State = SubmarineStateDiving
		e.Ticks = submarineTicksBeforeDiving
	}

	return e
}

func SubmarineTick(e *Entity) {
	switch e.State {
	case SubmarineStateRising:
		e.Ticks++
		if e.Ticks >= maxSubmarineRisingTicks {
			e.State = SubmarineStateIdle
		}
	case SubmarineStateIdle:
		e.Ticks++
		if e.Ticks >= maxSubmarineIdleTicks {
			e.Ticks = submarineTicksBeforeDiving
			e.State = SubmarineStateDiving
		}
	case SubmarineStateDiving:
		e.Ticks--
		if e.Ticks <= 0 {
			// Submarine dived - invisible and unreachable. Remove it.
			e.Remove()
			return
		}
	}

	if e.Health <= 0 {
		NewExplosion(e.Pos.X, e.Pos.Y)
		e.Remove()
		return
	}

	e.Pos.X += e.VelX
	e.Pos.Y += e.VelY

	if e.Pos.Y >= WindowH {
		e.Remove()
	}
}

func SubmarineRender(e *Entity) {
	frameNo := Clamp(int(e.Ticks)/10, 0, len(submarineFrames)-1)
	RenderSprite(e.Texture, e.Pos, submarineFrames[frameNo])
}
