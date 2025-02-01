package main

import "fmt"

type Student struct {
	name   string
	age    int
	id     int
	branch string
	hobby  string
}

func main() {
	s1 := Student{"Abhishek", 24, 1234, "chemical", "cricket"}
	fmt.Println(s1)
	//{Abhishek 24 1234 chemical cricket}

	//is it Mutable?? lets see
	s1.hobby = "Football"
	s1.id = 5678
	fmt.Println(s1)
	//{Abhishek 24 5678 chemical Football}

	//instance of Student with specific fields others are 0 or ""
	s2 := Student{name: "rahul"}
	fmt.Println(s2) //{rahul 0 0  }
}
