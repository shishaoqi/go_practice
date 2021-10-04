package main

import "fmt"

// It is common to append new elements to a slice, and so Go provides a build-in append function.
// The documentation of the build-in package describe append:
// func append(s []T, vs ...T) []T
func main() {
	var s []int
	printSliceAppend(s)

	// append works on nil slices.
	s = append(s, 0)
	printSliceAppend(s)

	// The slice grows as needed
	s = append(s, 1)
	printSliceAppend(s)

	// We can add more than one element at a time
	s = append(s, 2, 3, 4)
	printSliceAppend(s)
}

func printSliceAppend(s []int) {
	fmt.Printf("len=%d cap =%d %v\n", len(s), cap(s), s)
}
