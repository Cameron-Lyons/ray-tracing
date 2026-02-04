package main

import (
	"math"
	"math/rand"
)

type isotropic struct {
	albedo texture
}

func (iso isotropic) scatter(r_in ray, rec hit_record, srec *scatter_record) bool {
	srec.is_specular = true
	srec.specular_ray = ray{rec.p, random_in_unit_sphere(), r_in.time}
	srec.attenuation = iso.albedo.value(rec.u, rec.v, rec.p)
	srec.pdf_ptr = nil
	return true
}

func (iso isotropic) scattering_pdf(r_in ray, rec hit_record, scattered ray) float64 {
	return 0
}

func (iso isotropic) emitted(r_in ray, rec hit_record, u, v float64, p Vec3) Vec3 {
	return Vec3{0, 0, 0}
}

type constant_medium struct {
	boundary        hittable
	neg_inv_density float64
	phase_function  material
}

func new_constant_medium(boundary hittable, density float64, albedo texture) constant_medium {
	return constant_medium{
		boundary:        boundary,
		neg_inv_density: -1.0 / density,
		phase_function:  isotropic{albedo},
	}
}

func (cm constant_medium) hit(r ray, t_min float64, t_max float64, rec *hit_record) bool {
	var rec1, rec2 hit_record

	if !cm.boundary.hit(r, -math.MaxFloat64, math.MaxFloat64, &rec1) {
		return false
	}

	if !cm.boundary.hit(r, rec1.t+0.0001, math.MaxFloat64, &rec2) {
		return false
	}

	if rec1.t < t_min {
		rec1.t = t_min
	}
	if rec2.t > t_max {
		rec2.t = t_max
	}

	if rec1.t >= rec2.t {
		return false
	}

	if rec1.t < 0 {
		rec1.t = 0
	}

	ray_length := vec_len(r.direction)
	distance_inside_boundary := (rec2.t - rec1.t) * ray_length
	hit_distance := cm.neg_inv_density * math.Log(rand.Float64())

	if hit_distance > distance_inside_boundary {
		return false
	}

	rec.t = rec1.t + hit_distance/ray_length
	rec.p = point_at_parameter(r, rec.t)
	rec.normal = Vec3{1, 0, 0}
	rec.front_face = true
	rec.mat = cm.phase_function
	return true
}

func (cm constant_medium) bounding_box(time0, time1 float64) (aabb, bool) {
	return cm.boundary.bounding_box(time0, time1)
}
