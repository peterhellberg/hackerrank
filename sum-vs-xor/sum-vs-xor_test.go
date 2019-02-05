package main

import "testing"

func TestSumVsXOR(t *testing.T) {
	for n, want := range map[uint]int{
		0:                1,
		4:                4,
		5:                2,
		10:               4,
		3434444444333:    262144,
		1000000000000000: 1073741824,
	} {
		if got := sumVsXOR(n); got != want {
			t.Fatalf("sumVsXOR(%d) = %d, want %d", n, got, want)
		}
	}
}
