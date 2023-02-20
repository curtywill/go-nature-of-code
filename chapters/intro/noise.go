package intro

import (
	"fmt"
	"nature-of-code/processingutils"

	"github.com/fogleman/gg"
)

func NoiseExample() {
	dc := gg.NewContext(w, h)
	dc.SetRGB(0, 0, 0)
	dc.Clear()
	perlin := processingutils.NewPerlinNoiseGenerator()
	//(&perlin).NoiseDetail(2, 0.75)

	inc := 0.04
	for i := 0; i < 150; i++ {
		dx, dy := float64(i)*inc, float64(i)*inc
		for x := 0; x < h; x++ {
			dy = float64(i) * inc
			for y := 0; y < w; y++ {
				r := int(perlin.Noise(dx, dy) * 255)
				dc.SetRGB255(r, r, r)
				dc.SetPixel(x, y)
				dy += inc
			}
			dx += inc
		}
		dc.SavePNG(fmt.Sprintf("frames/frame%03d.png", i))
	}
	dc.SavePNG("out/noise.png")
}
