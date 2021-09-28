package main

import "fmt"

// Go has only one looping construct, the for loop
/**
	The basic for loop has three components separeted by semicolons:
	1.
	2.
	3.

 */

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
