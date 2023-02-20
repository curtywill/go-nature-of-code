package ch1

import (
	"nature-of-code/constants"
	"nature-of-code/structs"

	"github.com/fogleman/gg"
)

type Mover struct {
	position     *structs.Vector
	velocity     *structs.Vector
	acceleration *structs.Vector
	maxspeed     int
}

func NewMover() *Mover {
	position := structs.NewVector(constants.W/2, constants.H/2)
	velocity := structs.ZeroVector()
	return &Mover{position: position, velocity: velocity, maxspeed: 4}
}

// (x, y) is point of attraction
func (m *Mover) Update(x, y float64) {
	point := structs.NewVector(x, y)
	dir := structs.Sub(*point, *m.position)

	dir.Normalize()
	dir.Scale(0.5)

	m.acceleration = dir
	m.velocity.VAdd(*m.acceleration)
	m.velocity.Limit(float64(m.maxspeed))
	m.position.VAdd(*m.velocity)
}

func (m Mover) Show(ctx *gg.Context) {
	ctx.SetRGB(1, 1, 1)
	ctx.DrawCircle(m.position.X, m.position.Y, 8)
	ctx.Fill()
}
