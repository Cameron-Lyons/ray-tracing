package main

import "math"

type texture interface {
	value(u, v float64, p Vec3) Vec3
}

type solid_color struct {
	color_value Vec3
}

func (s solid_color) value(u, v float64, p Vec3) Vec3 {
	return s.color_value
}

type checker_texture struct {
	odd, even texture
}

func (c checker_texture) value(u, v float64, p Vec3) Vec3 {
	sines := math.Sin(10*p.X) * math.Sin(10*p.Y) * math.Sin(10*p.Z)
	if sines < 0 {
		return c.odd.value(u, v, p)
	}
	return c.even.value(u, v, p)
}

type noise_texture struct {
	noise perlin
	scale float64
}

func (n noise_texture) value(u, v float64, p Vec3) Vec3 {
	return vec_mul_scalar(Vec3{1, 1, 1}, 0.5*(1+math.Sin(n.scale*p.Z+10*n.noise.turbulence(p, 7))))
}
