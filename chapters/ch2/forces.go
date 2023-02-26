package ch2

// TODO: come back to this after 1D noise
import (
	"fmt"
	"math/rand"
	"nature-of-code/constants"
	"nature-of-code/structs"

	"github.com/fogleman/gg"
)

var m *structs.Ch2Mover
var dc *gg.Context

type AirBalloonSketch struct {
	helium *structs.Vector
	wind   *structs.Vector
}

// air balloon exercise
func (a *AirBalloonSketch) Setup() {
	m = structs.NewCh2Mover(constants.W/2, constants.H-50)
	dc = gg.NewContext(constants.W, constants.H)
	a.helium = structs.NewVector(0, -0.1)
	a.wind = structs.NewVector(0.08, 0)
}

func (a *AirBalloonSketch) loop(frameCount int) {
	if frameCount%(5*constants.FPS) == 0 {
		newWind := rand.Float64()*2 - 1
		a.wind = structs.NewVector(newWind, 0)
	}
	dc.SetRGB(0, 0, 0)
	dc.Clear()
	m.ApplyForce(*a.helium)
	m.ApplyForce(*a.wind)
	m.Update()
	m.Bounce()
	m.Show(dc)
	dc.SavePNG(fmt.Sprintf("frames/frame%03d.png", frameCount))
}

func (a *AirBalloonSketch) Draw() {
	for i := 0; i < constants.FPS*constants.SECS; i++ {
		a.loop(i)
	}
}

// exercise 2.3
type EdgeForceSketch struct {
	thrust *structs.Vector
	wind   *structs.Vector
	movers [100]*structs.Ch2Mover
}

func (e *EdgeForceSketch) Setup() {
	dc = gg.NewContext(constants.W, constants.H)
	e.thrust = structs.NewVector(0, -0.05)
	for i := 0; i < 100; i++ {
		e.movers[i] = structs.NewCh2Mover(rand.Float64()*constants.W, rand.Float64()*constants.H)
		e.movers[i].SetMass(rand.Float64() * 10)
	}
}

func (e *EdgeForceSketch) loop(frameCount int) {
	dc.SetRGB(0, 0, 0)
	dc.Clear()
	if frameCount%(2*constants.FPS) == 0 {
		newWind := rand.Float64()*2 - 1
		e.wind = structs.NewVector(newWind, 0)
	}
	for i := 0; i < 100; i++ {
		e.movers[i].ApplyForce(*e.thrust)
		e.movers[i].ApplyForce(*e.wind)
		e.movers[i].Constrain()
		e.movers[i].Update()
		e.movers[i].Show(dc)
	}
	dc.SavePNG(fmt.Sprintf("frames/frame%03d.png", frameCount))
}

func (e *EdgeForceSketch) Draw() {
	for i := 0; i < constants.FPS*constants.SECS; i++ {
		e.loop(i)
	}
}
