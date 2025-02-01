package main

import "fmt"

func oddSum(nums ...int) int {
	sum := 0;
	for _, num := range nums {
		sum += num;
	}
	return sum;
}

func fullname(arr ...string) string {
	fullName := "";
	for _, str := range arr {
		fullName += str;
	}
	return fullName;
}

func main() {
	//Variadic function: can take any number of arbitrary argument
	//and allow to use as a collection of slice in its scope
	res := oddSum(1, 3, 5, 7, 9)
	fmt.Println(res)//25

	res2 := fullname("Abhishek", "Singh")
	fmt.Println(res2)//AbhishekSingh
}
