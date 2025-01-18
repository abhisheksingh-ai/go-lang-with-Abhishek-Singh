package main
import "fmt"

func pointerExample(i *int) {
	// i pass here by reference, mean its pointing to the main func i
	*i = 0
}
func normalFunction(i int) {
	//i is pass here :- pass by value, as a diferent copy
	i = 2
}
func swap(x *int, y *int) {
	temp := *x
	*x = *y
	*y = temp
}
func main() {
	//Example 1
	i := 1
	normalFunction(i)
	fmt.Println("i after normal function: ", i) //1
	pointerExample(&i)
	fmt.Println("i after pointer example call", i) //0

	//Example 2
	//Swapping Values Using Pointers
	a := 4
	b := 10
	swap(&a, &b)
	fmt.Println("After swaping a: ", a, " and b: ", b)
	//After swaping a:  10  and b:  4
}
