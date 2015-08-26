package main

import (
	"fmt"
)

func main() {

	for n := 1; true; n++ {
		tr := n * (n + 1) / 2
		numd := 2
		for d := 2; d*d <= tr; d++ {
			if tr%d == 0 {
				numd = numd + 2
			}
		}
		if tr == d*d {
			numd--
		}

		if numd >= 500 {
			fmt.Println("First triangle number with more than 500 divisors: ", tr)
			return
		}
	}
}
