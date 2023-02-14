package intro

import (
	"nature-of-code/processingutils"

	"github.com/fogleman/gg"
)

func noiseExample() {
	dc := gg.NewContext(w, h)
	dc.SetRGB(0, 0, 0)
	dc.Clear()
	perlin := processingutils.NewPerlinNoiseGenerator()
	(&perlin).NoiseDetail(4, 0.75)
	dx, dy := 0.0, 0.0
	for x := 0; x < w; x++ {
		dx = 0.0
		for y := 0; y < h; y++ {
			r := int(perlin.Noise(dx, dy) * 255)
			dc.SetRGB255(r, r, r)
			dc.SetPixel(x, y)
			dx += 0.05
		}
		dy += 0.05
	}
	dc.SavePNG("out/noise.png")
}
