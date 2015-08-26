package main

import (
	"fmt"
)

var primes []int

func sieve(limit int) {

	primes = make([]int, limit+1)

	mark := func(n int) {
		for i := 2 * n; i <= limit; i = i + n {
			primes[i] = 1
		}
	}

	primes[0], primes[1] = 1, 1
	for i := 2; i <= limit; i++ {
		if primes[i] == 0 {
			mark(i)
		}
	}
}

func length(a, b int) int {
	for i := 0; i < 1000; i++ {
		n := i*i + a*i + b
		if n < 0 || primes[n] == 1 {
			return i
		}
	}
	return 1000
}

func main() {

	limit := 1000

	sieve(100000)
	max, x, y := 0, 0, 0
	for a := -limit; a < limit; a++ {
		for b := -limit; b < limit; b++ {
			if l := length(a, b); l > max {
				max, x, y = l, a, b
			}
		}
	}
	fmt.Printf("a:%v b:%v length:%v \n", x, y, max)
	fmt.Println("Product: ", x*y)
}
