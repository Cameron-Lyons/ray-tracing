package main

import "math"

type moving_sphere struct {
	center0, center1 Vec3
	time0, time1     float64
	radius           float64
	mat              material
}

func (s moving_sphere) center(time float64) Vec3 {
	return vec_add(s.center0, vec_mul_scalar(vec_sub(s.center1, s.center0), (time-s.time0)/(s.time1-s.time0)))
}

func (s moving_sphere) hit(r ray, t_min float64, t_max float64, rec *hit_record) bool {
	center := s.center(r.time)
	oc := vec_sub(r.origin, center)
	a := vec_len_squared(r.direction)
	half_b := vec_dot(oc, r.direction)
	c := vec_len_squared(oc) - s.radius*s.radius
	discriminant := half_b*half_b - a*c

	if discriminant < 0 {
		return false
	}

	sqrtd := math.Sqrt(discriminant)

	root := (-half_b - sqrtd) / a
	if root < t_min || t_max < root {
		root = (-half_b + sqrtd) / a
		if root < t_min || t_max < root {
			return false
		}
	}

	rec.t = root
	rec.p = point_at_parameter(r, rec.t)
	outward_normal := vec_div_scalar(vec_sub(rec.p, center), s.radius)
	rec.set_face_normal(r, outward_normal)
	rec.u, rec.v = get_sphere_uv(outward_normal)
	rec.mat = s.mat
	return true
}

func (s moving_sphere) bounding_box(time0, time1 float64) (aabb, bool) {
	box0 := aabb{
		vec_sub(s.center(time0), Vec3{s.radius, s.radius, s.radius}),
		vec_add(s.center(time0), Vec3{s.radius, s.radius, s.radius}),
	}
	box1 := aabb{
		vec_sub(s.center(time1), Vec3{s.radius, s.radius, s.radius}),
		vec_add(s.center(time1), Vec3{s.radius, s.radius, s.radius}),
	}
	return surrounding_box(box0, box1), true
}
