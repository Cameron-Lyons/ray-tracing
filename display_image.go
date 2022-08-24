package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math"
	"os"
)

func ray_color(r ray, world hittable, depth int) Vec3 {
	var rec hit_record

	if depth <= 0 {
		return Vec3{0, 0, 0}
	}

	if world.hit(r, 0.001, math.MaxFloat32, rec) {
		scattered := ray{}
		attenuation := Color{}
		if rec.material.scatter(r, rec, attenuation, scattered) {
			return vec_mul_scalar(ray_color(scattered, world, depth-1), attenuation)
		}
		return Color(0, 0, 0)
	}
	unit_direction := unit_vector(r.direction)
	t := 0.5 * (unit_direction.Y + 1.0)
	return vec_add(vec_mul_scalar(Vec3{1.0, 1.0, 1.0}, (1.0-t)), vec_mul_scalar(Vec3{0.5, 0.7, 1.0}, t))
}

func main() {
	// Image
	const aspectRatio = 16.0 / 9.0
	const image_width int = 400
	const image_height int = int(float32(image_width) / aspectRatio)
	const samples_per_pixel int = 100
	const max_depth int = 50

	image := image.NewRGBA(image.Rect(0, 0, image_width, image_height))

	// World
	var world hittable_list

	R := math.Cos(math.Pi / 4)

	material_left := lambertian{Color{0, 0, 1}, func(ray, hit_record, attenuation Color) bool {}}
	material_right := lambertian{Color{1, 0, 0}, func(ray, hit_record, attenuation Color) bool {}}

	world.list = append(world.list, sphere{Vec3{0, 0, -1}, 0.5, material_left})
	world.list = append(world.list, sphere{Vec3{0, 0, -1}, -0.45, material_right})
	// Camera
	cam := camera{90, aspectRatio, 0.1}

	// Render
	for j := image_height - 1; j >= 0; j-- {
		pixel_color := Color(0, 0, 0)
		for i := 0; i < image_width; i++ {
			u := float32(i) / float32(image_width-1)
			v := float32(j) / float32(image_height-1)
			ray := ray{origin, vec_add(vec_add(vec_add(lower_left_corner, vec_mul_scalar(horizontal, u)), vec_mul_scalar(vertical, v)), Vec3{0, 0, focal_length})}
			pixel_color += ray_color(ray, world, max_depth)
			image.Set(i, j, color.RGBA{uint8(255.99 * pixel_color.X), uint8(255.99 * pixel_color.Y), uint8(255.99 * pixel_color.Z), 255})
		}
	}
	draw.Draw(image, image.Bounds(), image, image.Bounds().Min, draw.Src)
	out, err := os.Create("./output.png")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	png.Encode(out, image)
}
