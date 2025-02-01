/*
🚀 What I Learned Today: Enums in Go!

Enums are a game-changer when you need a predefined set of values in your code! Instead of using random numbers or strings, enums provide:
✅ Readability: Replace confusing magic numbers like 2 with meaningful names like Shipped.
✅ Type Safety: Prevent invalid values from sneaking into your functions.
✅ Maintainability: Easily add new options without breaking existing logic.

👨‍💻 Problem I Solved:
I had to manage order statuses for an e-commerce system (like Pending, Shipped, Delivered). Without enums, it’s easy to lose track of what each value represents. By using enums with Go's iota, my code became clean, easy to read, and bug-free!

💡 Pro Tip: Combine enums with switch to handle logic for each value effortlessly.

📚 Keep learning, keep growing! 🌱

*/


package main

import "fmt"

type orderStatus int

const (
	Pending orderStatus = iota
	Processing
	Shipped
	Delivered
	Cancelled
)

func getInfo(num orderStatus) {
	switch num {
	case Pending:
		fmt.Println("Staus is Pending")
	case Processing:
		fmt.Println("Staus is Processing")
	case Shipped:
		fmt.Println("Staus is Shipped")
	case Delivered:
		fmt.Println("Staus is Delivered")
	case Cancelled:
		fmt.Println("Staus is Cancelled")
	}
}

func main() {
	status := Shipped
	getInfo(status) //Staus is Shipped
}
