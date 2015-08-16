package main

import (
	"fmt"
)

var g [][]int

func paths(dx int, dy int) (total int) {

	switch {
	case dx == 0 || dy == 0:
		total = 1
	case g[dx][dy] != 0:
		total = g[dx][dy]
	default:
		total = paths(dx-1, dy) + paths(dx, dy-1)
		g[dx][dy] = total
	}

	return total
}

func main() {

	size := 21
	g = make([][]int, 0, size)
	for i := 0; i < size; i++ {
		g = append(g, make([]int, size))
	}

	fmt.Println("Total lattice paths in 20x20 grid: ", paths(size-1, size-1))
}
