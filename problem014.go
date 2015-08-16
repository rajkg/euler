package main

import (
	"fmt"
)

var l = map[int]int{1: 1}

func lcollatz(n int) int {

	if l[n] != 0 {
		return l[n]
	}

	if n%2 == 0 {
		l[n] = lcollatz(n/2) + 1
	} else {
		l[n] = lcollatz(3*n+1) + 1
	}

	return l[n]
}

func main() {

	max, n := 0, 0
	for i := 1000000; i > 1; i-- {
		l := lcollatz(i)
		if l > max {
			max = l
			n = i
		}
	}

	fmt.Printf("Number under 1 Million generating longest collatz sequence: %v  length: %v\n", n, max)

}
