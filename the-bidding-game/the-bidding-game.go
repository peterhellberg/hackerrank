package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func (s state) move() int {
	if s.round == 0 {
		return 16
	}

	if s.distance == 1 && s.playerBalance >= s.enemyBalance {
		return s.playerBalance
	}

	if s.distance == 9 {
		if s.enemyBalance < s.playerBalance {
			return s.enemyBalance
		}

		return s.playerBalance
	}

	m := s.playerBalance / s.distance

	if m > s.playerBalance {
		m = s.playerBalance
	}

	if m < 1 {
		return 1
	}

	return m
}

func main() {
	fmt.Printf("%d\n", makeState(os.Stdin).move())
}

type state struct {
	player        int
	playerBalance int
	playerAverage int
	playerLastBet int
	enemy         int
	enemyBalance  int
	enemyAverage  int
	enemyLastBet  int
	draws         int
	scotch        int
	round         int
	distance      int
	moves         map[int][]int
}

func makeState(r io.Reader) state {
	s := readState(r)

	if s.player == 1 {
		s.enemy = 2
	} else {
		s.enemy = 1
	}

	pb, eb := 100, 100

	var draws int

	for i, pm := range s.moves[s.player] {
		em := s.moves[s.enemy][i]

		switch {
		case em == pm:
			debug("DRAW   p:%d == e:%d", pm, em)
			draws++

			if draws%2 == 0 {
				eb -= em
			} else {
				pb -= pm
			}
		case em > pm:
			debug("ENEMY  p:%d < e:%d", pm, em)
			eb -= em
		case pm > em:
			debug("PLAYER p:%d > e:%d", pm, em)
			pb -= pm
		}

		s.playerLastBet = pm
		s.enemyLastBet = em
	}

	if s.player == 1 {
		s.distance = s.scotch
	} else {
		s.distance = 12 - s.scotch
	}

	s.round = len(s.moves[1])
	s.draws = draws

	s.playerBalance = pb
	s.enemyBalance = eb

	if s.round > 0 {
		s.playerAverage = (100 - s.playerBalance) / s.round
		s.enemyAverage = (100 - s.enemyBalance) / s.round
	}

	if os.Getenv("DEBUG") == "true" {
		debug("%+v", s)
	}

	if s.playerBalance < 1 || s.enemyBalance < 1 {
		debug("END!")
	}

	return s
}

func readState(r io.Reader) state {
	scanner := bufio.NewScanner(r)

	s := state{moves: map[int][]int{}}

	if scanner.Scan() {
		s.player, _ = strconv.Atoi(scanner.Text())
	}

	if scanner.Scan() {
		s.scotch, _ = strconv.Atoi(scanner.Text())
	}

	if scanner.Scan() {
		s.moves[1] = movesFromString(scanner.Text())
	}

	if scanner.Scan() {
		s.moves[2] = movesFromString(scanner.Text())
	}

	return s
}

func movesFromString(s string) []int {
	moves := []int{}

	for _, v := range strings.Split(s, " ") {
		if m, err := strconv.Atoi(v); err == nil {
			moves = append(moves, m)
		}
	}

	return moves
}

func debug(format string, v ...interface{}) {
	fmt.Fprintf(os.Stderr, format+"\n", v...)
}
