package main

import "fmt"

const Pi = 3.14

// Constants are declared like variables, but with the const keyword

func main() {
	const World = "hello world"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
