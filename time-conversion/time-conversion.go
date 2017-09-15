package main

import (
	"fmt"
	"time"
)

func main() {
	var value string

	fmt.Scan(&value)

	if t, err := time.Parse("3:04:05PM", value); err == nil {
		fmt.Println(t.Format("15:04:05"))
	}
}
