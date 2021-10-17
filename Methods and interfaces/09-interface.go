package main

import (
	"fmt"
	"math"
)

// An interface type is defined as a set of method signatures
// A value of interface type can hold any value that implements those methods
type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat_m9(-math.Sqrt2)
	v := Vertex_m9{3, 4}

	a = f // a MyFloat implements Abser
	fmt.Println(a.Abs())

	a = &v // a * Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does Not implement Abser
	//a = v


	fmt.Println(a.Abs())
}

type MyFloat_m9 float64

func (f MyFloat_m9) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex_m9 struct {
	X, Y float64
}

func (v *Vertex_m9) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}