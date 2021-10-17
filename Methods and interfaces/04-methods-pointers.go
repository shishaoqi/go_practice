package main

import (
	"fmt"
	"math"
)

type Vertex_m3 struct {
	X, Y float64
}

func (v Vertex_m3) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

// You can declare methods with pointer revievers.
// Method with pointer recervers can modify the value to which the receiver pointer.
func (v *Vertex_m3)Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex_m3{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
