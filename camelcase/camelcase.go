package main

import (
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	scanInput().Output(os.Stdout)
}

func scanInput() Input {
	var s Input

	fmt.Scan(&s)

	return s
}

type Input string

func (input Input) Output(w io.Writer) {
	var c = 1

	for _, r := range input {
		if unicode.IsUpper(r) {
			c++
		}
	}

	fmt.Fprintln(w, c)
}
