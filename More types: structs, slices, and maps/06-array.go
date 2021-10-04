package main

import "fmt"

// The type [n]T is an array of n values of type T
// The expression
// var a [10]int // [n]T

func main() {
	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
