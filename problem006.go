package main

import (
	"fmt"
)

func main() {

	n := 100

	sum := (n * (n + 1)) / 2
	squaresum := sum * sum
	sumsquare := (n * (n + 1) * (2*n + 1)) / 6

	fmt.Printf("Sum: %v SquareSum: %v Sumsquare: %v \n", sum, sumsquare, squaresum)
	fmt.Println("Difference between Squaresum and Sumsquare:", squaresum-sumsquare)
}
