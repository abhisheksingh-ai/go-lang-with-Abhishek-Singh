package main

import "fmt"

type Set struct {
	mpp map[string]bool
}

// constructor
func NewSet() *Set {
	return &Set{make(map[string]bool)}
}

// adding value in the set
func (s *Set) add(key string) {
	_, prs := s.mpp[key]

	if prs == false {
		//key not present in the map
		s.mpp[key] = true
	}
}

// checking is this key present in the map
func (s *Set) check(key string) {
	_, prs := s.mpp[key]

	if prs == false {
		fmt.Println(key, "Key not found")
	} else {
		fmt.Println(key, "Key found")
	}
}

// delete any key
func (s *Set) deleteKey(key string) {
	_, prs := s.mpp[key]
	if prs == true {
		delete(s.mpp, key)
	}
}

func main() {
	st := NewSet()

	st.mpp["Abhishek"] = true
	st.mpp["Singh"] = true

	st.check("Abhishek") //Abhishek Key found
	st.check("Mtech")    //Mtech Key not found

	st.add("Kharagpur")
	st.add("GoLang")

	st.check("GoLang") //GoLang Key found

	st.deleteKey("GoLang")
	st.check("GoLang")
}
