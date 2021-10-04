package main

import "fmt"

type Vertex_19 struct {
	Lat, Long float64
}

var m map[string]Vertex_19

// A map maps keys to values
func main() {
	m = make(map[string]Vertex_19)
	m["Bell Labs"] = Vertex_19{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
