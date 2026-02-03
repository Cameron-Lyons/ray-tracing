package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/rand"
	"os"
	"sync"
	"sync/atomic"
)

func ray_color(r ray, world hittable, background Vec3, depth int) Vec3 {
	if depth <= 0 {
		return Vec3{0, 0, 0}
	}

	var rec hit_record
	if !world.hit(r, 0.001, math.MaxFloat64, &rec) {
		return background
	}

	emitted := rec.mat.emitted(rec.u, rec.v, rec.p)

	var scattered ray
	var attenuation Vec3
	if !rec.mat.scatter(r, rec, &attenuation, &scattered) {
		return emitted
	}

	return vec_add(emitted, vec_mul(attenuation, ray_color(scattered, world, background, depth-1)))
}

func random_scene() hittable {
	world := &hittable_list{}

	ground_material := lambertian{checker_texture{solid_color{Vec3{0.2, 0.3, 0.1}}, solid_color{Vec3{0.9, 0.9, 0.9}}}}
	world.list = append(world.list, sphere{Vec3{0, -1000, 0}, 1000, ground_material})

	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			choose_mat := rand.Float64()
			center := Vec3{float64(a) + 0.9*rand.Float64(), 0.2, float64(b) + 0.9*rand.Float64()}

			if vec_len(vec_sub(center, Vec3{4, 0.2, 0})) > 0.9 {
				if choose_mat < 0.8 {
					albedo := vec_mul(random_vec3(), random_vec3())
					center2 := vec_add(center, Vec3{0, rand.Float64() * 0.5, 0})
					world.list = append(world.list, moving_sphere{center, center2, 0.0, 1.0, 0.2, lambertian{solid_color{albedo}}})
				} else if choose_mat < 0.95 {
					albedo := random_vec3_range(0.5, 1)
					fuzz := rand.Float64() * 0.5
					world.list = append(world.list, sphere{center, 0.2, metal{albedo, fuzz}})
				} else {
					world.list = append(world.list, sphere{center, 0.2, dielectric{1.5}})
				}
			}
		}
	}

	world.list = append(world.list, sphere{Vec3{0, 1, 0}, 1.0, dielectric{1.5}})
	world.list = append(world.list, sphere{Vec3{-4, 1, 0}, 1.0, lambertian{noise_texture{new_perlin(), 4}}})
	world.list = append(world.list, sphere{Vec3{4, 1, 0}, 1.0, metal{Vec3{0.7, 0.6, 0.5}, 0.0}})

	return new_bvh_node(world.list, 0.0, 1.0)
}

func cornell_smoke() hittable {
	world := &hittable_list{}

	red := lambertian{solid_color{Vec3{0.65, 0.05, 0.05}}}
	white := lambertian{solid_color{Vec3{0.73, 0.73, 0.73}}}
	green := lambertian{solid_color{Vec3{0.12, 0.45, 0.15}}}
	light := diffuse_light{solid_color{Vec3{7, 7, 7}}}

	world.list = append(world.list, yz_rect{0, 555, 0, 555, 555, green})
	world.list = append(world.list, yz_rect{0, 555, 0, 555, 0, red})
	world.list = append(world.list, xz_rect{113, 443, 127, 432, 554, light})
	world.list = append(world.list, xz_rect{0, 555, 0, 555, 555, white})
	world.list = append(world.list, xz_rect{0, 555, 0, 555, 0, white})
	world.list = append(world.list, xy_rect{0, 555, 0, 555, 555, white})

	box1 := new_box(Vec3{0, 0, 0}, Vec3{165, 330, 165}, white)
	box1_rotated := new_rotate_y(box1, 15)
	box1_translated := translate{box1_rotated, Vec3{265, 0, 295}}

	box2 := new_box(Vec3{0, 0, 0}, Vec3{165, 165, 165}, white)
	box2_rotated := new_rotate_y(box2, -18)
	box2_translated := translate{box2_rotated, Vec3{130, 0, 65}}

	world.list = append(world.list, new_constant_medium(box1_translated, 0.01, solid_color{Vec3{0, 0, 0}}))
	world.list = append(world.list, new_constant_medium(box2_translated, 0.01, solid_color{Vec3{1, 1, 1}}))

	return new_bvh_node(world.list, 0.0, 1.0)
}

func main() {
	scene := 2

	var world hittable
	var background Vec3
	var lookfrom, lookat Vec3
	var vfov, aperture float64
	aspect_ratio := 1.0
	image_width := 600
	samples_per_pixel := 200
	max_depth := 50

	switch scene {
	case 1:
		world = random_scene()
		background = Vec3{0.7, 0.8, 1.0}
		lookfrom = Vec3{13, 2, 3}
		lookat = Vec3{0, 0, 0}
		vfov = 20
		aperture = 0.1
		aspect_ratio = 3.0 / 2.0
		image_width = 400
		samples_per_pixel = 100
	default:
		world = cornell_smoke()
		background = Vec3{0, 0, 0}
		lookfrom = Vec3{278, 278, -800}
		lookat = Vec3{278, 278, 0}
		vfov = 40
		aperture = 0.0
	}

	vup := Vec3{0, 1, 0}
	dist_to_focus := 10.0
	image_height := int(float64(image_width) / aspect_ratio)

	cam := new_camera(lookfrom, lookat, vup, vfov, aspect_ratio, aperture, dist_to_focus, 0.0, 1.0)

	img := image.NewRGBA(image.Rect(0, 0, image_width, image_height))

	remaining := int64(image_height)
	var wg sync.WaitGroup

	for j := image_height - 1; j >= 0; j-- {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			for i := 0; i < image_width; i++ {
				pixel_color := Vec3{0, 0, 0}
				for s := 0; s < samples_per_pixel; s++ {
					u := (float64(i) + rand.Float64()) / float64(image_width-1)
					v := (float64(j) + rand.Float64()) / float64(image_height-1)
					r := cam.get_ray(u, v)
					pixel_color = vec_add(pixel_color, ray_color(r, world, background, max_depth))
				}

				scale := 1.0 / float64(samples_per_pixel)
				cr := math.Sqrt(clamp(pixel_color.X*scale, 0, 0.999))
				cg := math.Sqrt(clamp(pixel_color.Y*scale, 0, 0.999))
				cb := math.Sqrt(clamp(pixel_color.Z*scale, 0, 0.999))

				img.Set(i, image_height-1-j, color.RGBA{
					uint8(256 * cr),
					uint8(256 * cg),
					uint8(256 * cb),
					255,
				})
			}
			r := atomic.AddInt64(&remaining, -1)
			fmt.Printf("\rScanlines remaining: %d   ", r)
		}(j)
	}

	wg.Wait()
	fmt.Println("\nDone.")

	out, err := os.Create("./output.png")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := png.Encode(out, img); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := out.Close(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
