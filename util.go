package main

import "math"

type Float interface {
	~float32 | ~float64
}

type Integer interface {
	~int8 | ~int16 | ~int32 | ~int64 | ~int
}

type Number interface {
	Float | Integer
}

const (
	radToDeg = math.Pi / 180
	degToRad = 180 / math.Pi
)

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

func DegToRad[T Float](degrees T) T {
	return degrees * radToDeg
}

func RadToDeg[T Float](radians T) T {
	return radians * degToRad
}

func RotateAround[T Float](x, y T, originX, originY T, radians T) (rx, ry T) {
	sin, cos := math.Sincos(float64(radians))
	return RotateAroundSinCos(x, y, originX, originY, T(sin), T(cos))
}

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

	return
}
