package main

import (
	"bufio"
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestNewQueries(t *testing.T) {
	for _, tt := range []struct {
		r   io.Reader
		len int
	}{
		{strings.NewReader("3\n1 42\n2\n3"), 3},
		{strings.NewReader("2\n1 42\n2"), 2},
	} {
		queries := newQueries(tt.r)

		if got, want := len(queries), tt.len; got != want {
			t.Fatalf("len(queries) = %d, want %d", got, want)
		}
	}
}

func TestRun(t *testing.T) {
	for _, tt := range []struct {
		queries queries
		want    string
	}{
		{nil, ""},
		{queries{}, ""},
		{queries{{print, 0}}, "-1\n"},
		{queries{{push, 1}, {print, 0}}, "1\n"},
		{queries{{push, 1}, {push, 2}, {print, 0}}, "2\n"},
		{queries{{push, 1}, {print, 0}, {push, 2}, {print, 0}}, "1\n2\n"},
		{queries{
			{push, 1}, {push, 44}, {print, 0}, {print, 0}, {delete, 0},
			{print, 0}, {print, 0}, {push, 3}, {push, 37}, {delete, 0},
			{print, 0}, {push, 29}, {print, 0}, {push, 73}, {push, 51},
			{print, 0}, {print, 0}, {print, 0}, {push, 70}, {print, 0},
			{push, 8}, {delete, 0}, {push, 49}, {push, 56}, {push, 81},
			{delete, 0}, {push, 59}, {push, 44}, {delete, 0}, {print, 0},
			{print, 0}, {delete, 0}, {print, 0}, {print, 0}, {push, 4},
			{print, 0}, {push, 89}, {delete, 0}, {push, 37}, {push, 50},
			{push, 64}, {delete, 0}, {push, 49}, {push, 35}, {push, 85},
			{print, 0}, {push, 41}, {delete, 0}, {print, 0}, {print, 0},
			{push, 86}, {print, 0}, {push, 60}, {push, 8}, {print, 0},
			{push, 100}, {print, 0}, {push, 83}, {print, 0}, {push, 47},
			{delete, 0}, {push, 78}, {delete, 0}, {push, 55}, {push, 97},
			{delete, 0}, {print, 0}, {push, 40}},
			"44\n44\n1\n1\n3\n29\n73\n73\n73\n73\n73\n73\n73\n73\n73\n85\n85\n85\n86\n86\n100\n100\n100\n",
		},
	} {
		var b bytes.Buffer

		w := bufio.NewWriter(&b)

		tt.queries.Run(w)

		w.Flush()

		if got := b.String(); got != tt.want {
			t.Fatalf("b.String() = %q, want %q", got, tt.want)
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
