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

func RandomUnitVector() *Vector {
	// r := rand.New(rand.NewSource(time.Now().UnixNano()))
	v := &Vector{rand.Float64()*2 - 1, rand.Float64()*2 - 1}
	v.Normalize()
	return v
}

func (v *Vector) Add(x, y float64) {
	v.X += x
	v.Y += y
}

func (v *Vector) Dot(o Vector) float64 {
	return v.X*o.X + v.Y*o.Y
}

func (v *Vector) Normalize() {
	mag := v.Mag()
	v.X /= mag
	v.Y /= mag
}

func (v Vector) Mag() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
