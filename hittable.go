package main

type hit_record struct {
	t          float32
	p          Vec3
	normal     Vec3
	front_face bool
}

type hittable interface {
	hit(r ray, t_min float32, t_max float32, h hit_record) bool
	hit_record(r ray, t_min float32, t_max float32) hit_record
}

func set_face_normal(r ray, outward_normal Vec3) (bool, Vec3) {
	if vec_dot(r.direction, outward_normal) < 0 {
		return true, outward_normal
	}
	return false, vec_mul_scalar(outward_normal, -1)
}
