package main

import (
	"image"
	"image/png"
	"math"
	"os"
)

var _ texture = new_image_texture("")

type image_texture struct {
	data          image.Image
	width, height int
}

func new_image_texture(filename string) image_texture {
	f, err := os.Open(filename)
	if err != nil {
		return image_texture{}
	}
	defer f.Close() //nolint:errcheck

	img, err := png.Decode(f)
	if err != nil {
		return image_texture{}
	}

	bounds := img.Bounds()
	return image_texture{
		data:   img,
		width:  bounds.Dx(),
		height: bounds.Dy(),
	}
}

func (t image_texture) value(u, v float64, p Vec3) Vec3 {
	if t.data == nil {
		return Vec3{0, 1, 1}
	}

	u = math.Max(0, math.Min(u, 1))
	v = 1.0 - math.Max(0, math.Min(v, 1))

	i := int(u * float64(t.width))
	j := int(v * float64(t.height))

	if i >= t.width {
		i = t.width - 1
	}
	if j >= t.height {
		j = t.height - 1
	}

	color_scale := 1.0 / 255.0
	r, g, b, _ := t.data.At(i+t.data.Bounds().Min.X, j+t.data.Bounds().Min.Y).RGBA()

	return Vec3{
		float64(r>>8) * color_scale,
		float64(g>>8) * color_scale,
		float64(b>>8) * color_scale,
	}
}
