package processingutils

import (
	"math"
	"math/rand"
	"time"
)

// Recreates the random Gaussian/normal generator used in Processing
type normalGenerator struct {
	uniformGenerator *rand.Rand
	nextNormal       float64
	hasNextNormal    bool
}

func NewNormalGenerator() *normalGenerator {
	return &normalGenerator{rand.New(rand.NewSource(time.Now().UnixNano())), 0, false}
}

// https://en.wikipedia.org/wiki/Normal_distribution#Generating_values_from_normal_distribution
func (n *normalGenerator) Next() float64 {
	if n.hasNextNormal {
		n.hasNextNormal = false
		return n.nextNormal
	}
	var u, v float64
	s := 1.0
	for s >= 1 {
		u = n.uniformGenerator.Float64()*2 - 1
		v = n.uniformGenerator.Float64()*2 - 1
		s = u*u + v*v
	}
	a := math.Sqrt(-2 * math.Log(s) / s)
	n.nextNormal = a * v
	n.hasNextNormal = true
	return a * u
}
