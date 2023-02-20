package structs

import (
	"math"
	"math/rand"
)

type Vector struct {
	X float64
	Y float64
}

func NewVector(x, y float64) *Vector {
	return &Vector{x, y}
}

func ZeroVector() *Vector {
	return &Vector{0, 0}
}

func RandomUnitVector() *Vector {
	v := &Vector{rand.Float64()*2 - 1, rand.Float64()*2 - 1}
	v.Normalize()
	return v
}

func (v *Vector) Add(x, y float64) {
	v.X, v.Y = v.X+x, v.Y+y
}

func (v *Vector) VAdd(o Vector) {
	v.X, v.Y = v.X+o.X, v.Y+o.Y
}

func Sub(v1, v2 Vector) *Vector {
	return NewVector(v1.X-v2.X, v1.Y-v2.Y)
}

func (v *Vector) Scale(num float64) {
	v.X, v.Y = v.X*num, v.Y*num
}

func (v *Vector) Div(num float64) {
	v.X, v.Y = v.X/num, v.Y/num
}

func (v *Vector) Limit(max float64) {
	mag := v.Mag()
	if mag > max {
		v.Div(mag / max)
	}
}

func (v *Vector) Dot(o Vector) float64 {
	return v.X*o.X + v.Y*o.Y
}

func (v *Vector) Normalize() {
	mag := v.Mag()
	v.X, v.Y = v.X/mag, v.Y/mag
}

func (v Vector) Mag() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
