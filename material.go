package main

type material struct {
	scatter func(ray, hit_record, attenuation Color) (Vec3, ray)
}