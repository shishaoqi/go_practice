package main

import "fmt"

// Go function may be closuers
// A closure is a function value that references variable from outside its body
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-1*i),
		)
	}
}
