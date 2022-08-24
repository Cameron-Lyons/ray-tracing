package main

type camera struct {
	lookfrom Vec3
	lookat Vec3
	vup Vec3
	vfov float32
	aspect_ratio float32
	theta := vfov * (math.Pi / 180)
	h := math.Tan(theta / 2)
	viewport_hieght = 2 * half_height
	viewport_width = aspect_ratio * viewport_height
	
	w = vec_unit_vector(vec_sub(lookat, lookfrom))
	u = vec_unit_vector(vec_cross(vup, w))
	v = vec_cross(w, u)
	
	origin = lookfrom
	horizontal = viewport_width * u
	vertical = viewport_height * v
	lower_left_corner = vec_sub(origin, vec_mul_scalar(horizontal, 0.5), vec_mul_scalar(vertical, 0.5), vec_mul_scalar(w, focal_length))

	ray get_ray(s, t) ray {
		return ray(origin, vec_add(vec_add(vec_add(lower_left_corner, vec_mul_scalar(horizontal, s)), vec_mul_scalar(vertical, t)), Vec3{0, 0, focal_length}))
	}
}

func get_ray(u, v) ray {
	return ray{origin, vec_add(vec_add(vec_add(lower_left_corner, vec_mul_scalar(horizontal, u)), vec_mul_scalar(vertical, v)), Vec3{0, 0, focal_length})}
}