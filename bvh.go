package main

import (
	"math/rand"
	"sort"
)

type bvh_node struct {
	left, right hittable
	box         aabb
}

func new_bvh_node(objects []hittable, time0, time1 float64) *bvh_node {
	node := &bvh_node{}

	axis := rand.Intn(3)
	comparator := func(a, b hittable) bool {
		box_a, _ := a.bounding_box(time0, time1)
		box_b, _ := b.bounding_box(time0, time1)
		switch axis {
		case 0:
			return box_a.minimum.X < box_b.minimum.X
		case 1:
			return box_a.minimum.Y < box_b.minimum.Y
		default:
			return box_a.minimum.Z < box_b.minimum.Z
		}
	}

	switch n := len(objects); n {
	case 1:
		node.left = objects[0]
		node.right = objects[0]
	case 2:
		if comparator(objects[0], objects[1]) {
			node.left = objects[0]
			node.right = objects[1]
		} else {
			node.left = objects[1]
			node.right = objects[0]
		}
	default:
		sort.Slice(objects, func(i, j int) bool {
			return comparator(objects[i], objects[j])
		})
		mid := n / 2
		node.left = new_bvh_node(objects[:mid], time0, time1)
		node.right = new_bvh_node(objects[mid:], time0, time1)
	}

	box_left, _ := node.left.bounding_box(time0, time1)
	box_right, _ := node.right.bounding_box(time0, time1)
	node.box = surrounding_box(box_left, box_right)

	return node
}

func (n *bvh_node) hit(r ray, t_min float64, t_max float64, rec *hit_record) bool {
	if !aabb_hit(n.box, r, t_min, t_max) {
		return false
	}
	hit_left := n.left.hit(r, t_min, t_max, rec)
	if hit_left {
		t_max = rec.t
	}
	hit_right := n.right.hit(r, t_min, t_max, rec)
	return hit_left || hit_right
}

func (n *bvh_node) bounding_box(time0, time1 float64) (aabb, bool) {
	return n.box, true
}
