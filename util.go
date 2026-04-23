package main

import (
	"math"
	"math/rand/v2"
)

// Float is a type constraint for floating-point types.
type Float interface {
	~float32 | ~float64
}

// Integer is a type constraint for signed and unsigned integer types.
type Integer interface {
	~int8 | ~int16 | ~int32 | ~int64 | ~int |
		~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uint |
		uintptr
}

// Number is a type constraint for both float and integer types.
type Number interface {
	Float | Integer
}

// Angle conversion constants.
const (
	radToDeg = math.Pi / 180 // Radians to degrees
	degToRad = 180 / math.Pi // Degrees to radians
)

// Clamp limits a value to be within [minValue, maxValue].
func Clamp[T Number](value, minValue, maxValue T) T {
	if minValue > maxValue {
		minValue, maxValue = maxValue, minValue
	}
	switch {
	case value < minValue:
		return minValue
	case value > maxValue:
		return maxValue
	default:
		return value
	}
}

// DegToRad converts degrees to radians.
func DegToRad[T Float](degrees T) T {
	return degrees * radToDeg
}

// RadToDeg converts radians to degrees.
func RadToDeg[T Float](radians T) T {
	return radians * degToRad
}

// RotateAround rotates a point (x, y) around (originX, originY) by radians.
func RotateAround[T Float](x, y T, originX, originY T, radians T) (rx, ry T) {
	sin, cos := math.Sincos(float64(radians))
	return RotateAroundSinCos(x, y, originX, originY, T(sin), T(cos))
}

// RotateAroundSinCos rotates a point around an origin using precomputed sin/cos.
func RotateAroundSinCos[T Float](x, y T, originX, originY T, sin, cos T) (rx, ry T) {
	// Make the point relative to the origin.
	x -= originX
	y -= originY

	// Rotate the point.
	rx = x*cos - y*sin
	ry = x*sin + y*cos

	// Move the point back to the original coordinate system.
	rx += originX
	ry += originY

	return rx, ry
}

// RandIntRange returns a random integer in [minValue, maxValue].
func RandIntRange(minValue, maxValue int) int {
	if minValue > maxValue {
		minValue, maxValue = maxValue, minValue
	}
	return minValue + rand.IntN(maxValue-minValue+1)
}

// RandInt32Range returns a random int32 in [minValue, maxValue].
func RandInt32Range(minValue, maxValue int32) int32 {
	if minValue > maxValue {
		minValue, maxValue = maxValue, minValue
	}
	return minValue + rand.Int32N(maxValue-minValue+1)
}
