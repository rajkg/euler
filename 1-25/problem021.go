package main

import (
	"fmt"
)

func main() {

	size := 10000
	D := make([]int, size)

	for i := 4; i < size; i++ {
		d := 1
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				if q := i / j; q == j {
					d = d + j
				} else {
					d = d + j + i/j
				}
			}
		}
		D[i] = d
	}

	sum := 0
	for i := 4; i < size; i++ {
		if D[i] != 1 {
			for j := i + 1; j < size; j++ {
				if D[i] == j && D[j] == i {
					fmt.Printf("(%v, %v), ", i, j)
					sum = sum + i + j
				}
			}
		}
	}
	fmt.Println("")

	fmt.Println("Sum of amicable numbers : ", sum)

}
