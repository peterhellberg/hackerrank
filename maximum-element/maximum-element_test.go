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
		{queries{{push, 1}, {cmd: print}}, "1\n"},
		{queries{{push, 1}, {push, 2}, {cmd: print}}, "2\n"},
		{queries{{push, 1}, {cmd: print}, {push, 2}, {cmd: print}}, "1\n2\n"},
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
