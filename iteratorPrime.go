package main
import "fmt"

type PrimeIterator struct {
	num   int
	limit int
}

func isPrime(num int) bool {
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func (p *PrimeIterator) next() (int, bool) {
	p.num++
	for p.num <= p.limit {
		if isPrime(p.num) {
			return p.num, true
		}
		p.num++
	}
	return 0, false
}

func main() {
	PI := PrimeIterator{num: 1, limit: 20}
	//printing all prime numbers till 20
	for {
		val, ok := PI.next()
		if !ok {
			break
		}
		fmt.Println(val)
	}
}
