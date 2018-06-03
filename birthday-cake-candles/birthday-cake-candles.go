package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	scanCandles().Output(os.Stdout)
}

func scanCandles() Candles {
	var n int

	fmt.Scan(&n)

	c := Candles(make([]int, n))

	for i := 0; i < n; i++ {
		fmt.Scan(&c[i])
	}

	return c
}

type Candles []int

func (c Candles) Output(w io.Writer) {
	counter := map[int]int{}

	var tallest int

	for _, candle := range c {
		if candle > tallest {
			tallest = candle
		}

		counter[candle]++
	}

	fmt.Fprintln(w, counter[tallest])
}
