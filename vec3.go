package main

import (
	"math"
	"math/rand"
)

// Vec3 is a 3D vector
type Vec3 struct {
	X float32
	Y float32
	Z float32
}

// Vector Functions
func vec_add(a, b Vec3) Vec3 {
	return Vec3{a.X + b.X, a.Y + b.Y, a.Z + b.Z}
}

func vec_sub(a, b Vec3) Vec3 {
	return Vec3{a.X - b.X, a.Y - b.Y, a.Z - b.Z}
}

func vec_mul(a, b Vec3) Vec3 {
	return Vec3{a.X * b.X, a.Y * b.Y, a.Z * b.Z}
}

func vec_div(a, b Vec3) Vec3 {
	return Vec3{a.X / b.X, a.Y / b.Y, a.Z / b.Z}
}

func vec_len(a Vec3) float32 {
	return float32(math.Sqrt(float64(a.X*a.X + a.Y*a.Y + a.Z*a.Z)))
}

func vec_len_squared(a Vec3) float32 {
	return float32(a.X*a.X + a.Y*a.Y + a.Z*a.Z)
}

func vec_add_scalar(a Vec3, b float32) Vec3 {
	return Vec3{a.X + b, a.Y + b, a.Z + b}
}

func vec_sub_scalar(a Vec3, b float32) Vec3 {
	return Vec3{a.X - b, a.Y - b, a.Z - b}
}

func vec_mul_scalar(a Vec3, b float32) Vec3 {
	return Vec3{a.X * b, a.Y * b, a.Z * b}
}

func vec_div_scalar(a Vec3, b float32) Vec3 {
	return Vec3{a.X / b, a.Y / b, a.Z / b}
}

func vec_dot(a, b Vec3) float32 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

func vec_cross(a, b Vec3) Vec3 {
	return Vec3{a.Y*b.Z - a.Z*b.Y, a.Z*b.X - a.X*b.Z, a.X*b.Y - a.Y*b.X}
}

func unit_vector(a Vec3) Vec3 {
	return vec_div_scalar(a, vec_len(a))
}

func random_in_unit_sphere() Vec3 {
	for {
		p := Vec3{rand.Float32(), rand.Float32(), rand.Float32()}
		if vec_len_squared(p) < 1.0 {
			return p
		}
	}
}
