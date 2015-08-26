package main

import (
	"fmt"
)

func main() {

	for i := 1; i < 1000; i++ {
		for j := i + 1; j < 1000; j++ {
			k := 1000 - i - j
			if i*i+j*j == k*k {
				fmt.Println("Special pythagorean triplet :", i, j, k)
				fmt.Println("Their product :", i*j*k)
				return
			}
		}
	}
}
