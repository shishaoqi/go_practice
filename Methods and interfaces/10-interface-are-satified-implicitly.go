package main

import "fmt"

type I interface {
	M()
}

// A type implements an interface by implementing its methods.
// There are no explicit declaration of intent, no "implements" keyword.
type T struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello world"}
	i.M()
}
