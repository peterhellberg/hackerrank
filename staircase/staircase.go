package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	scanInput().Output(os.Stdout)
}

func scanInput() Input {
	var n Input

	fmt.Scan(&n)

	return n
}

type Input int

func (input Input) Output(w io.Writer) {
	for i := 1; i <= int(input); i++ {
		fmt.Print(strings.Repeat(" ", int(input)-i), strings.Repeat("#", i), "\n")
	}
}
