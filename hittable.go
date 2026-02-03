package main

type hit_record struct {
	t          float64
	u, v       float64
	p          Vec3
	normal     Vec3
	front_face bool
	mat        material
}

func (rec *hit_record) set_face_normal(r ray, outward_normal Vec3) {
	rec.front_face = vec_dot(r.direction, outward_normal) < 0
	if rec.front_face {
		rec.normal = outward_normal
	} else {
		rec.normal = vec_mul_scalar(outward_normal, -1)
	}
}

type hittable interface {
	hit(r ray, t_min float64, t_max float64, rec *hit_record) bool
	bounding_box(time0, time1 float64) (aabb, bool)
}
