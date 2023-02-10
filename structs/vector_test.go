package structs

import "testing"

func TestZeroVector(t *testing.T) {
	v := NewVector(0, 0)
	if v.X != 0 || v.Y != 0 {
		t.Fatal("Zero vector should be (0, 0)")
	}
}
