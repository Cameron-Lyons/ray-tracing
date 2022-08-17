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
}

