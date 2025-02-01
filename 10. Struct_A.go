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
	//How to make the instances of Student Struct
	s1 := Student{}

	fmt.Println(s1)
	// { 0 0  }

	s1.name = "Abhishek"
	s1.age = 24
	s1.id = 1234
	s1.branch = "Chemical"
	s1.hobby = "Football"

	fmt.Println(s1)
	//{Abhishek 24 1234 Chemical Football}
	//if i want to print s1 with its field name than i have to do
	
	fmt.Printf("%+v", s1)
	//{name:Abhishek age:24 id:1234 branch:Chemical hobby:Football}
}
