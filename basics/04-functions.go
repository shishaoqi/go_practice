package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

// A function can take zero or more arguments.
func main() {
	fmt.Println(add(8, 90))
}
