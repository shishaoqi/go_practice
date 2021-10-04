package main

import "fmt"

// mutate  [mju:`teit] v. (使)变异；【生】（使）突变

// Insert or update an element in map m:
// m[key] = elem

// Retrieve an element:
// elem = m[key]

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	// Delete an element
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	// Test That a key is present with a two-value assignment
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Presend?", ok)
}
