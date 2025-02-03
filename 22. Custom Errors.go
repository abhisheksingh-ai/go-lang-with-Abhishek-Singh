package main

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Can not devide by zero")
	}
	return a / b, nil
}

func main() {
	_, err := divide(4, 0)

	if err != nil {
		fmt.Println(err)
		//Can not devide by zero
	}
}
