package main
import "fmt"

type EvenIterator struct {
	current int
	max     int
}

func (e *EvenIterator) next() (int, bool) {
	e.current++
	for e.current <= e.max {
		if e.current%2 == 0 {
			return e.current, true
		}
		e.current++
	}
	return 0, false
}

func main() {
	iter := EvenIterator{current: 1, max: 10}
	for {
		val, ok := iter.next()

		if !ok {
			break
		}
		fmt.Println(val)
	}
}
