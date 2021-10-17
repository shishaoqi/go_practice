package main

import (
	"fmt"
	"math"
)

// Go does not have classes. However, you can define methods on types.
type Vertex_method struct {
	X, Y float64
}

func (v Vertex_method) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func main() {
	v := Vertex_method{3, 4}
	fmt.Println(v.Abs())
}
