package main

import (
	"fmt"
)

func palindrome(n int) bool {

	r := 0
	for i := n; i > 0; i = i / 10 {
		r = r*10 + i%10
	}

	return r == n
}

func main() {

	largest := 0
	for i := 999; i > 100; i-- {
		for j := i; j > 100; j-- {
			p := i * j
			if p > largest && palindrome(p) {
				largest = p
			}
		}
	}

	fmt.Println("largest palindrome product of 3 digit numbers: ", largest)
}
