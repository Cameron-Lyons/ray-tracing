package main

import "math/rand"

type hittable_list struct {
	list []hittable
}

func (l *hittable_list) hit(r ray, t_min float64, t_max float64, rec *hit_record) bool {
	var temp_rec hit_record
	hit_anything := false
	closest_so_far := t_max

	for _, h := range l.list {
		if h.hit(r, t_min, closest_so_far, &temp_rec) {
			hit_anything = true
			closest_so_far = temp_rec.t
			*rec = temp_rec
		}
	}
	return hit_anything
}

func (l *hittable_list) bounding_box(time0, time1 float64) (aabb, bool) {
	if len(l.list) == 0 {
		return aabb{}, false
	}

	var output_box aabb
	first_box := true

	for _, object := range l.list {
		temp_box, has_box := object.bounding_box(time0, time1)
		if !has_box {
			return aabb{}, false
		}
		if first_box {
			output_box = temp_box
		} else {
			output_box = surrounding_box(output_box, temp_box)
		}
		first_box = false
	}
	return output_box, true
}

func (l *hittable_list) pdf_value(o, v Vec3) float64 {
	weight := 1.0 / float64(len(l.list))
	sum := 0.0
	for _, h := range l.list {
		if ph, ok := h.(pdf_hittable); ok {
			sum += weight * ph.pdf_value(o, v)
		}
	}
	return sum
}

func (l *hittable_list) random(o Vec3) Vec3 {
	return l.list[rand.Intn(len(l.list))].(pdf_hittable).random(o)
}
