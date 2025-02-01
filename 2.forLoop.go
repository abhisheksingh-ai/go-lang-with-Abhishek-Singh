package main

import "fmt"

func main() {
	//For loop and if condtions in go lang by example

	//eg 1: Prinitng even numbers from 1 to 10
	for i := range 10 {
		if (i % 2) == 0 {
			continue
		}
		fmt.Println(i) // 13579
	}

	//eg 2: printing stars
	// *****
	// *****
	// *****
	// *****
	// *****
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}

	//eg 3: In Go, the for loop is very flexible and can be used to mimic a while loop
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}
}
