package main

import "fmt"

func main() {
	var q int

	fmt.Scan(&q)

	for i := 0; i < q; i++ {
		var n uint32

		fmt.Scan(&n)
		fmt.Println(flippingBits(n))
	}
}

func flippingBits(n uint32) uint32 {
	return ^n
}
