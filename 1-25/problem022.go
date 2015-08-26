package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func score(s string) (score int) {
	for _, c := range s {
		score = score + int(c-'A'+1)
	}
	return
}

func main() {
	names := make([]string, 0, 6000)

	line, _ := ioutil.ReadFile("./data/p022_names.txt")

	for _, name := range strings.Split(string(line), ",") {
		name = strings.TrimSuffix(strings.TrimPrefix(name, "\""), "\"")
		names = append(names, name)
	}

	sort.Sort(sort.StringSlice(names))

	total := 0
	for i, name := range names {
		total = total + (i+1)*score(name)
	}

	fmt.Println("Total name score: ", total)
}
