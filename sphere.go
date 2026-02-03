package main

import "math"

type sphere struct {
	center Vec3
	radius float64
	mat    material
}

func get_sphere_uv(p Vec3) (float64, float64) {
	theta := math.Acos(-p.Y)
	phi := math.Atan2(-p.Z, p.X) + math.Pi
	u := phi / (2 * math.Pi)
	v := theta / math.Pi
	return u, v
}

func (s sphere) hit(r ray, t_min float64, t_max float64, rec *hit_record) bool {
	oc := vec_sub(r.origin, s.center)
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
	outward_normal := vec_div_scalar(vec_sub(rec.p, s.center), s.radius)
	rec.set_face_normal(r, outward_normal)
	rec.u, rec.v = get_sphere_uv(outward_normal)
	rec.mat = s.mat
	return true
}

func (s sphere) bounding_box(time0, time1 float64) (aabb, bool) {
	return aabb{
		vec_sub(s.center, Vec3{s.radius, s.radius, s.radius}),
		vec_add(s.center, Vec3{s.radius, s.radius, s.radius}),
	}, true
}
