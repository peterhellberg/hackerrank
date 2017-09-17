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
	m := &stack{}

	for _, query := range q {
		switch query.cmd {
		case push:
			if m.Empty() {
				m.Push(query.arg)
			} else {
				if query.arg > m.Peek() {
					m.Push(query.arg)
				} else {
					m.Push(m.Peek())
				}
			}
		case delete:
			if !m.Empty() {
				m.Pop()
			}
		case print:
			fmt.Fprintln(w, m.Peek())
		}
	}
}

type stack struct {
	// You could have a mutex here if you need the stack to be thread safe
	items []int
	size  int
}

func (s *stack) Empty() bool {
	return s.size == 0
}

func (s *stack) Peek() int {
	if s.Empty() {
		return -1
	}

	return s.items[s.size-1]
}

func (s *stack) Pop() int {
	if s.Empty() {
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
