package main

import (
	"fmt"
)

var primes = make([]int, 2, 10000)

func nextprime(max int) int {

	index := len(primes) - 1

NotPrime: //begins at next iteration
	for next := primes[index] + 2; next <= max; next = next + 2 {
		for _, d := range primes {
			if d*d > next {
				break
			}
			if next%d == 0 {
				continue NotPrime
			}
		}
		primes = append(primes, next)
		return primes[index+1]
	}

	return 0
}

func main() {

	primes[0], primes[1] = 2, 3
	sum, next := 5, 0

	for {
		if next = nextprime(2000000); next == 0 {
			break
		}
		sum = sum + next
	}

	fmt.Println("Sum of primes below 2000000: ", sum)
}
