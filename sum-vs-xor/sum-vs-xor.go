package main

import (
	"fmt"
	"math/bits"
)

func main() {
	var n uint

	fmt.Scan(&n)
	fmt.Println(sumVsXOR(n))
}

func sumVsXOR(n uint) int {
	return 1 << uint(bits.Len(n)-bits.OnesCount(n))
}
