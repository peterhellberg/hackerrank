package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

	var input = make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&input[i])
	}

	fmt.Println(largestRectangle(input))
}

func largestRectangle(input []int) int {
	var (
		largest, pos, top int

		positions = &stack{}
		heights   = &stack{}

		update = func() {
			top = positions.Pop()

			if area := heights.Pop() * (pos - top); area > largest {
				largest = area
			}
		}
	)

	for pos = 0; pos < len(input); pos++ {
		h := input[pos]

		if heights.Size() == 0 || h > heights.Peek() {
			heights.Push(h)
			positions.Push(pos)
		} else if h < heights.Peek() {
			for heights.Size() > 0 && h < heights.Peek() {
				update()
			}

			heights.Push(h)
			positions.Push(top)
		}
	}

	for heights.Size() > 0 {
		update()
	}

	return largest
}

type stack struct {
	// You could have a mutex here if you need the stack to be thread safe
	items []int
}

func (s *stack) Peek() int {
	size := len(s.items)
	if size == 0 {
		return -1
	}

	return s.items[size-1]
}

func (s *stack) Pop() int {
	size := len(s.items)
	if size == 0 {
		return -1
	}

	item := s.items[size-1]

	s.items = s.items[:size-1]

	return item
}

func (s *stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *stack) Size() int {
	return len(s.items)
}
