package main

// main package is used to run the setup and draw functions of the current sketch
import (
	"math/rand"
	"nature-of-code/chapters/ch2"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	sketch := ch2.EdgeForceSketch{}
	sketch.Setup()
	sketch.Draw()
}
