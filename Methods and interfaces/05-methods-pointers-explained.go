package main

import (
	"fmt"
	"math"
)

type Vertex_m5 struct {
	X, Y float64
}

func Abs(v Vertex_m5) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex_m5, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex_m5{3, 4}
	Scale(&v, 10)
	fmt.Println(Abs(v))
}
