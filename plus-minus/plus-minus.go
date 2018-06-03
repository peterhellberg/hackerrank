package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Print(scanInput().Output(6))
}

func scanInput() Input {
	var input Input

	fmt.Scan(&input.n)

	input.elements = make([]int64, input.n)

	for i := 0; i < int(input.n); i++ {
		fmt.Scan(&input.elements[i])
	}

	return input
}

type Input struct {
	n        int64
	elements []int64
}

func (i Input) Output(precision int) Output {
	var pc, zc, nc int64

	for _, e := range i.elements {
		switch {
		case e > 0:
			pc++
		case e < 0:
			nc++
		case e == 0:
			zc++
		}
	}

	return Output{
		precision,
		big.NewRat(pc, i.n),
		big.NewRat(nc, i.n),
		big.NewRat(zc, i.n),
	}
}

type Output struct {
	precision int
	positive  *big.Rat
	negative  *big.Rat
	zeros     *big.Rat
}

func (o Output) String() string {
	return fmt.Sprintf("%s\n%s\n%s\n",
		o.positive.FloatString(o.precision),
		o.negative.FloatString(o.precision),
		o.zeros.FloatString(o.precision),
	)
}
