package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// You can declare a method on non-struct types, too.
func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs)

	f2 := MyFloat(3)
	fmt.Println(f2.Abs())
}
