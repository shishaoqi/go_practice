package main

import (
	"fmt"
	"math/rand"
)

// Every Go program is made up of packages.
// Programs start running in pacgage main.

func main() {
	fmt.Println("My favorite number is", rand.Intn(100))
}
