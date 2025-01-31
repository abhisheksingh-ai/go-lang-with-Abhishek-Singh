package main

import (
	"errors"
	"fmt"
)

func main() {
	//In go lang error is value not the exception
	//Go works on explicit error handling

	// 1. Creating a simple errors
	orginalErr := errors.New("Orginal error")

	// 2. Wrapping errors
	//Sometimes you want to add more context to an error
	//while preseerving the original error
	wrappedErr := fmt.Errorf("Wrapped error: %w", orginalErr)
	fmt.Println(wrappedErr) //Wrapped error: Orginal error

	// 3. Unwrapped error
	//To retrive the original error from a wrapped error
	//When you need to inspect underlying cause of error
	unwrappedErr := errors.Unwrap(wrappedErr)
	fmt.Println(unwrappedErr) // Orginal error

	// 4. Checking for specific error
	//Instead of checking errors directly, use errors.Is()
	//to check if an error (or any error it wrap) matches a
	//specific error
	if errors.Is(wrappedErr, orginalErr) {
		fmt.Println("The error is original error") //The error is original error
	}

	//Dont do check like this, It wont work
	// if errors.Is(orginalErr, wrappedErr) {
	// 	fmt.Println("The error is original error")
	// }

	// 5. Combinig multiple error
	e1 := errors.New("Error 1")
	e2 := errors.New("Error 2")

	combineErr := errors.Join(e1, e2)
	fmt.Println(combineErr)
	//Error 1
	//Error 2
}
