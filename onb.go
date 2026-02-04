package main

import "math"

type onb struct {
	u, v, w Vec3
}

func new_onb(n Vec3) onb {
	w := unit_vector(n)
	var a Vec3
	if math.Abs(w.X) > 0.9 {
		a = Vec3{0, 1, 0}
	} else {
		a = Vec3{1, 0, 0}
	}
	v := unit_vector(vec_cross(w, a))
	u := vec_cross(w, v)
	return onb{u, v, w}
}

func (o onb) local(a Vec3) Vec3 {
	return vec_add(vec_add(vec_mul_scalar(o.u, a.X), vec_mul_scalar(o.v, a.Y)), vec_mul_scalar(o.w, a.Z))
}
