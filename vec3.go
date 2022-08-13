package vec3

import "math"

type Vec3 struct {
	X float32
	Y float32
	Z float32
} // Vec3 is a 3D vector

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
