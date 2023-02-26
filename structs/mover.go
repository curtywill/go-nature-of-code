package structs

import (
	"nature-of-code/constants"

	"github.com/fogleman/gg"
)

type Ch1Mover struct {
	position     *Vector
	velocity     *Vector
	acceleration *Vector
	maxspeed     float64
}

type Ch2Mover struct {
	position     *Vector
	velocity     *Vector
	acceleration *Vector
	mass         float64
}

func NewCh1Mover(x, y float64) *Ch1Mover {
	position := NewVector(x, y)
	velocity := ZeroVector()
	acceleration := ZeroVector()
	return &Ch1Mover{position, velocity, acceleration, 4}
}

func NewCh2Mover(x, y float64) *Ch2Mover {
	position := NewVector(x, y)
	velocity := ZeroVector()
	acceleration := ZeroVector()
	return &Ch2Mover{position, velocity, acceleration, 1}
}

func (m *Ch2Mover) SetMass(mass float64) {
	m.mass = mass
}

// force = mass * acceleration
func (m *Ch2Mover) ApplyForce(force Vector) {
	force.Div(m.mass)
	m.acceleration.VAdd(force)
}

// (x, y) is point of attraction
func (m *Ch1Mover) Update(x, y float64) {
	point := NewVector(x, y)
	dir := Sub(*point, *m.position)

	dir.Normalize()
	dir.Scale(0.5)

	m.acceleration = dir
	m.velocity.VAdd(*m.acceleration)
	m.velocity.Limit(float64(m.maxspeed))
	m.position.VAdd(*m.velocity)
}

func (m *Ch2Mover) Update() {
	m.velocity.VAdd(*m.acceleration)
	m.velocity.Limit(1)
	m.position.VAdd(*m.velocity)
	m.acceleration.Scale(0)
}

func (m *Ch2Mover) Bounce() {
	if !(m.position.Y-4 < 0) {
		return
	}
	m.velocity.Y *= -1
}

// for exercise 2.3
func (m *Ch2Mover) Constrain() {
	scale := 3000.0
	top := NewVector(0, (constants.H-m.position.Y)/scale) // distance from the bottom
	btm := NewVector(0, -m.position.Y/scale)
	left := NewVector((constants.W-m.position.X)/scale, 0)
	right := NewVector(-m.position.X/scale, 0)

	m.ApplyForce(*top)
	m.ApplyForce(*btm)
	m.ApplyForce(*left)
	m.ApplyForce(*right)
}

func (m Ch1Mover) Show(ctx *gg.Context) {
	ctx.SetRGB(1, 1, 1)
	ctx.DrawCircle(m.position.X, m.position.Y, 8)
	ctx.Fill()
}

func (m Ch2Mover) Show(ctx *gg.Context) {
	ctx.SetRGB(1, 1, 1)
	ctx.DrawCircle(m.position.X, m.position.Y, m.mass)
	ctx.Fill()
}
