package processingutils

import (
	"math/rand"
	"nature-of-code/structs"
)

const (
	TABLE_SIZE = 256
	MASK       = TABLE_SIZE - 1
)

type PerlinNoise struct {
	gradients    [TABLE_SIZE]structs.Vector
	permutations [TABLE_SIZE * 2]int
	octaves      int
	falloff      float64
}

func NewPerlinNoiseGenerator() PerlinNoise {
	// r := rand.New(rand.NewSource(time.Now().UnixNano()))
	gradients := [TABLE_SIZE]structs.Vector{}
	permutations := [TABLE_SIZE * 2]int{}
	for i := 0; i < TABLE_SIZE; i++ {
		gradients[i] = *structs.RandomUnitVector()
		permutations[i] = rand.Intn(TABLE_SIZE)
		permutations[i+TABLE_SIZE] = permutations[i]
	}
	return PerlinNoise{gradients, permutations, 4, 0.5}
}

func (p PerlinNoise) gradientDotOffset(cx int, cy int, px float64, py float64) float64 {
	gradient := p.gradients[p.permutations[p.permutations[cx]+cy]]
	offset := structs.NewVector(px-float64(cx), py-float64(cy))
	return gradient.Dot(*offset)
}

func lerp(w, a0, a1 float64) float64 {
	return (a1-a0)*((w*(w*6.0-15.0)+10.0)*w*w*w) + a0
}

func (p *PerlinNoise) NoiseDetail(octaves int, falloff float64) {
	p.octaves = octaves
	p.falloff = falloff
}

func (p PerlinNoise) Noise(x float64, y float64) float64 {
	sum := 0.0
	freq := 1.0
	amp := 1.0
	totalAmp := 0.0
	for i := 0; i < p.octaves; i++ {
		sum += p.noise(x*freq, y*freq) * amp
		totalAmp += amp
		amp *= p.falloff
		freq *= 2
	}
	return sum / totalAmp
}

func (p PerlinNoise) noise(x, y float64) float64 {
	x1 := int(x) & MASK
	y1 := int(y) & MASK
	x2 := (x1 + 1) & MASK
	y2 := (y1 + 1) & MASK

	u, v := x-float64(x1), y-float64(y1)

	d1 := p.gradientDotOffset(x1, y1, x, y)
	d2 := p.gradientDotOffset(x2, y1, x, y)
	i1 := lerp(u, d1, d2)

	d3 := p.gradientDotOffset(x1, y2, x, y)
	d4 := p.gradientDotOffset(x2, y2, x, y)
	i2 := lerp(u, d3, d4)

	return lerp(v, i1, i2)*0.5 + 0.5
}
