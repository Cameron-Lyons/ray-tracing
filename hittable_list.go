package main

type hittable_list struct {
	list []hittable
}

func (l hittable_list) hit(r ray, t_min float32, t_max float32, rec hit_record) bool {
	var temp_rec hit_record
	var hit_anything bool = false
	var closest_so_far float32 = t_max

	for _, h := range l.list {
		if h.hit(r, t_min, closest_so_far, temp_rec) {
			hit_anything = true
			closest_so_far = temp_rec.t
			rec = temp_rec
		}
	}
	return hit_anything
}
