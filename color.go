package main

import "math"

func clamp(x, min, max float64) float64 {
	return math.Min(math.Max(x, min), max)
}
