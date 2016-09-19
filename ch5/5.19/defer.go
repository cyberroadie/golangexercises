package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%s\n", panicnoreturn())
}

func panicnoreturn() (txt string) {
	defer func() {
		if p := recover(); p != nil {
			txt = "hello world!"
		}
	}()
	panic("attack!")
}
