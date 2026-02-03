package main

import (
	"math"
	"math/rand"
)

type material interface {
	scatter(r_in ray, rec hit_record, attenuation *Vec3, scattered *ray) bool
	emitted(u, v float64, p Vec3) Vec3
}

type lambertian struct {
	albedo texture
}

func (l lambertian) scatter(r_in ray, rec hit_record, attenuation *Vec3, scattered *ray) bool {
	scatter_direction := vec_add(rec.normal, random_unit_vector())
	if near_zero(scatter_direction) {
		scatter_direction = rec.normal
	}
	*scattered = ray{rec.p, scatter_direction, r_in.time}
	*attenuation = l.albedo.value(rec.u, rec.v, rec.p)
	return true
}

func (l lambertian) emitted(u, v float64, p Vec3) Vec3 {
	return Vec3{0, 0, 0}
}

type metal struct {
	albedo Vec3
	fuzz   float64
}

func (m metal) scatter(r_in ray, rec hit_record, attenuation *Vec3, scattered *ray) bool {
	reflected := reflect(unit_vector(r_in.direction), rec.normal)
	*scattered = ray{rec.p, vec_add(reflected, vec_mul_scalar(random_in_unit_sphere(), m.fuzz)), r_in.time}
	*attenuation = m.albedo
	return vec_dot(scattered.direction, rec.normal) > 0
}

func (m metal) emitted(u, v float64, p Vec3) Vec3 {
	return Vec3{0, 0, 0}
}

type dielectric struct {
	ref_idx float64
}

func (d dielectric) scatter(r_in ray, rec hit_record, attenuation *Vec3, scattered *ray) bool {
	*attenuation = Vec3{1.0, 1.0, 1.0}
	var refraction_ratio float64
	if rec.front_face {
		refraction_ratio = 1.0 / d.ref_idx
	} else {
		refraction_ratio = d.ref_idx
	}

	unit_direction := unit_vector(r_in.direction)
	cos_theta := math.Min(vec_dot(vec_mul_scalar(unit_direction, -1), rec.normal), 1.0)
	sin_theta := math.Sqrt(1.0 - cos_theta*cos_theta)

	cannot_refract := refraction_ratio*sin_theta > 1.0

	var direction Vec3
	if cannot_refract || reflectance(cos_theta, refraction_ratio) > rand.Float64() {
		direction = reflect(unit_direction, rec.normal)
	} else {
		direction = refract(unit_direction, rec.normal, refraction_ratio)
	}

	*scattered = ray{rec.p, direction, r_in.time}
	return true
}

func (d dielectric) emitted(u, v float64, p Vec3) Vec3 {
	return Vec3{0, 0, 0}
}

func reflectance(cosine, ref_idx float64) float64 {
	r0 := (1 - ref_idx) / (1 + ref_idx)
	r0 = r0 * r0
	return r0 + (1-r0)*math.Pow(1-cosine, 5)
}

type diffuse_light struct {
	emit texture
}

func (d diffuse_light) scatter(r_in ray, rec hit_record, attenuation *Vec3, scattered *ray) bool {
	return false
}

func (d diffuse_light) emitted(u, v float64, p Vec3) Vec3 {
	return d.emit.value(u, v, p)
}
