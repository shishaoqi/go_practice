package main

import "fmt"

// Let's have some fun with functions

func fibonacci(i int) int {
	if i <= 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return fibonacci(i - 1) + fibonacci(i - 2)
}

func main() {
	//f := fibonacci(int)
	for i := 0; i < 10; i++ {
		fmt.Println(fibonacci(i))
	}
}
