package main

import (
	"fmt"
	"math"
	"os"
)

func clamp(x float64, min float64, max float64) float64 {
	return math.Min(math.Max(x, min), max)
}

func write_color(file *os.File, pixel_color color, samples_per_pixel int) {
	r := pixel_color.r()
	g := pixel_color.g()
	b := pixel_color.b()

	scale := 1.0 / float64(samples_per_pixel)
	r = scale * float64(r)
	g = scale * float64(g)
	b = scale * float64(b)

	r = clamp(r, 0.0, 0.999)
	g = clamp(g, 0.0, 0.999)
	b = clamp(b, 0.0, 0.999)

	fmt.Fprintf(file, "%d %d %d\n", int(255.99*r), int(255.99*g), int(255.99*b))
}
