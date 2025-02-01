package main

import "fmt"

func main() {
	// Using explicit declaration with type
	var x int = 5
	var y string = "Go Programming"
	// Using short declaration
	z := 10.5
	// Multiple variables
	var a, b, c int = 1, 2, 3
	var d, e = 10, "Hello"
	// Constant declaration
	const pi = 3.14159
	// Printing the variables
	fmt.Println(x, y, a, b, c, d, e, pi, z)
}
