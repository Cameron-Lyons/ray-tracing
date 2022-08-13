package main

import "math"

type sphere struct {
	center Vec3
	radius float32
}

func (s sphere) hit(r ray, t_min float32, t_max float32) bool {
	oc := vec_sub(r.origin, s.center)
	a := vec_dot(r.direction, r.direction)
	b := vec_dot(oc, r.direction)
	c := vec_dot(oc, oc) - s.radius*s.radius
	discriminant := b*b - a*c
	if discriminant > 0 {
		temp := (-b - float32(math.Sqrt(float64(discriminant)))) / a
		if temp < t_max && temp > t_min {
			return true
		}
		temp = (-b + float32(math.Sqrt(float64(discriminant)))) / a
		if temp < t_max && temp > t_min {
			return true
		}
	}
	return false
}
