package main

import "testing"

func TestTheGreatXOR(t *testing.T) {
	for n, want := range map[uint]uint{
		2:         1,
		10:        5,
		217:       38,
		111888516: 22329211,
		303924675: 232946236,
	} {
		if got := theGreatXOR(n); got != want {
			t.Fatalf("theGreatXOR(%d) = %d, want %d", n, got, want)
		}
	}
}
