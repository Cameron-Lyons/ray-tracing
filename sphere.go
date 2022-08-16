package main

import "math"

type sphere struct {
	center Vec3
	radius float32
	mat    material
}

func (s sphere) hit(r ray, t_min float32, t_max float32, rec hit_record) bool {
	oc := vec_sub(r.origin, s.center)
	a := vec_dot(r.direction, r.direction)
	b := vec_dot(oc, r.direction)
	c := vec_dot(oc, oc) - s.radius*s.radius
	discriminant := b*b - a*c
	if discriminant > 0 {
		root := (-b - float32(math.Sqrt(float64(discriminant)))) / a
		if root < t_min || t_max < root {
			return false
		}
		rec.t = root
		rec.p = point_at_parameter(r, root)
		outward_normal := vec_div_scalar(vec_sub(rec.p, s.center), s.radius)
		rec.front_face, rec.normal = set_face_normal(r, outward_normal)
		rec.mat = s.mat
		return true
	}
	return false
}
