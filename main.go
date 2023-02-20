package main

import (
	"fmt"
	"math/rand"
	"nature-of-code/chapters/ch1"
	"nature-of-code/constants"
	"time"

	"github.com/fogleman/gg"
)

var m *ch1.Mover
var dc *gg.Context

func main() {
	rand.Seed(time.Now().UnixNano())
	setup()
	draw()
}

func setup() {
	m = ch1.NewMover()
	dc = gg.NewContext(constants.W, constants.H)
}

func draw() {
	x, y := 500.0, 300.0
	for i := 0; i < constants.FPS*constants.SECS; i++ {
		if i%(3*constants.FPS) == 0 {
			x, y = rand.Float64()*constants.W, rand.Float64()*constants.H
		}
		dc.SetRGB(0, 0, 0)
		dc.Clear()
		m.Update(x, y)
		m.Show(dc)
		dc.SavePNG(fmt.Sprintf("frames/frame%03d.png", i))
	}
}
