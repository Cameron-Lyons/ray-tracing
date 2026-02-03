package main

import (
	"math"
	"math/rand"
)

type camera struct {
	origin            Vec3
	lower_left_corner Vec3
	horizontal        Vec3
	vertical          Vec3
	u, v, w           Vec3
	lens_radius       float64
	time0, time1      float64
}

func new_camera(lookfrom, lookat, vup Vec3, vfov, aspect_ratio, aperture, focus_dist, time0, time1 float64) camera {
	theta := vfov * math.Pi / 180
	h := math.Tan(theta / 2)
	viewport_height := 2.0 * h
	viewport_width := aspect_ratio * viewport_height

	w := unit_vector(vec_sub(lookfrom, lookat))
	u := unit_vector(vec_cross(vup, w))
	v := vec_cross(w, u)

	origin := lookfrom
	horizontal := vec_mul_scalar(u, viewport_width*focus_dist)
	vertical := vec_mul_scalar(v, viewport_height*focus_dist)
	lower_left_corner := vec_sub(vec_sub(vec_sub(origin, vec_mul_scalar(horizontal, 0.5)), vec_mul_scalar(vertical, 0.5)), vec_mul_scalar(w, focus_dist))

	return camera{
		origin:            origin,
		lower_left_corner: lower_left_corner,
		horizontal:        horizontal,
		vertical:          vertical,
		u:                 u,
		v:                 v,
		w:                 w,
		lens_radius:       aperture / 2,
		time0:             time0,
		time1:             time1,
	}
}

func (c camera) get_ray(s, t float64) ray {
	rd := vec_mul_scalar(random_in_unit_disk(), c.lens_radius)
	offset := vec_add(vec_mul_scalar(c.u, rd.X), vec_mul_scalar(c.v, rd.Y))
	return ray{
		vec_add(c.origin, offset),
		vec_sub(
			vec_add(vec_add(c.lower_left_corner, vec_mul_scalar(c.horizontal, s)), vec_mul_scalar(c.vertical, t)),
			vec_add(c.origin, offset),
		),
		c.time0 + rand.Float64()*(c.time1-c.time0),
	}
}
