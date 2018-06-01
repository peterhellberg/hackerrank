package main

import "fmt"

func main() {
	fmt.Println(scanMatrix().Result())
}

func scanMatrix() Matrix {
	var n int

	fmt.Scan(&n)

	var m Matrix

	for i := 0; i < n; i++ {
		var row []int

		for j := 0; j < n; j++ {
			var e int

			fmt.Scan(&e)

			row = append(row, e)
		}

		m = append(m, row)
	}

	return m
}

type Matrix [][]int

func (m Matrix) Result() int {
	return abs(sum(m.Primary()) - sum(m.Secondary()))
}

func (m Matrix) Primary() []int {
	var p []int

	for i, r := range m {
		p = append(p, r[i])
	}

	return p
}

func (m Matrix) Secondary() []int {
	var s []int

	o := len(m) - 1

	for i, r := range m {
		s = append(s, r[o-i])
	}

	return s
}

func sum(elements []int) int {
	var n int

	for _, e := range elements {
		n += e
	}

	return n
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}
