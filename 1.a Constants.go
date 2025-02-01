package main

import (
	"fmt"
	"math"
)

func main() {
	//1. max value in go lang
	fmt.Println("max int", math.MaxInt)
	//max int 9223372036854775807
	fmt.Println(math.MaxInt32)
	// 2147483647

	//2. uint8 unsigned int 8 bit (0-255) range
	fmt.Println(math.MaxUint8)
	//255

	//3.float values
	var ratio float64 = 34 / 7688 // int / int ---> int result
	fmt.Println(ratio)
	// 0

	var ratio2 float64 = float64(34) / float64(7688)
	fmt.Println(ratio2)
	//0.004422476586888657
}
