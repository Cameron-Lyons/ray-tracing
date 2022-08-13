package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

func main() {

	const image_width int = 256
	const image_height int = 256

	image := image.NewRGBA(image.Rect(0, 0, image_width, image_height))

	for j := image_height - 1; j >= 0; j-- {
		for i := 0; i < image_width; i++ {

			r := float64(i) / float64(image_width-1)
			g := float64(j) / float64(image_height-1)
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
