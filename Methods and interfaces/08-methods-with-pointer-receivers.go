package main

import (
	"fmt"
	"math"
)

type Vertex_m8 struct {
	X, Y float64
}

// There are two reasons to use a pointer reveiver.
// The first is so that the method can modify the value that its reveiver pointer to.
// The second is to avoid copying the value on each method call.
//This can be more efficient if the reveiver is a large struct.

func (v *Vertex_m8) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex_m8) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex_m8{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}