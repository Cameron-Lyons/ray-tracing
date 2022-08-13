package main

type hit_record struct {
	t      float32
	p      Vec3
	normal Vec3
}

type hittable interface {
	hit(r ray, t_min float32, t_max float32) bool
	hit_record(r ray, t_min float32, t_max float32) hit_record
}
