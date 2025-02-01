package main
import "fmt"
// closure
func counter() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
// Recursion
func fact(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	ans := fact(n-1) * n
	return ans
}

func main() {
	//Closure:-A function value that references variables from its surrounding scope
	count := counter()
	fmt.Println(count())    //1
	fmt.Println(count())    //2
	fmt.Println(count())    //3
	fmt.Println(count())    //4

	newCount := counter()   //New count is reffereng to 1 here
	fmt.Println(newCount()) //1
	
	res := fact(7)
	fmt.Println(res) //5070
}
