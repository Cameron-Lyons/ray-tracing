package main

import (
	"math"
	"math/rand"
)

const point_count = 256

type perlin struct {
	ranvec []Vec3
	perm_x []int
	perm_y []int
	perm_z []int
}

func new_perlin() perlin {
	ranvec := make([]Vec3, point_count)
	for i := 0; i < point_count; i++ {
		ranvec[i] = unit_vector(random_vec3_range(-1, 1))
	}
	return perlin{
		ranvec: ranvec,
		perm_x: perlin_generate_perm(),
		perm_y: perlin_generate_perm(),
		perm_z: perlin_generate_perm(),
	}
}

func perlin_generate_perm() []int {
	p := make([]int, point_count)
	for i := 0; i < point_count; i++ {
		p[i] = i
	}
	for i := point_count - 1; i > 0; i-- {
		target := rand.Intn(i + 1)
		p[i], p[target] = p[target], p[i]
	}
	return p
}

func (p perlin) noise(point Vec3) float64 {
	u := point.X - math.Floor(point.X)
	v := point.Y - math.Floor(point.Y)
	w := point.Z - math.Floor(point.Z)

	u = u * u * (3 - 2*u)
	v = v * v * (3 - 2*v)
	w = w * w * (3 - 2*w)

	i := int(math.Floor(point.X))
	j := int(math.Floor(point.Y))
	k := int(math.Floor(point.Z))

	var c [2][2][2]Vec3
	for di := 0; di < 2; di++ {
		for dj := 0; dj < 2; dj++ {
			for dk := 0; dk < 2; dk++ {
				c[di][dj][dk] = p.ranvec[p.perm_x[(i+di)&255]^p.perm_y[(j+dj)&255]^p.perm_z[(k+dk)&255]]
			}
		}
	}

	return perlin_interp(c, u, v, w)
}

func perlin_interp(c [2][2][2]Vec3, u, v, w float64) float64 {
	uu := u * u * (3 - 2*u)
	vv := v * v * (3 - 2*v)
	ww := w * w * (3 - 2*w)
	accum := 0.0

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				fi := float64(i)
				fj := float64(j)
				fk := float64(k)
				weight := Vec3{u - fi, v - fj, w - fk}
				accum += (fi*uu + (1-fi)*(1-uu)) *
					(fj*vv + (1-fj)*(1-vv)) *
					(fk*ww + (1-fk)*(1-ww)) *
					vec_dot(c[i][j][k], weight)
			}
		}
	}
	return accum
}

func (p perlin) turbulence(point Vec3, depth int) float64 {
	accum := 0.0
	temp_p := point
	weight := 1.0

	for i := 0; i < depth; i++ {
		accum += weight * p.noise(temp_p)
		weight *= 0.5
		temp_p = vec_mul_scalar(temp_p, 2)
	}

	return math.Abs(accum)
}
