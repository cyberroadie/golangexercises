package main

import (
	"fmt"
)

func main() {
	s := []int{0, 1, 2, 3, 4, 5}

	s = rotate(s, 3)
	fmt.Println(s)

}

func rotate(s []int, times int) []int {
	return append(s[times:len(s)], s[0:times]...)
}
