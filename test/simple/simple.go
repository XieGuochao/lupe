package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func concatString(a, b string) string {
	return a + b
}

func main() {
	// Testing the add function
	sum := add(5, 3)
	fmt.Println("Sum:", sum)

	// Testing the concatString function
	result := concatString("Hello, ", "World!")
	fmt.Println("Concatenated String:", result)
}
