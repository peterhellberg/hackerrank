package main

import (
	"flag"
	"fmt"
	"image/color"
	"io"
	"os"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

var w, h int

func run() {
	var fw, fh = float64(w), float64(h)

	win, err := pixelgl.NewWindow(pixelgl.WindowConfig{
		Bounds:      pixel.R(0, 0, fw, fh),
		VSync:       true,
		Undecorated: true,
	})
	if err != nil {
		panic(err)
	}

	var (
		imd = imdraw.New(nil)
		p   = newPiles(os.Stdin)
		u   = (fh - 20) / float64(p.maxHeight)
		cw  = (fw - 20) / 3
		eh  = p.Tallest().height
	)

	hc := color.RGBA{0xFF, 0xFF, 0xFF, 0xFF}

	update := func() {
		if p.Equal() {
			hc = color.RGBA{0x75, 0xEF, 0x5B, 0xFF}
		}

		imd.Clear()
		imd.Color = hc
		imd.Push(pixel.V(0, 0), pixel.V(fw, 12+(float64(eh)*u)))
		imd.Rectangle(0)

		imd.Color = color.RGBA{0xA0, 0xA0, 0xA0, 0xFF}
		imd.Push(pixel.V(0, 16+(float64(eh)*u)), pixel.V(fw, 16+(float64(eh)*u)))
		imd.Line(6)

		for i, c := range p.a.cylinders {
			var h int

			for _, c := range p.a.cylinders[:i] {
				h += c
			}

			hu := 10 + float64(h)*u
			cu := float64(c) * u

			imd.Color = color.RGBA{0x33, 0x99, 0xFF, 0xFF}
			imd.Push(pixel.V(20, hu), pixel.V(cw, hu+cu))
			imd.Rectangle(0)
			imd.Color = color.RGBA{0x01, 0x57, 0xAA, 0xFF}
			imd.Push(pixel.V(20, hu), pixel.V(cw, hu+cu))
			imd.Rectangle(6)
		}

		for i, c := range p.b.cylinders {
			var h int

			for _, c := range p.b.cylinders[:i] {
				h += c
			}

			hu := 10 + float64(h)*u
			cu := float64(c) * u

			imd.Color = color.RGBA{0x00, 0xCC, 0xCC, 0xFF}
			imd.Push(pixel.V(20+cw, hu), pixel.V(2*cw, hu+cu))
			imd.Rectangle(0)
			imd.Color = color.RGBA{0x00, 0x79, 0x79, 0xFF}
			imd.Push(pixel.V(20+cw, hu), pixel.V(2*cw, hu+cu))
			imd.Rectangle(6)
		}

		for i, c := range p.c.cylinders {
			var h int

			for _, c := range p.c.cylinders[:i] {
				h += c
			}

			hu := 10 + float64(h)*u
			cu := float64(c) * u

			imd.Color = color.RGBA{0x99, 0x99, 0xFF, 0xFF}
			imd.Push(pixel.V(20+2*cw, hu), pixel.V(3*cw, hu+cu))
			imd.Rectangle(0)
			imd.Color = color.RGBA{0x5A, 0x5A, 0xCE, 0xFF}
			imd.Push(pixel.V(20+2*cw, hu), pixel.V(3*cw, hu+cu))
			imd.Rectangle(6)
		}
	}

	update()

	for !win.Closed() {
		win.SetClosed(win.JustPressed(pixelgl.KeyEscape) || win.JustPressed(pixelgl.KeyQ))

		if win.JustPressed(pixelgl.KeySpace) {
			if !p.Equal() {
				p.Tallest().Pop()

				eh = p.Tallest().height

				update()
			} else {
				win.SetClosed(true)
			}
		}

		win.Clear(color.RGBA{0xBB, 0xBB, 0xBB, 0xFF})

		imd.Draw(win)
		win.Update()
	}

	fmt.Println(p.EqualHeight())
}

func main() {
	flag.IntVar(&w, "w", 640, "Width")
	flag.IntVar(&h, "h", 360, "Height")

	flag.Parse()

	pixelgl.Run(run)
}

type piles struct {
	maxCount  int
	maxHeight int
	a         *pile
	b         *pile
	c         *pile
}

func newPiles(r io.Reader) *piles {
	var n1, n2, n3 int

	go func() {
		time.Sleep(100 * time.Millisecond)
		if n3 == 0 {
			fmt.Println("Got no input")
			os.Exit(1)
		}
	}()

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

	maxCount := n1

	switch {
	case n2 > n1 && n2 > n3:
		maxCount = n3
	case n3 > n1 && n3 > n2:
		maxCount = n3
	}

	a, b, c := newPile(c1), newPile(c2), newPile(c3)

	maxHeight := a.height

	switch {
	case b.height > a.height && b.height > c.height:
		maxHeight = b.height
	case c.height > a.height && c.height > b.height:
		maxHeight = c.height
	}

	return &piles{maxCount, maxHeight, a, b, c}
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
