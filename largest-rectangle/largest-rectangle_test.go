package main

import "testing"

func TestLargestRectangleArea(t *testing.T) {
	for _, tt := range []struct {
		data []int
		want int
	}{
		{nil, 0},
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{1, 3}, 3},
		{[]int{2, 3, 1}, 4},
		{[]int{2, 3, 2}, 6},
		{[]int{3, 0, 2, 1, 2}, 3},
		{[]int{1, 3, 2, 1, 2}, 5},
		{[]int{1, 2, 1, 3, 2, 0, 1}, 5},
		{[]int{2, 2, 1, 3, 0, 3, 3}, 6},
	} {
		if got := largestRectangle(tt.data); got != tt.want {
			t.Fatalf("largestRectangle(%v) = %d, want %d", tt.data, got, tt.want)
		}
	}
}

func TestStack(t *testing.T) {
	t.Run("Empty stack", func(t *testing.T) {
		t.Run("Peek", func(t *testing.T) {
			s := &stack{}
			s.Peek()
		})

		t.Run("Pop", func(t *testing.T) {
			s := &stack{}
			s.Pop()
		})

		t.Run("Push", func(t *testing.T) {
			s := &stack{}
			s.Push(1)
			s.Push(2)

			if got, want := s.Size(), 2; got != want {
				t.Fatalf("s.Size() = %d, want %d", got, want)
			}
		})

		t.Run("Size", func(t *testing.T) {
			s := &stack{}

			if got, want := s.Size(), 0; got != want {
				t.Fatalf("s.Size() = %d, want %d", got, want)
			}
		})
	})
}
