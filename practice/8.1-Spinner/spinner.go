package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner()
	time.Sleep(10 * time.Second)
}

func spinner() {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(100 * time.Millisecond)
		}
	}
}
