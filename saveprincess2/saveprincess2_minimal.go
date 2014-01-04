package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
)

func (g *Game) NextMove() string {
	if g.Bot.Y == g.Princess.Y {
		if g.Bot.X > g.Princess.X {
			return "LEFT"
		} else {
			return "RIGHT"
		}
	}

	if g.Bot.X == g.Princess.X {
		if g.Bot.Y > g.Princess.Y {
			return "UP"
		} else {
			return "DOWN"
		}
	}

	if g.Bot.Y > g.Princess.Y {
		return "UP"
	} else {
		return "DOWN"
	}

	if g.Bot.X > g.Princess.X {
		return "LEFT"
	} else {
		return "RIGHT"
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	size := getInt(nextLine(reader))
	bot := getBot(nextLine(reader))
	board, princess := getBoardAndPrincess(reader)

	game := Game{size, bot, princess, board}

	fmt.Println(game.NextMove())
}

func getBoardAndPrincess(reader *bufio.Reader) ([][]byte, Pos) {
	var board [][]byte
	princess := Pos{}

	y := 0

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		x := bytes.Index(line, []byte("p"))

		if x > -1 {
			princess.X = x
			princess.Y = y
		}

		board = append(board, line)

		y++
	}

	return board, princess
}

func getBot(line []byte) Pos {
	pos := bytes.Split(line, []byte(" "))

	return Pos{getInt(pos[1]), getInt(pos[0])}
}

func getInt(bytes []byte) int {
	i, _ := strconv.Atoi(string(bytes))

	return i
}

func nextLine(reader *bufio.Reader) []byte {
	line, _, _ := reader.ReadLine()

	return line
}

type Game struct {
	Size     int
	Bot      Pos
	Princess Pos
	Board    [][]byte
}

type Pos struct {
	X int
	Y int
}
