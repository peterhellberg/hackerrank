package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	m := readImage(os.Stdin)

	saveImage(m)

	fmt.Println(dayOrNight(m))
}

func dayOrNight(m *image.RGBA) string {
	if averageDay(m) {
		return "day"
	}

	return "night"
}

func averageDay(m *image.RGBA) bool {
	c := len(m.Pix)

	var t int

	for _, v := range m.Pix {
		t += int(v)
	}

	return t/c > 110
}

func readImage(r io.Reader) *image.RGBA {
	grid := [][]color.RGBA{}

	s := bufio.NewScanner(r)

	for s.Scan() {
		row := []color.RGBA{}

		for _, cs := range strings.Split(s.Text(), " ") {
			var r, g, b uint8

			fmt.Sscanf(cs, "%d,%d,%d", &r, &g, &b)

			row = append(row, color.RGBA{r, g, b, 255})
		}

		grid = append(grid, row)
	}

	w, h := len(grid[0]), len(grid)

	m := image.NewRGBA(image.Rect(0, 0, w, h))

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			m.Set(x, y, grid[y][x])
		}
	}

	return m
}

func saveImage(m *image.RGBA) {
	if os.Getenv("DEBUG") == "true" {
		prefix := "digital-camera-day-or-night-"

		if file, err := ioutil.TempFile("/tmp/", prefix); err == nil {
			defer file.Close()
			png.Encode(file, m)
			fmt.Fprintln(os.Stderr, file.Name())
		}
	}
}
