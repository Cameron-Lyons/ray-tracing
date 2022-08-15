package main

type camera struct {
	aspectRatio := 16.0 / 9.0
	viewport_hieght := 2.0
	viewport_width := aspectRatio * viewport_hieght
	focal_length := 1.0

	origin := Vec3{0, 0, 0}
	horizontal := Vec3{viewport_width, 0, 0}
	vertical := Vec3{0, viewport_hieght, 0}
	lower_left_corner := Vec3{-viewport_width / 2, -viewport_hieght / 2, -focal_length}
	ray ray
}

func get_ray(u, v) ray {
	return ray{origin, vec_add(vec_add(vec_add(lower_left_corner, vec_mul_scalar(horizontal, u)), vec_mul_scalar(vertical, v)), Vec3{0, 0, focal_length})}
}