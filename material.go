package main

import (
	"math"
	"math/rand"
)

type material interface {
	scatter(r_in ray, rec hit_record, srec *scatter_record) bool
	scattering_pdf(r_in ray, rec hit_record, scattered ray) float64
	emitted(r_in ray, rec hit_record, u, v float64, p Vec3) Vec3
}

type lambertian struct {
	albedo texture
}

func (l lambertian) scatter(r_in ray, rec hit_record, srec *scatter_record) bool {
	srec.is_specular = false
	srec.attenuation = l.albedo.value(rec.u, rec.v, rec.p)
	srec.pdf_ptr = new_cosine_pdf(rec.normal)
	return true
}

func (l lambertian) scattering_pdf(r_in ray, rec hit_record, scattered ray) float64 {
	cosine := vec_dot(rec.normal, unit_vector(scattered.direction))
	if cosine < 0 {
		return 0
	}
	return cosine / math.Pi
}

func (l lambertian) emitted(r_in ray, rec hit_record, u, v float64, p Vec3) Vec3 {
	return Vec3{0, 0, 0}
}

type metal struct {
	albedo Vec3
	fuzz   float64
}

func (m metal) scatter(r_in ray, rec hit_record, srec *scatter_record) bool {
	reflected := reflect(unit_vector(r_in.direction), rec.normal)
	srec.specular_ray = ray{rec.p, vec_add(reflected, vec_mul_scalar(random_in_unit_sphere(), m.fuzz)), r_in.time}
	srec.attenuation = m.albedo
	srec.is_specular = true
	srec.pdf_ptr = nil
	return true
}

func (m metal) scattering_pdf(r_in ray, rec hit_record, scattered ray) float64 {
	return 0
}

func (m metal) emitted(r_in ray, rec hit_record, u, v float64, p Vec3) Vec3 {
	return Vec3{0, 0, 0}
}

type dielectric struct {
	ref_idx float64
}

func (d dielectric) scatter(r_in ray, rec hit_record, srec *scatter_record) bool {
	srec.is_specular = true
	srec.pdf_ptr = nil
	srec.attenuation = Vec3{1.0, 1.0, 1.0}

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

	srec.specular_ray = ray{rec.p, direction, r_in.time}
	return true
}

func (d dielectric) scattering_pdf(r_in ray, rec hit_record, scattered ray) float64 {
	return 0
}

func (d dielectric) emitted(r_in ray, rec hit_record, u, v float64, p Vec3) Vec3 {
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

func (d diffuse_light) scatter(r_in ray, rec hit_record, srec *scatter_record) bool {
	return false
}

func (d diffuse_light) scattering_pdf(r_in ray, rec hit_record, scattered ray) float64 {
	return 0
}

func (d diffuse_light) emitted(r_in ray, rec hit_record, u, v float64, p Vec3) Vec3 {
	if rec.front_face {
		return d.emit.value(u, v, p)
	}
	return Vec3{0, 0, 0}
}
