package main

import "fmt"

type Set[K, V comparable] struct {
	mpp map[K]V
}

// constructor
func NewSet[K, V comparable]() *Set[K, V] {
	return &Set[K, V]{mpp: make(map[K]V)}
}

// adding value in Set
func (s *Set[K, V]) add(key K, val V) {
	_, prs := s.mpp[key]
	if prs == false {
		s.mpp[key] = val
	}
}

// Checking is this key present in the map
func (s *Set[K, V]) check(key K) {
	_, prs := s.mpp[key]
	if prs == false {
		fmt.Println(key, "key not found")
	} else {
		fmt.Println(key, "key found")
	}
}

// Delete particular key
func (s *Set[K, V]) deleteKey(key K) {
	_, prs := s.mpp[key]
	if prs == true {
		delete(s.mpp, key)
	}
}

func main() {
	st := NewSet[string, bool]()

	st.mpp["Abhishek"] = true
	st.mpp["Singh"] = true

	st.check("Abhishek") //Abhishek Key found
	st.check("Mtech")    //Mtech Key not found

	st.add("Kharagpur", true)
	st.add("GoLang", true)

	st.check("GoLang") //GoLang Key found

	st.deleteKey("GoLang")
	st.check("GoLang")

	//making different set
	//[string]---->int
	st2 := NewSet[string, int]()

	st2.add("Abhishek", 1)
	st2.add("Singh", 1)

	st2.check("Abhishek") //Abhishek key found
	st2.check("Mtech") //Mtech key not found
}
