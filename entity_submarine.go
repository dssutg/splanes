package main

import (
	"math/rand/v2"

	"github.com/veandco/go-sdl2/sdl"
)

// Submarine state constants.
const (
	SubmarineStateSurfacing = iota // Surfacing (animated up)
	SubmarineStateIdle             // On the surface
	SubmarineStateDiving           // Diving (animated down)
)

// Submarine timing constants.
const (
	maxSubmarineSurfacingTicks = 60 * 3
	maxSubmarineIdleTicks      = maxSubmarineSurfacingTicks + 40
	submarineTicksBeforeDiving = maxSubmarineSurfacingTicks / 3
)

// submarineFrames are the animation frames for submarines.
var submarineFrames = []sdl.Rect{
	{X: 529, Y: 100, W: 32, H: 98},
	{X: 496, Y: 100, W: 32, H: 98},
	{X: 463, Y: 100, W: 32, H: 98},
	{X: 430, Y: 100, W: 32, H: 98},
	{X: 397, Y: 100, W: 32, H: 98},
	{X: 364, Y: 100, W: 32, H: 98},
}

// NewSubmarine creates a submarine with random initial state.
// Submarines can rise, idle on the surface, or dive.
func NewSubmarine() *Entity {
	e := NewEntity(EntityTypeSubmarine)

	frame := &submarineFrames[0]

	e.Pos.W = frame.W * 1
	e.Pos.H = frame.H * 1
	e.Pos.X = rand.Int32N(WindowW)
	e.Pos.Y = -rand.Int32N(WindowH) - e.Pos.H

	e.VelY = 11

	e.Health = 100

	// Random initial state.
	switch rand.IntN(4) {
	case 0:
		e.State = SubmarineStateSurfacing
		e.Ticks = rand.Int32N(maxSubmarineSurfacingTicks)
	case 1:
		e.State = SubmarineStateIdle
		e.Ticks = RandInt32Range(maxSubmarineSurfacingTicks, maxSubmarineIdleTicks)
	default:
		e.State = SubmarineStateDiving
		e.Ticks = submarineTicksBeforeDiving
	}

	return e
}

// SubmarineTick handles the submarine state machine.
func SubmarineTick(e *Entity) {
	switch e.State {
	case SubmarineStateSurfacing:
		e.Ticks++
		if e.Ticks >= maxSubmarineSurfacingTicks {
			e.State = SubmarineStateIdle
		}
	case SubmarineStateIdle:
		e.Ticks++
		if e.Ticks >= maxSubmarineIdleTicks {
			e.Ticks = submarineTicksBeforeDiving
			e.State = SubmarineStateDiving
		}
	case SubmarineStateDiving:
		// Play the animation backwards.
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
		return
	}
}

// SubmarineRender draws the submarine with animation based on state.
func SubmarineRender(e *Entity) {
	frameNo := Clamp(int(e.Ticks)/10, 0, len(submarineFrames)-1)
	RenderSprite(e.Texture, e.Pos, submarineFrames[frameNo])
}
