package main

import (
	"fmt"
)

func main() {

	pow := 1000
	n := make([]int, 1, 1000)

	n[0] = 1
	for i := 0; i < pow; i++ {
		r := 0
		for j := 0; j < len(n); j++ {
			p := n[j]*2 + r
			n[j], r = p%10, p/10
		}
		if r > 0 {
			n = append(n, r)
		}
	}

	sum := 0
	for _, d := range n {
		sum = sum + d
	}

	fmt.Println("Sum of digits: ", sum)
}
