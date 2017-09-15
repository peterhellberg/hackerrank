package main

import "fmt"

func main() {
	var n, sum int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var v int

		if _, err := fmt.Scan(&v); err == nil {
			sum += v
		}
	}

	fmt.Println(sum)
}
