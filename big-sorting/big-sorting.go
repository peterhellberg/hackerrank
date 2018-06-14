package main

import (
	"fmt"
	"math/big"
	"sort"
	"strconv"
)

func main() {
	elements := scan()

	sort.Slice(elements, func(i, j int) bool {
		il, jl := len(elements[i]), len(elements[j])

		if il != jl {
			return il < jl
		}

		if il > 18 {
			ib, _ := new(big.Int).SetString(elements[i], 10)
			jb, _ := new(big.Int).SetString(elements[j], 10)

			return ib.Cmp(jb) < 0
		}

		si, _ := strconv.Atoi(elements[i])
		sj, _ := strconv.Atoi(elements[j])

		return si < sj
	})

	for _, e := range elements {
		fmt.Println(e)
	}
}

func scan() []string {
	var n int

	fmt.Scan(&n)

	elements := make([]string, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&elements[i])
	}

	return elements
}
