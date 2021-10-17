package main

import "fmt"

type Vertex_m6 struct {
	X, Y float64
}

func (v *Vertex_m6) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex_m6, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex_m6{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex_m6{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}

