package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
)

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

func (g *Game) FoundPrincess() bool {
	if g.Bot == g.Princess {
		return true
	} else {
		return false
	}
}

func main() {
	game := setupGame()

	for {
		if game.FoundPrincess() {
			break
		}

		fmt.Println(game.NextMove())
	}
}

func (g *Game) NextMove() string {
	if g.Bot.Y == g.Princess.Y {
		if g.Bot.X > g.Princess.X {
			g.Bot.X--
			return "LEFT"
		} else {
			g.Bot.X++
			return "RIGHT"
		}
	}

	if g.Bot.X == g.Princess.X {
		if g.Bot.Y > g.Princess.Y {
			g.Bot.Y--
			return "UP"
		} else {
			g.Bot.Y++
			return "DOWN"
		}
	}

	if g.Bot.Y > g.Princess.Y {
		g.Bot.Y--
		return "UP"
	} else {
		g.Bot.Y++
		return "DOWN"
	}

	if g.Bot.X > g.Princess.X {
		g.Bot.X--
		return "LEFT"
	} else {
		g.Bot.X++
		return "RIGHT"
	}
}

func setupGame() Game {
	size, bot, princess, board := getSizeBotPrincessAndBoard()

	return Game{size, bot, princess, board}
}

func getSizeBotPrincessAndBoard() (int, Pos, Pos, [][]byte) {
	reader := bufio.NewReader(os.Stdin)

	size := getInt(nextLine(reader))

	c := (size - 1) / 2

	bot := Pos{c, c}
	princess := Pos{}

	var board [][]byte

	y := 0

	for {
		line, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}

		handleError(err)

		x := bytes.Index(line, []byte("p"))

		if x > -1 {
			princess.X = x
			princess.Y = y
		}

		board = append(board, line)

		y++
	}

	return size, bot, princess, board
}

func getInt(bytes []byte) int {
	i, err := strconv.Atoi(string(bytes))

	handleError(err)

	return i
}

func nextLine(reader *bufio.Reader) []byte {
	line, _, err := reader.ReadLine()

	handleError(err)

	return line
}

func handleError(err error) {
	if err != nil {
		os.Exit(1)
	}
}
