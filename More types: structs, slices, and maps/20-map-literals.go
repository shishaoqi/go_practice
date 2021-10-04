package main

import "fmt"

type Vertex_20 struct {
	Lat, Long float64
}

var m20 = map[string]Vertex_20 {
	"Bell Labs": Vertex_20 {
		40.48433, 74.39967,
	},
	"Gooble": Vertex_20{
		37.42202, -122.08408,
	},
}

// Map literals are like struct literals, but the keys are requied
func main() {
	fmt.Println(m20)
}
