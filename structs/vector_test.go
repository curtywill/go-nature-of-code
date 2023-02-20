package structs

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestZeroVector(t *testing.T) {
	v := NewVector(0, 0)
	if v.X != 0 || v.Y != 0 {
		t.Fatal("Zero vector should be (0, 0)")
	}
}

func TestRandomUnitVector(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		v := RandomUnitVector()
		t.Log(v.Mag())
		if math.Round(v.Mag()) != 1.0 {
			t.Fatal("Normalized vector should have length of ~1")
		}
	}
}

func TestVectorLimit(t *testing.T) {
	v := NewVector(10, 20)
	v.Limit(10)
	t.Log(v.Mag())
}
