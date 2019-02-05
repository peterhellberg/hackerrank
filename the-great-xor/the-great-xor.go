package main

import (
	"fmt"
	"math/bits"
)

func main() {
	for _, q := range scanQueries() {
		fmt.Println(theGreatXOR(q))
	}
}

func scanQueries() []uint {
	var n int

	fmt.Scan(&n)

	var queries []uint

	for i := 0; i < n; i++ {
		var q uint

		if _, err := fmt.Scan(&q); err == nil {
			queries = append(queries, q)
		}
	}

	return queries
}

func theGreatXOR(x uint) uint {
	return 1<<uint(bits.Len(x)) - x - 1
}
