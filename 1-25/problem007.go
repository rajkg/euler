package main

import (
	"fmt"
)

func prime(index int) int {
	primes := make([]int, 2, 10001)

	primes[0], primes[1] = 2, 3

	for i := 2; i < index; i++ {

	NotPrime: //begins at next iteration
		for next := primes[i-1] + 2; true; next = next + 2 {
			for _, d := range primes {
				if d*d > next {
					break
				}
				if next%d == 0 {
					continue NotPrime
				}
			}
			primes = append(primes, next)
			break
		}
	}
	return primes[index-1]
}

func main() {
	fmt.Println("Prime 10001:", prime(10001))
}
