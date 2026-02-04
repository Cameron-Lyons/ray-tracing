package main

import "math"

type flip_face struct {
	ptr hittable
}

func (f flip_face) hit(r ray, t_min float64, t_max float64, rec *hit_record) bool {
	if !f.ptr.hit(r, t_min, t_max, rec) {
		return false
	}
	rec.front_face = !rec.front_face
	return true
}

func (f flip_face) bounding_box(time0, time1 float64) (aabb, bool) {
	return f.ptr.bounding_box(time0, time1)
}

type translate struct {
	ptr    hittable
	offset Vec3
}

func (t translate) hit(r ray, t_min float64, t_max float64, rec *hit_record) bool {
	moved_r := ray{vec_sub(r.origin, t.offset), r.direction, r.time}
	if !t.ptr.hit(moved_r, t_min, t_max, rec) {
		return false
	}
	rec.p = vec_add(rec.p, t.offset)
	rec.set_face_normal(moved_r, rec.normal)
	return true
}

func (t translate) bounding_box(time0, time1 float64) (aabb, bool) {
	output_box, has_box := t.ptr.bounding_box(time0, time1)
	if !has_box {
		return aabb{}, false
	}
	return aabb{
		vec_add(output_box.minimum, t.offset),
		vec_add(output_box.maximum, t.offset),
	}, true
}

type rotate_y struct {
	ptr                hittable
	sin_theta          float64
	cos_theta          float64
	has_box            bool
	bbox               aabb
}

func new_rotate_y(p hittable, angle float64) rotate_y {
	radians := angle * math.Pi / 180
	sin_theta := math.Sin(radians)
	cos_theta := math.Cos(radians)

	bbox, has_box := p.bounding_box(0, 1)

	min := Vec3{math.Inf(1), math.Inf(1), math.Inf(1)}
	max := Vec3{math.Inf(-1), math.Inf(-1), math.Inf(-1)}

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				fi := float64(i)
				fj := float64(j)
				fk := float64(k)
				x := fi*bbox.maximum.X + (1-fi)*bbox.minimum.X
				y := fj*bbox.maximum.Y + (1-fj)*bbox.minimum.Y
				z := fk*bbox.maximum.Z + (1-fk)*bbox.minimum.Z

				newx := cos_theta*x + sin_theta*z
				newz := -sin_theta*x + cos_theta*z

				if newx < min.X {
					min.X = newx
				}
				if newx > max.X {
					max.X = newx
				}
				if y < min.Y {
					min.Y = y
				}
				if y > max.Y {
					max.Y = y
				}
				if newz < min.Z {
					min.Z = newz
				}
				if newz > max.Z {
					max.Z = newz
				}
			}
		}
	}

	return rotate_y{
		ptr:       p,
		sin_theta: sin_theta,
		cos_theta: cos_theta,
		has_box:   has_box,
		bbox:      aabb{min, max},
	}
}

func (ry rotate_y) hit(r ray, t_min float64, t_max float64, rec *hit_record) bool {
	origin := Vec3{
		ry.cos_theta*r.origin.X - ry.sin_theta*r.origin.Z,
		r.origin.Y,
		ry.sin_theta*r.origin.X + ry.cos_theta*r.origin.Z,
	}
	direction := Vec3{
		ry.cos_theta*r.direction.X - ry.sin_theta*r.direction.Z,
		r.direction.Y,
		ry.sin_theta*r.direction.X + ry.cos_theta*r.direction.Z,
	}

	rotated_r := ray{origin, direction, r.time}

	if !ry.ptr.hit(rotated_r, t_min, t_max, rec) {
		return false
	}

	p := Vec3{
		ry.cos_theta*rec.p.X + ry.sin_theta*rec.p.Z,
		rec.p.Y,
		-ry.sin_theta*rec.p.X + ry.cos_theta*rec.p.Z,
	}
	normal := Vec3{
		ry.cos_theta*rec.normal.X + ry.sin_theta*rec.normal.Z,
		rec.normal.Y,
		-ry.sin_theta*rec.normal.X + ry.cos_theta*rec.normal.Z,
	}

	rec.p = p
	rec.set_face_normal(rotated_r, normal)
	return true
}

func (ry rotate_y) bounding_box(time0, time1 float64) (aabb, bool) {
	return ry.bbox, ry.has_box
}
