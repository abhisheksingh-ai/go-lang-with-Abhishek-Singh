package main

import "fmt"

type fibbonaciIterator struct {
	a, b, limit int
}

func (fi *fibbonaciIterator) next() (int, bool) {
	fi.a, fi.b = fi.b, fi.a+fi.b
	if fi.b > fi.limit {
		return 0, false
	}
	return fi.b, true
}

func main() {
	fib := fibbonaciIterator{a: -1, b: 1, limit: 10}
	for {
		val, ok := fib.next()
		if !ok {
			break
		}
		fmt.Println(val)
	}
}
