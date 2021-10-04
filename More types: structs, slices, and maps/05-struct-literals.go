package main

import "fmt"

type Vertex_5 struct {
	X, Y int
}

// A struct literal denotes a newly allocated struct value by listing the values of its fields
var (
	v1 = Vertex_5{1, 2} // has type Vertex
	v2 = Vertex_5{X: 1} // Y:0 is implicit
	v3 =Vertex_5{} // X: 0 and Y:0
	p = &Vertex_5{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3)
}