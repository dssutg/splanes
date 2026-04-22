package main

func Clamp[T ~int | ~int32](value, minValue, maxValue T) T {
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
