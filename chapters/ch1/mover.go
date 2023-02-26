package ch1

import (
	"fmt"
	"math/rand"
	"nature-of-code/constants"
	"nature-of-code/structs"

	"github.com/fogleman/gg"
)

// created mover0 through mover2 gifs
var m *structs.Ch1Mover
var dc *gg.Context
var x, y float64

func Setup() {
	m = structs.NewCh1Mover(constants.W/2, constants.H/2)
	dc = gg.NewContext(constants.W, constants.H)
}

func loop(frameCount int) {
	if frameCount%(3*constants.FPS) == 0 {
		x, y = rand.Float64()*constants.W, rand.Float64()*constants.H
	}
	dc.SetRGB(0, 0, 0)
	dc.Clear()
	m.Update(x, y)
	m.Show(dc)
	dc.SavePNG(fmt.Sprintf("frames/frame%03d.png", frameCount))
}

func Draw() {
	for i := 0; i < constants.FPS*constants.SECS; i++ {
		loop(i)
	}
}
