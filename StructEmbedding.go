/*
Problem Statement: Build a Vehicle Management System
You are tasked with designing a simple Vehicle Management System using Go. Each Vehicle has common fields like Make, Model, and Year. Depending on the type of vehicle (e.g., Car or Truck), additional fields and methods are required.

Requirements:
Vehicle is the base struct that contains:

Make (string)
Model (string)
Year (int)
Car and Truck are two types of vehicles:

A Car has an additional field Seats (int) and a method Honk() that prints: "Car Honk: Beep Beep!"
A Truck has an additional field PayloadCapacity (int) and a method Honk() that prints: "Truck Honk: Horn Horn!"
Both Car and Truck should embed the Vehicle struct so that the Make, Model, and Year fields can be accessed directly.

Create a function PrintDetails(v Vehicle) that prints the details of a Car or Truck.

In the main function:

Create an array of Car and Truck instances.
Loop through them and:
Print their details using the PrintDetails function.
Call their respective Honk() method.

*/

package main

import "fmt"

type Vehical struct {
	Make  string
	Model string
	Year  int
}

type Car struct {
	Vehical
	Seat int
}

type Truck struct {
	Vehical
	PayLoadCapacity int
}

type PrintDetailsInter interface {
	PrintDetails()
	Honk()
}

// Print Horn for Car
func (c Car) Honk() {
	fmt.Println("\nCar Honk: Beep Beep")
}

// Print Horn for Truck
func (t Truck) Honk() {
	fmt.Println("\nTruck Honk: Horn Horn")
}

// Print details for Car
func (c Car) PrintDetails() {
	fmt.Printf("Make: %s, Model: %s, Year: %d, Seat: %d", c.Make, c.Model, c.Year, c.Seat)
}

// //Print details for Truck
func (t Truck) PrintDetails() {
	fmt.Printf("Make: %s, Model: %s, Year: %d, PayLoadCapacity: %d", t.Make, t.Model, t.Year, t.PayLoadCapacity)
}

func main() {
	car := Car{
		Vehical: Vehical{Make: "Toyota", Model: "Corola", Year: 2022},
		Seat:    5,
	}

	truck := Truck{
		Vehical:         Vehical{Make: "Ford", Model: "F-150", Year: 2000},
		PayLoadCapacity: 120,
	}

	vehicals := []PrintDetailsInter{car, truck} //Key

	for _, val := range vehicals {
		val.PrintDetails()
		val.Honk()
	}
}
