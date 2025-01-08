package main

import "fmt"

func main() {
	var a [4]int // bydefault value 0
	fmt.Println("emp", a)

	a[3] = 100
	fmt.Println("emp", a) // [0, 0 , 0, 100]

	length := len(a) // calculating the length of array
	fmt.Println("length:", length)

	// declaration and initalization in single line
	b := [3]int{100, 101, 102}
	fmt.Println(b) //[100 101 102]

	c := [...]int{200, 201, 202}
	fmt.Println(c)
	fmt.Println(len(c)) // 3

	d := [...]int{400, 3: 500, 600}
	fmt.Println(d)

	//2D vector
	var twoD [2][3]int
	//taking input
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	//printing
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(twoD[i][j])
		}
		fmt.Print("\n")
	}
	//012
	//123
}
