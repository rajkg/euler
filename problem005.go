package main

import (
	"fmt"
)

func main() {

	var i, j int
	lcm, num := 1, []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	for i = 2; i < 20; i++ {
		for factor := true; factor == true; {
			for factor, j = false, 0; j < len(num); j++ {
				if num[j]%i == 0 {
					factor = true
					num[j] = num[j] / i
				}
			}
			if factor {
				lcm = lcm * i
			}
		}
	}

	fmt.Println("LCM :", lcm)
}
