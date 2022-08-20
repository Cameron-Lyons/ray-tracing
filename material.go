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
		attenuation = Color{1.0, 1.0, 1.0}
		reflected = reflect(ray.direction, rec.normal)
		var outward_normal Vec3
		var ni_over_nt float64
		var cosine float64
		if ray.direction.dot(rec.normal) > 0.0 {
			outward_normal = -rec.normal
			ni_over_nt = ref_idx
			cosine = ref_idx * ray.direction.dot(rec.normal) / ray.direction.len()
		} else {
			outward_normal = rec.normal
			ni_over_nt = 1.0 / ref_idx
			cosine = -ray.direction.dot(rec.normal) / ray.direction.len()
		}
		var refracted Vec3
		var reflect_prob float64
		var refracted_prob float64
		if refract(ray.direction, outward_normal, ni_over_nt, &refracted) {
			reflect_prob = schlick(cosine, ref_idx)
			refracted_prob = 1.0 - reflect_prob
		} else {
			scattered = ray(rec.p, reflected)
			return true
		}
		if rand.Float64() < reflect_prob {
			scattered = ray(rec.p, reflected)
		} else {
			scattered = ray(rec.p, refracted)
		}
		return true
	}

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
