package main

import (
	"fmt"
	"runtime"
)

// A switch statement is a shorte way to write a sequence of if - else statements. It runs the first case whose value is
// equal to the condition expression
func main() {
	fmt.Println("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}
