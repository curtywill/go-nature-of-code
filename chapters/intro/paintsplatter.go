package intro

import (
	"math"
	"nature-of-code/processingutils"

	"github.com/fogleman/gg"
)

func NormalPaintSplatter() {
	dc := gg.NewContext(w, h)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	generator := processingutils.NewNormalGenerator()
	xsd := 25.0
	xmean := 300.0
	ysd := 25.0
	ymean := 200.0
	for i := 0; i < 500; i++ {
		num := generator.Next()
		x := xsd*num + xmean
		num = generator.Next()
		y := ysd*num + ymean
		dc.DrawCircle(x, y, 1)
		dc.SetRGBA(0, 0, 0, 0.5)
		dc.Fill()
	}
	dc.SavePNG("out/NormalPaintSplatter.png")
}

func NormalColorPalette() {
	dc := gg.NewContext(w, h)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	generator := processingutils.NewNormalGenerator()
	xsd := 25.0
	xmean := 300.0
	ysd := 25.0
	ymean := 200.0
	rsd := 0.25
	rmean := 0.5
	for i := 0; i < 500; i++ {
		num := generator.Next()
		x := xsd*num + xmean
		num = generator.Next()
		y := ysd*num + ymean
		dc.DrawCircle(x, y, 5)
		num = generator.Next()
		r := rsd*num + rmean
		dc.SetRGB(math.Abs(r), 0, 1)
		dc.Fill()
	}
	dc.SavePNG("out/NormalColorPalette.png")
}
