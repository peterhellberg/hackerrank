package main

import (
	"fmt"
	"math/big"
)

func main() {
	var n, v int

	fmt.Scan(&n)

	bi := new(big.Int)

	for i := 0; i < n; i++ {
		if _, err := fmt.Scan(&v); err == nil {
			bi.Add(bi, big.NewInt(int64(v)))
		}
	}

	fmt.Println(bi)
}
