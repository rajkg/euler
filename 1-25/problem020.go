package main

import (
	"fmt"
)

func main() {

	number := 100
	n := make([]int, 1, 1000)
	n[0] = 1

	for i := 1; i <= number; i++ {
		r := 0
		for j := 0; j < len(n); j++ {
			p := n[j]*i + r
			n[j], r = p%10, p/10
		}
		for ; r > 0; r = r / 10 {
			n = append(n, r%10)
		}
	}

	sum := 0
	for _, d := range n {
		sum = sum + d
	}

	fmt.Println("Sum of digits: ", sum)
}
