package structs

type vector struct {
	X float64
	Y float64
}

func NewVector(x, y float64) *vector {
	return &vector{x, y}
}

func (v *vector) Add(x, y float64) {
	v.X += x
	v.Y += y
}
