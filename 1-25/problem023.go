package main

import (
	"fmt"
)

func SumOfDivisors(x int) int {

	n, p := x, 1
	for n%2 == 0 {
		n, p = n/2, p*2
	}
	s := p*2 - 1
	//fmt.Println(2, s)

	for i := 3; i < n; i = i + 2 {
		p = 1
		for n%i == 0 {
			n, p = n/i, p*i
		}
		if p != 1 {
			d := ((p*i - 1) / (i - 1))
			s = s * d
			//fmt.Println(i, p, d)
		}
	}

	if n != 1 {
		d := ((n*n - 1) / (n - 1))
		s = s * d
		//fmt.Println(n, d)
	}

	return s - x
}

func main() {

	limit := 28123
	abundant := make([]bool, limit+1)
	numbers := make([]int, 0, limit)

	for i := 1; i <= limit; i++ {
		if SumOfDivisors(i) > i {
			abundant[i], numbers = true, append(numbers, i)
		}
	}

	sum := 0
NotFit:
	for i := 1; i <= limit; i++ {
		for _, n := range numbers {
			if i >= n && abundant[i-n] == true {
				continue NotFit
			}
		}
		sum = sum + i
	}

	fmt.Println("\nSum: ", sum)
}
