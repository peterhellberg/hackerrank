package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

var alphabet = strings.Split("ABCDEFGHIJKLMNOPQRSTUVWXYZ", "")

type message struct {
	keyword    string
	ciphertext string
}

func makeMessages(r io.Reader) []message {
	var n int

	fmt.Scan(&n)

	messages := make([]message, n)

	s := bufio.NewScanner(r)

	for i := range messages {
		s.Scan()
		messages[i].keyword = s.Text()

		s.Scan()
		messages[i].ciphertext = s.Text()
	}

	return messages
}

func (m message) String() string {
	var key, letters, chunk string

	for _, k := range strings.Split(m.keyword, "") {
		if !strings.Contains(key, k) {
			key += k
		}
	}

	for _, l := range alphabet {
		if !strings.Contains(key, l) {
			letters += l
		}
	}

	sk := strings.Split(key, "")

	cols := make([][]string, len(sk))

	for _, l := range strings.Split(key+letters, "") {
		chunk += l

		if len(chunk) == len(key) {
			for i, c := range strings.Split(chunk, "") {
				cols[i] = append(cols[i], c)
			}
			chunk = ""
		}
	}

	if len(chunk) > 0 {
		for i, c := range strings.Split(chunk, "") {
			cols[i] = append(cols[i], c)
		}
		chunk = ""
	}

	sort.Strings(sk)

	var subs, oldnews []string

	for _, k := range sk {
		for _, c := range cols {
			if c[0] == k {
				subs = append(subs, c...)
			}
		}
	}

	for i, l := range alphabet {
		oldnews = append(oldnews, subs[i], l)
	}

	return strings.NewReplacer(oldnews...).Replace(m.ciphertext)
}

func main() {
	for _, m := range makeMessages(os.Stdin) {
		fmt.Println(m)
	}
}
