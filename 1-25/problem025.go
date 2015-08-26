package main

import (
	"fmt"
	"math/big"
)

func main() {

	a := big.NewInt(1)
	b := big.NewInt(1)

	var limit big.Int
	limit.Exp(big.NewInt(10), big.NewInt(999), nil)

	var i int
	for i = 1; a.Cmp(&limit) < 0; i++ {
		a.Add(a, b)
		a, b = b, a
	}
	fmt.Printf("Fib %v: %v\n", i, a)
}
