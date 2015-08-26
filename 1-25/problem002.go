package main

import (
	"fmt"
)

func fibonacci(max int) <-chan int {
	c := make(chan int)

	go func() {
		defer close(c)
		i0, i1 := 0, 1
		for {
			i0, i1 = i1, i0+i1
			if i1 > max {
				return
			}
			c <- i1
		}
	}()

	return c
}

func IsEven(num int) bool {
	return num&0x1 == 0
}

func main() {

	c := fibonacci(4000000)

	sum := 0
	for i := range c {
		if IsEven(i) {
			sum = sum + i
		}
	}

	fmt.Println("Sum of even fibonacci numbers less than 4 Million: ", sum)
}
