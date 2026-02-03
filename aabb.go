package main

import "math"

type aabb struct {
	minimum, maximum Vec3
}

func aabb_hit(box aabb, r ray, t_min, t_max float64) bool {
	for a := 0; a < 3; a++ {
		var min_val, max_val, origin, dir float64
		switch a {
		case 0:
			min_val, max_val, origin, dir = box.minimum.X, box.maximum.X, r.origin.X, r.direction.X
		case 1:
			min_val, max_val, origin, dir = box.minimum.Y, box.maximum.Y, r.origin.Y, r.direction.Y
		case 2:
			min_val, max_val, origin, dir = box.minimum.Z, box.maximum.Z, r.origin.Z, r.direction.Z
		}
		invD := 1.0 / dir
		t0 := (min_val - origin) * invD
		t1 := (max_val - origin) * invD
		if invD < 0 {
			t0, t1 = t1, t0
		}
		if t0 > t_min {
			t_min = t0
		}
		if t1 < t_max {
			t_max = t1
		}
		if t_max <= t_min {
			return false
		}
	}
	return true
}

func surrounding_box(box0, box1 aabb) aabb {
	small := Vec3{
		math.Min(box0.minimum.X, box1.minimum.X),
		math.Min(box0.minimum.Y, box1.minimum.Y),
		math.Min(box0.minimum.Z, box1.minimum.Z),
	}
	big := Vec3{
		math.Max(box0.maximum.X, box1.maximum.X),
		math.Max(box0.maximum.Y, box1.maximum.Y),
		math.Max(box0.maximum.Z, box1.maximum.Z),
	}
	return aabb{small, big}
}
