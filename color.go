package main

import (
	"image"
	"image/color"
	"math"
	"os"
)

type Color struct {
	R, G, B float64
}

func clamp(x float64, min float64, max float64) float64 {
	return math.Min(math.Max(x, min), max)
}

func write_color(file *os.File, pixel_color Color, samples_per_pixel int) {
	r := pixel_color.R
	g := pixel_color.G
	b := pixel_color.B

	scale := 1.0 / float64(samples_per_pixel)
	r = scale * float64(r)
	g = scale * float64(g)
	b = scale * float64(b)

	r = clamp(r, 0.0, 0.999)
	g = clamp(g, 0.0, 0.999)
	b = clamp(b, 0.0, 0.999)

	image.Set(i, j, color.RGBA{uint8(255.99 * r), uint8(255.99 * g), uint8(255.99 * b), 255})
}
