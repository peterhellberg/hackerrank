package main

import (
	"fmt"
	"sort"
)

func main() {
	var v = make([]int, 5)

	for i := 0; i < 5; i++ {
		fmt.Scan(&v[i])
	}

	sort.Ints(v)

	fmt.Println(sum(v[:4]), sum(v[1:]))
}

func sum(v []int) int {
	var s int

	for _, v := range v {
		s += v
	}

	return s
}
