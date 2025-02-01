// Create a Car Inventory Management System
// Define a Car struct with fields: Model, Brand, Year, and Price.
// Write functions to:
// Add a new car to the inventory (a slice of *Car).
// Update the price of a car given its model and brand.
// Find the most expensive car in the inventory.

package main
import "fmt"

type Car struct {
	model string
	year  int
	brand string
	price int
}

// Task:0  Printing The Cars in Inventory
func print(cars []*Car) {
	//Printing
	for _, car := range cars {
		fmt.Printf("Model: %s, Year: %d, Brand: %s, Price: %d\n", car.model, car.year, car.brand, car.price)
	}
}

// Task:1  Apend New Car
func addCar(cars []*Car, model string, year int, brand string, price int) []*Car {
	//This function create the copy 
	newCar := &Car{model, year, brand, price}
	cars = append(cars, newCar)
	return cars
}

// Task: 2 Update the car entry
func update(cars *[]*Car, price int, model string, brand string) {
	//pass by reference , Memory efficient approach
	for _, value := range *cars {
		if model == value.model && brand == value.brand {
			value.price = price
		}
	}
}

// Task:3 Most expensive car
func getMostExpensiveCar(cars *[]*Car) {
	//pass by reference , Memory efficient approach
	maxPrice := -1
	for _, value := range *cars {
		if maxPrice < value.price {
			maxPrice = value.price
		}
	}
	fmt.Println("Most Expensive Car Price is: ", maxPrice)
}

func main() {
	cars := []*Car{
		{model: "A", year: 2024, brand: "BMW", price: 100000},
		{"B", 2020, "Mercedese", 1800000},
	}
	//fmt.Println(cars) //[0x140001120c0 0x140001120f0]
	//Go's fmt.Println by default prints the memory addresses of the pointers

	fmt.Println("before")
	print(cars)
	//Task:1 Adding New Car in Inventory
	cars = addCar(cars, "C", 2010, "Alto", 10000)

	fmt.Println("\nAfter Adding new car")
	print(cars)

	//Task:2 Update the price of a car given its model and brand.
	update(&cars, 8000000, "A", "BMW")
	fmt.Println("\nAfter updating new car")
	print(cars)

	//task:3 Find the most expensive car in the inventory.
	getMostExpensiveCar(&cars)
}
