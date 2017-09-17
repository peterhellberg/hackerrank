package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	newQueries(os.Stdin).Run(os.Stdout)
}

type command int

const (
	_ command = iota
	push
	delete
	print
)

type query struct {
	cmd command
	arg int
}

type queries []query

func newQueries(r io.Reader) queries {
	var n int

	fmt.Fscan(r, &n)

	var q = make([]query, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(r, &q[i].cmd)

		if q[i].cmd == push {
			fmt.Fscan(r, &q[i].arg)
		}
	}

	return q
}

func (q queries) Run(w io.Writer) {
	s := &stack{}

	for _, query := range q {
		switch query.cmd {
		case push:
			s.Push(query.arg)
		case delete:
			s.Pop()
		case print:
			fmt.Fprintln(w, s.Max())
		}
	}
}

type stack struct {
	// You could have a mutex here if you need the stack to be thread safe
	items []int
	size  int
}

func (s *stack) Max() int {
	var max int = -1

	for i, item := range s.items {
		if i == 0 || item > max {
			max = item
		}
	}

	return max
}

func (s *stack) Peek() int {
	if s.size == 0 {
		return -1
	}

	return s.items[s.size-1]
}

func (s *stack) Pop() int {
	if s.size == 0 {
		return -1
	}

	s.size--

	item := s.items[s.size]
	s.items = s.items[:s.size]

	return item
}

func (s *stack) Push(item int) {
	s.items = append(s.items, item)
	s.size++
}

func (s *stack) Size() int {
	return s.size
}
