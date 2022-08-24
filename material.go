package main

type material struct {
	scatter func(ray, hit_record, attenuation Color) (Vec3, ray)
}

type lambertian struct {
	albedo  Color
	scatter func(ray, hit_record, attenuation Color) bool {
		if near_zero(scatter_direction){
			scatter_direction = rec.normal
		}
		scattered = ray(rec.p, scatter_direction)
		attenuation = albedo
		return true
	}
	dielectric float64
	albedo Color
}

type metal struct {
	albedo  Color
	fuzz    float64
	scatter func(ray, hit_record, attenuation Color) bool {
		reflected = reflect(ray.direction, rec.normal)
		scattered = ray(rec.p, reflected + fuzz*random_in_unit_sphere())
		attenuation = albedo
		return scattered.direction.dot(rec.normal) > 0.0
	}
}

type dielectric struct {
	ref_idx float64
	scatter func(ray, hit_record, attenuation Color) bool {
		attenuation := Color{1.0, 1.0, 1.0}
		refraction_ratio := rec.front_face

		unit_direction := unit_vector(scattered.direction)
		cos_theta := math.Min(dot(unit_direction, rec.normal), 1.0)
		sin_theta := math.Sqrt(1.0 - cos_theta * cos_theta)

		cannot_refract := refraction_ratio * sin_theta > 1.0
		
		if cannot_refract {
			direction := reflect(unit_direction, rec.normal)
	}
		else {
			direction := refract(unit_direction, rec.normal, refraction_ratio)
		}
		scatterd := ray(rec.p, scattered.direction)

		return true

	func reflectance(ray, hit_record, attenuation Color) float64 {
		r0 := (1.0 - ref_idx) / (1.0 + ref_idx)
		r0 = r0 * r0
		return r0 + (1.0 - r0) * math.Pow(1.0 - cosine, 5)
	}

func (m *material) scatter(ray, hit_record, attenuation Color) (Vec3, ray) {
	return m.scatter(ray, hit_record, attenuation)
}

func (m *lambertian) scatter(ray, hit_record, attenuation Color) (Vec3, ray) {
	return m.scatter(ray, hit_record, attenuation)
}

func (m *metal) scatter
