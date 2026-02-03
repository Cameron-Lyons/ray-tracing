package main

type box_shape struct {
	box_min, box_max Vec3
	sides            hittable_list
}

func new_box(p0, p1 Vec3, mat material) *box_shape {
	b := &box_shape{box_min: p0, box_max: p1}

	b.sides.list = append(b.sides.list, xy_rect{p0.X, p1.X, p0.Y, p1.Y, p1.Z, mat})
	b.sides.list = append(b.sides.list, xy_rect{p0.X, p1.X, p0.Y, p1.Y, p0.Z, mat})

	b.sides.list = append(b.sides.list, xz_rect{p0.X, p1.X, p0.Z, p1.Z, p1.Y, mat})
	b.sides.list = append(b.sides.list, xz_rect{p0.X, p1.X, p0.Z, p1.Z, p0.Y, mat})

	b.sides.list = append(b.sides.list, yz_rect{p0.Y, p1.Y, p0.Z, p1.Z, p1.X, mat})
	b.sides.list = append(b.sides.list, yz_rect{p0.Y, p1.Y, p0.Z, p1.Z, p0.X, mat})

	return b
}

func (b *box_shape) hit(r ray, t_min float64, t_max float64, rec *hit_record) bool {
	return b.sides.hit(r, t_min, t_max, rec)
}

func (b *box_shape) bounding_box(time0, time1 float64) (aabb, bool) {
	return aabb{b.box_min, b.box_max}, true
}
