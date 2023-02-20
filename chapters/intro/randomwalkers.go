package intro

import (
	"fmt"
	"math/rand"
	"nature-of-code/processingutils"
	"nature-of-code/structs"
	"time"

	"github.com/fogleman/gg"
)

const (
	w = 200
	h = 200
)

func BiasedWalker() {
	walker := structs.NewVector(w/2, h/2)
	dc := gg.NewContext(w, h)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 10000; i++ {
		dc.DrawCircle(walker.X, walker.Y, 1)
		dc.SetRGB(0, 0, 0)
		dc.Fill()
		// dc.SavePNG(fmt.Sprintf("out/output%v.png", i))
		r1 := random.Float64()
		x, y := 0.0, 0.0
		if r1 < 0.4 {
			x = -1
		} else {
			x = 1
		}
		r2 := random.Float64()
		if r2 < 0.4 {
			y = -1
		} else {
			y = 1
		}
		walker.Add(x, y)
	}
	dc.SavePNG("out/BiasedWalker.png")
}

func GaussianWalker() {
	walker := structs.NewVector(w/2, h/2)
	dc := gg.NewContext(w, h)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	generator := processingutils.NewNormalGenerator()
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 150; i++ {
		dc.DrawCircle(walker.X, walker.Y, 5)
		dc.SetRGB(0, 0, 0)
		dc.Fill()
		dc.SavePNG(fmt.Sprintf("out/frame%03d.png", i))
		dc.SetRGB(1, 1, 1)
		dc.Clear()
		x, y := 0.0, 0.0
		num := generator.Next()
		stepsize := 10.0 * num
		r := int(random.Float64() * 2)
		if r == 0 {
			x = stepsize
		} else {
			y = stepsize
		}
		walker.Add(x, y)
	}
}

func NoiseWalker() {
	walker := structs.NewVector(w/2, h/2)
	dc := gg.NewContext(w, h)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	x, y := 0.0, 0.0
	perlin := processingutils.NewPerlinNoiseGenerator()
	for i := 0; i < 150; i++ {
		dc.DrawCircle(walker.X, walker.Y, 5)
		dc.SetRGB(0, 0, 0)
		dc.Fill()
		dc.SavePNG(fmt.Sprintf("out/frame%03d.png", i))
		dc.SetRGB(1, 1, 1)
		dc.Clear()
		x += 0.01
		y += 0.01
		perlin.Noise(x, y)
		walker.Add(x, y)
	}
}
