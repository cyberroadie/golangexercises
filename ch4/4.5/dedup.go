package main

import (
	"fmt"
)

func main() {
	s := []int{1, 1, 2, 2, 3, 4, 4, 5, 5}

	for i, elem := range s {
		println(i)
		if i > 0 && i < len(s) && elem == s[i-1] {
			s = append(s[:i], s[i+1:]...)
		}
	}
	fmt.Println(s)
}
