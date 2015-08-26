package main

import (
	"fmt"
)

func prime(n int) bool {

	if n%2 == 0 {
		return false
	}

	for i := 3; i*i <= n; i = i + 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}

/*
 Finds length of repeating cycle for prime n by long division
*/
func cyclic(n int) bool {
	t, r := 0, 1

	for {
		t = t + 1
		x := r * 10
		r = x % n
		if r == 1 {
			break
		}
	}

	if t == n-1 {
		return true
	}

	return false

}

func main() {

	// Longest repeating cycles are produced by prime numbers
	//
	// if 10 is a primitive root modulo p, then length of repeat cycle is p-1
	//
	for i := 999; i > 1; i-- {
		if prime(i) && cyclic(i) {
			fmt.Println("Largest cyclic < 1000: ", i)
			return
		}
	}
}
