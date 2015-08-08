package main

import (
	"fmt"
)

func main() {

	num := 600851475143
	for i := 2; i*i < num; i++ {
		if num%i == 0 {
			fmt.Printf("%v, ", i)
			num = num / i
		}
	}

	fmt.Println(num)
	fmt.Println("Largest prime factor: ", num)
}
