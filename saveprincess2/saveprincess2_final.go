package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strconv"
)

type Game struct {
	Size     int
	Bot      Bot
	Princess Princess
	Board    Board
}

type Bot struct {
	X int
	Y int
}

type Princess struct {
	X int
	Y int
}

type Board [][]byte

func (g *Game) NextMove() string {
	// Same line
	if g.Bot.Y == g.Princess.Y {
		if g.Bot.X > g.Princess.X {
			return "LEFT"
		} else {
			return "RIGHT"
		}
	}

	// Same column
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
	game := setupGame()

	os.Stdout.Write([]byte(game.NextMove()))
}

func setupGame() Game {
	reader := bufio.NewReader(os.Stdin)

	size := getInt(nextLine(reader))
	bot := getBot(nextLine(reader))

	board, princess := getBoardAndPrincess(reader)

	return Game{size, bot, princess, board}
}

func getInt(bytes []byte) int {
	i, err := strconv.Atoi(string(bytes))

	handleError(err)

	return i
}

func getBot(line []byte) Bot {
	pos := bytes.Split(line, []byte(" "))

	return Bot{getInt(pos[1]), getInt(pos[0])}
}

func getBoardAndPrincess(reader *bufio.Reader) (Board, Princess) {
	board := Board{}
	princess := Princess{}

	y := 0

	for {
		line, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}

		handleError(err)

		x := bytes.Index(line, []byte("p"))

		if x > -1 {
			// Found the princess!
			princess.X = x
			princess.Y = y
		}

		y++

		board = append(board, line)
	}

	return board, princess
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
