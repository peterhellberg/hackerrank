package main

import "fmt"

func main() {
	var n, v int

	fmt.Scan(&n)

	values := map[int]int{}

	for i := 0; i < n; i++ {
		if _, err := fmt.Scan(&v); err == nil {
			values[v]++
		}
	}

	for k, v := range values {
		if v == 1 {
			fmt.Println(k)
		}
	}
}
