package main

type camera struct {
	vfov float32
	aspect_ratio float32
	theta = vfov * (math.Pi / 180)
	half_height = math.Tan(theta / 2)
	viewport_hieght = 2 * half_height
	viewport_width = aspect_ratio * viewport_height
	focal_length = 1.0
	
	origin = Vec3{0, 0, 0}
	horizontal = Vec3{viewport_width, 0, 0}
	vertical = Vec3{0, viewport_hieght, 0}
	lower_left_corner = Vec3{-viewport_width / 2, -viewport_hieght / 2, -focal_length}

	ray get_ray(u, v) ray {
		origin,
		direction := vec_add(vec_add(vec_add(lower_left_corner, vec_mul_scalar(horizontal, u)), vec_mul_scalar(vertical, v)), Vec3{0, 0, focal_length})
		direction = vec_normalize(direction)
		return ray{origin, direction}
	}
}

func get_ray(u, v) ray {
	return ray{origin, vec_add(vec_add(vec_add(lower_left_corner, vec_mul_scalar(horizontal, u)), vec_mul_scalar(vertical, v)), Vec3{0, 0, focal_length})}
}