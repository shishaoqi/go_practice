package main

import "fmt"

// Struct fields can be accessed through a struct pointer
type Pvertex struct {
	X int
	Y int
}

func main() {
	v := Pvertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}
