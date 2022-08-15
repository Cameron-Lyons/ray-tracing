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

func ray_color(r ray, world hittable) Vec3 {
	var rec hit_record
	if world.hit(r, 0.0, math.MaxFloat32, rec) {
		return vec_mul_scalar(vec_add(rec.normal, Vec3{1, 1, 1}), 0.5)
	}
	unit_direction := unit_vector(r.direction)
	t := 0.5 * (unit_direction.Y + 1.0)
	return vec_add(vec_mul_scalar(Vec3{1.0, 1.0, 1.0}, (1.0-t)), vec_mul_scalar(Vec3{0.5, 0.7, 1.0}, t))
}

func main() {
	// Image
	const aspectRatio = 16.0 / 9.0
	const image_width int = 256
	const image_height int = int(float32(image_width) / aspectRatio)

	image := image.NewRGBA(image.Rect(0, 0, image_width, image_height))

	// World
	var world hittable_list
	world.list = append(world.list, sphere{Vec3{0, 0, -1}, 0.5})
	world.list = append(world.list, sphere{Vec3{0, -100.5, -1}, 100})

	// Camera
	const viewport_hieght = 2.0
	const viewport_width = aspectRatio * viewport_hieght
	const focal_length = 1.0

	origin = Vec3{0, 0, 0}
	horizontal := Vec3{viewport_width, 0, 0}
	vertical := Vec3{0, viewport_hieght, 0}
	lower_left_corner := Vec3{-viewport_width / 2, -viewport_hieght / 2, -focal_length}

	// Render
	for j := image_height - 1; j >= 0; j-- {
		for i := 0; i < image_width; i++ {

			r := float32(i) / float32(image_width-1)
			g := float32(j) / float32(image_height-1)
			b := 0.25

			ir := int(255.999 * r)
			ig := int(255.999 * g)
			ib := int(255.990 * b)

			image.Set(i, j, color.RGBA{uint8(ir), uint8(ig), uint8(ib), 255})
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
