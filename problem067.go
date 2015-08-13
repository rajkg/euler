package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var tree = make([][]int64, 0, 100)

func load(file string) error {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		nums := strings.Split(line, " ")
		row := make([]int64, len(nums))
		for i, num := range nums {
			if row[i], err = strconv.ParseInt(num, 10, 64); err != nil {
				return err
			}
		}
		tree = append(tree, row)
	}

	return nil
}

func main() {

	if err := load("./data/p067_triangle.txt"); err ! nil {
		fmt.Println(err)
	}

	for i := len(tree) - 2; i >= 0; i-- {
		for j := 0; j < len(tree[i]); j++ {
			if tree[i+1][j] > tree[i+1][j+1] {
				tree[i][j] = tree[i][j] + tree[i+1][j]
			} else {
				tree[i][j] = tree[i][j] + tree[i+1][j+1]
			}
		}
	}

	fmt.Println("Max path sum: ", tree[0][0])
}
