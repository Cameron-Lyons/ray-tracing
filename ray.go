package main

type ray struct {
	origin    Vec3
	direction Vec3
	time      float64
}

func point_at_parameter(r ray, t float64) Vec3 {
	return vec_add(r.origin, vec_mul_scalar(r.direction, t))
}
