package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println(newPiles(os.Stdin).EqualHeight())
}

type piles struct {
	a *pile
	b *pile
	c *pile
}

func newPiles(r io.Reader) *piles {
	var n1, n2, n3 int

	fmt.Fscan(r, &n1, &n2, &n3)

	var c1 = make([]int, n1)
	var c2 = make([]int, n2)
	var c3 = make([]int, n3)

	for i := n1 - 1; i >= 0; i-- {
		fmt.Fscan(r, &c1[i])
	}

	for i := n2 - 1; i >= 0; i-- {
		fmt.Fscan(r, &c2[i])
	}

	for i := n3 - 1; i >= 0; i-- {
		fmt.Fscan(r, &c3[i])
	}

	return &piles{newPile(c1), newPile(c2), newPile(c3)}
}

func (p *piles) EqualHeight() int {
	for !p.Equal() {
		p.Tallest().Pop()
	}

	return p.Tallest().height
}

func (p *piles) Equal() bool {
	return p.a.height == p.b.height && p.b.height == p.c.height
}

func (p *piles) Tallest() *pile {
	switch {
	case p.a.height >= p.b.height && p.a.height >= p.c.height:
		return p.a
	case p.b.height >= p.a.height && p.b.height >= p.c.height:
		return p.b
	default:
		return p.c
	}
}

type pile struct {
	cylinders []int
	height    int
}

func newPile(cylinders []int) *pile {
	var height int

	for _, c := range cylinders {
		height += c
	}

	return &pile{cylinders, height}
}

func (p *pile) Pop() int {
	if p.height == 0 {
		return -1
	}

	last := len(p.cylinders) - 1

	item := p.cylinders[last]

	p.cylinders = p.cylinders[:last]

	p.height -= item

	return item
}
