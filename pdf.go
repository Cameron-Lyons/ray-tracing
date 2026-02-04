package main

import (
	"math"
	"math/rand"
)

type scatter_record struct {
	specular_ray ray
	is_specular  bool
	attenuation  Vec3
	pdf_ptr      pdf
}

type pdf interface {
	value(direction Vec3) float64
	generate() Vec3
}

type pdf_hittable interface {
	pdf_value(o, v Vec3) float64
	random(o Vec3) Vec3
}

type cosine_pdf struct {
	uvw onb
}

func new_cosine_pdf(w Vec3) cosine_pdf {
	return cosine_pdf{new_onb(w)}
}

func (c cosine_pdf) value(direction Vec3) float64 {
	cosine := vec_dot(unit_vector(direction), c.uvw.w)
	if cosine <= 0 {
		return 0
	}
	return cosine / math.Pi
}

func (c cosine_pdf) generate() Vec3 {
	return c.uvw.local(random_cosine_direction())
}

type hittable_pdf struct {
	o   Vec3
	ptr pdf_hittable
}

func new_hittable_pdf(p pdf_hittable, origin Vec3) hittable_pdf {
	return hittable_pdf{origin, p}
}

func (h hittable_pdf) value(direction Vec3) float64 {
	return h.ptr.pdf_value(h.o, direction)
}

func (h hittable_pdf) generate() Vec3 {
	return h.ptr.random(h.o)
}

type mixture_pdf struct {
	p [2]pdf
}

func new_mixture_pdf(p0, p1 pdf) mixture_pdf {
	return mixture_pdf{[2]pdf{p0, p1}}
}

func (m mixture_pdf) value(direction Vec3) float64 {
	return 0.5*m.p[0].value(direction) + 0.5*m.p[1].value(direction)
}

func (m mixture_pdf) generate() Vec3 {
	if rand.Float64() < 0.5 {
		return m.p[0].generate()
	}
	return m.p[1].generate()
}
