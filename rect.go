package main

type xy_rect struct {
	x0, x1, y0, y1, k float64
	mat                material
}

func (r xy_rect) hit(ray ray, t_min float64, t_max float64, rec *hit_record) bool {
	t := (r.k - ray.origin.Z) / ray.direction.Z
	if t < t_min || t > t_max {
		return false
	}
	x := ray.origin.X + t*ray.direction.X
	y := ray.origin.Y + t*ray.direction.Y
	if x < r.x0 || x > r.x1 || y < r.y0 || y > r.y1 {
		return false
	}
	rec.u = (x - r.x0) / (r.x1 - r.x0)
	rec.v = (y - r.y0) / (r.y1 - r.y0)
	rec.t = t
	rec.set_face_normal(ray, Vec3{0, 0, 1})
	rec.mat = r.mat
	rec.p = point_at_parameter(ray, t)
	return true
}

func (r xy_rect) bounding_box(time0, time1 float64) (aabb, bool) {
	return aabb{Vec3{r.x0, r.y0, r.k - 0.0001}, Vec3{r.x1, r.y1, r.k + 0.0001}}, true
}

type xz_rect struct {
	x0, x1, z0, z1, k float64
	mat                material
}

func (r xz_rect) hit(ray ray, t_min float64, t_max float64, rec *hit_record) bool {
	t := (r.k - ray.origin.Y) / ray.direction.Y
	if t < t_min || t > t_max {
		return false
	}
	x := ray.origin.X + t*ray.direction.X
	z := ray.origin.Z + t*ray.direction.Z
	if x < r.x0 || x > r.x1 || z < r.z0 || z > r.z1 {
		return false
	}
	rec.u = (x - r.x0) / (r.x1 - r.x0)
	rec.v = (z - r.z0) / (r.z1 - r.z0)
	rec.t = t
	rec.set_face_normal(ray, Vec3{0, 1, 0})
	rec.mat = r.mat
	rec.p = point_at_parameter(ray, t)
	return true
}

func (r xz_rect) bounding_box(time0, time1 float64) (aabb, bool) {
	return aabb{Vec3{r.x0, r.k - 0.0001, r.z0}, Vec3{r.x1, r.k + 0.0001, r.z1}}, true
}

type yz_rect struct {
	y0, y1, z0, z1, k float64
	mat                material
}

func (r yz_rect) hit(ray ray, t_min float64, t_max float64, rec *hit_record) bool {
	t := (r.k - ray.origin.X) / ray.direction.X
	if t < t_min || t > t_max {
		return false
	}
	y := ray.origin.Y + t*ray.direction.Y
	z := ray.origin.Z + t*ray.direction.Z
	if y < r.y0 || y > r.y1 || z < r.z0 || z > r.z1 {
		return false
	}
	rec.u = (y - r.y0) / (r.y1 - r.y0)
	rec.v = (z - r.z0) / (r.z1 - r.z0)
	rec.t = t
	rec.set_face_normal(ray, Vec3{1, 0, 0})
	rec.mat = r.mat
	rec.p = point_at_parameter(ray, t)
	return true
}

func (r yz_rect) bounding_box(time0, time1 float64) (aabb, bool) {
	return aabb{Vec3{r.k - 0.0001, r.y0, r.z0}, Vec3{r.k + 0.0001, r.y1, r.z1}}, true
}
