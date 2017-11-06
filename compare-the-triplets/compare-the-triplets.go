package main

import "fmt"

func main() {
	var input [2][3]int

	fmt.Scan(
		&input[0][0], &input[0][1], &input[0][2],
		&input[1][0], &input[1][1], &input[1][2],
	)

	var alice, bob int

	for i := 0; i < 3; i++ {
		a, b := input[0][i], input[1][i]

		switch {
		case a > b:
			alice++
		case a < b:
			bob++
		}
	}

	fmt.Println(alice, bob)
}
