package main
import (
	"fmt"
	"slices"
)
func main() {
	//Step 1: Appending slices
	slicesA := []int{3, 5, 7, 3}
	slicesB := []int{1, 3, 5, 7, 9}
	slicesC := []int{}
	slicesC = append(slicesC, slicesA...)
	slicesC = append(slicesC, slicesB...)
	// fmt.Println(slicesC) //Output: [3 5 7 3 1 3 5 7 9]
	//2. finding duplicate using MAP
	//creating a map to track the frequency of each element
	//key of the map will be uniqu slices elements
	mpp := make(map[int]int)
	for i := 0; i < len(slicesC); i++ {
		key := slicesC[i]
		mpp[key]++
	}
	//fmt.Println("mpp", mpp) //Output:  mpp map[1:1 3:3 5:2 7:2 9:1]
	uniqueSlice := []int{}
	for key := range mpp {
		uniqueSlice = append(uniqueSlice, key)
	}
	//fmt.Println(uniqueSlice) //[7 1 9 3 5]
	//sorting of this slice
	slices.Sort(uniqueSlice)
	fmt.Println(uniqueSlice)
	//Result: [1 3 5 7 9]
}
