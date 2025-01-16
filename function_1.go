package main

import "fmt"

// 1. func {name of function} ({parameter}) {return type}  {logic of function}
func add(a int, b int) int {
	return a + b
}

// 2.multiple consecutive parameters
func add2(a, b, c, d int) int {
	return a + b + c + d
}

// 3. multiple return values from a function
func myName() (string, string) {
	return "Abhishek", "Singh"
}

func main() {
	res := add(5, 6)
	fmt.Println(res) //11

	res2 := add2(1, 4, 5, 7)
	fmt.Println(res2) //17

	firstName, _ := myName()
	fmt.Println(firstName) //Abhishek

	_, LastName := myName()
	fmt.Println(LastName)//Singh
}
