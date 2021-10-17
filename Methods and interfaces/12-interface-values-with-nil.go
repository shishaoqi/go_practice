package main

import "fmt"

// If the concrete value inside the interface itself is nil, the method will be called with a nil receiver.

// In some languages thi would trigger a null pointer exception, but in Go it is common to write methods that gracefully
// handle being called with a nil reveiver (as with the method Min this example)

// Note that an interface value that holds a nil concrete value is itself non-nil.
type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

