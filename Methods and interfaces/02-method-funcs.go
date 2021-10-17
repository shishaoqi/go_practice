package main

import (
	"fmt"
	"math"
)

type Vertex_m2 struct {
	X, Y float64
}

// Remeber: a method is just a function with a receiver argument
func Abs(v Vertex_m2) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex_m2{3, 4}
	fmt.Println(Abs(v))
}