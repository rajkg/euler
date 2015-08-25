package main

import (
	"fmt"
)

func main() {

	limit := 10
	permute, bag := make([]int, 0, limit), make([]int, limit)
	for i := 0; i < limit; i++ {
		bag[i] = i
	}

	place := func() {
		permute, bag = append(permute, bag[0]), bag[1:]
	}

	replace := func() bool {
		p := permute[len(permute)-1]

		for i := 0; i < len(bag); i++ {
			if bag[i] > p {
				permute[len(permute)-1], bag[i] = bag[i], p
				return true
			}
		}

		permute, bag = permute[:len(permute)-1], append(bag, p)
		return len(permute) == 0
	}

	place()

Done:
	for i := 0; true; {
		switch {
		case len(permute) < limit:
			place()
		case len(permute) == limit:
			if i = i + 1; i == 1000000 {
				break Done
			}
			for !replace() {
			}
		}
	}

	number := 0
	for _, r := range permute {
		number = number*10 + r
	}
	fmt.Println(number)
}
