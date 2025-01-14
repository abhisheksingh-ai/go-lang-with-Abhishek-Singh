package main

import (
	"fmt"
	"reflect"
)

func main() {
	//creating an empty slice array
	arr := []int{}
	fmt.Println(arr, "len:", len(arr)) //Output: [] len: 0

	//appending in slices
	arr = append(arr, 2, 4, 5)
	fmt.Println(arr, "len:", len(arr))
	//Output: [2 4 5] len: 3

	//slices provide the reference to the array
	newArr := [7]int{0, 1, 2, 3, 4, 5, 6} // this static array with size 7
	slicedArr := newArr[1:6]
	fmt.Println("slicedArr:", slicedArr)
	//Output: [1 2 3 4 5], index 6 is excluded
	//changing in sliceArr will change in newArr beacause it is reffering to it
	slicedArr[0] = 100
	fmt.Println("newArr:", newArr)

	// ["","",""] Creating array of sting dynamically
	s := []string{}
	s = append(s, "Hello")
	s = append(s, "world")
	fmt.Println("s:", s, ", s[0]:", s[0])
	//output: s: [Hello world] , s[0]: Hello

	//Cheking the data type using: reflect
	fmt.Println("Type of data:", reflect.TypeOf(s))
	//Output: Type of data: []string

	fmt.Println("Type of data:", reflect.TypeOf(s[0]))
	//Output: Type of data: string

	fmt.Println(s[0][3])
	//Output: 108 (ASCII value of 'l')

	fmt.Println(string(s[0][3]))
	//Output: l

}
