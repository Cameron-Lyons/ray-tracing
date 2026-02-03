package main

import (
	"math"
	"math/rand"
)

type Vec3 struct {
	X, Y, Z float64
}

func vec_add(a, b Vec3) Vec3 {
	return Vec3{a.X + b.X, a.Y + b.Y, a.Z + b.Z}
}

func vec_sub(a, b Vec3) Vec3 {
	return Vec3{a.X - b.X, a.Y - b.Y, a.Z - b.Z}
}

func vec_mul(a, b Vec3) Vec3 {
	return Vec3{a.X * b.X, a.Y * b.Y, a.Z * b.Z}
}

func vec_len(a Vec3) float64 {
	return math.Sqrt(a.X*a.X + a.Y*a.Y + a.Z*a.Z)
}

func vec_len_squared(a Vec3) float64 {
	return a.X*a.X + a.Y*a.Y + a.Z*a.Z
}

func vec_mul_scalar(a Vec3, b float64) Vec3 {
	return Vec3{a.X * b, a.Y * b, a.Z * b}
}

func vec_div_scalar(a Vec3, b float64) Vec3 {
	return Vec3{a.X / b, a.Y / b, a.Z / b}
}

func vec_dot(a, b Vec3) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

func vec_cross(a, b Vec3) Vec3 {
	return Vec3{a.Y*b.Z - a.Z*b.Y, a.Z*b.X - a.X*b.Z, a.X*b.Y - a.Y*b.X}
}

func unit_vector(a Vec3) Vec3 {
	return vec_div_scalar(a, vec_len(a))
}

func random_vec3() Vec3 {
	return Vec3{rand.Float64(), rand.Float64(), rand.Float64()}
}

func random_vec3_range(min, max float64) Vec3 {
	return Vec3{
		min + (max-min)*rand.Float64(),
		min + (max-min)*rand.Float64(),
		min + (max-min)*rand.Float64(),
	}
}

func random_in_unit_sphere() Vec3 {
	for {
		p := random_vec3_range(-1, 1)
		if vec_len_squared(p) < 1.0 {
			return p
		}
	}
}

func random_unit_vector() Vec3 {
	return unit_vector(random_in_unit_sphere())
}

func random_in_unit_disk() Vec3 {
	for {
		p := Vec3{rand.Float64()*2 - 1, rand.Float64()*2 - 1, 0}
		if vec_len_squared(p) < 1.0 {
			return p
		}
	}
}

func near_zero(v Vec3) bool {
	s := 1e-8
	return math.Abs(v.X) < s && math.Abs(v.Y) < s && math.Abs(v.Z) < s
}

func reflect(v, n Vec3) Vec3 {
	return vec_sub(v, vec_mul_scalar(n, 2*vec_dot(v, n)))
}

func refract(uv, n Vec3, etai_over_etat float64) Vec3 {
	cos_theta := math.Min(vec_dot(vec_mul_scalar(uv, -1), n), 1.0)
	r_out_perp := vec_mul_scalar(vec_add(uv, vec_mul_scalar(n, cos_theta)), etai_over_etat)
	r_out_parallel := vec_mul_scalar(n, -math.Sqrt(math.Abs(1.0-vec_len_squared(r_out_perp))))
	return vec_add(r_out_perp, r_out_parallel)
}
