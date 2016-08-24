// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 86.

// Rev reverses a slice.
package main

import (
	"fmt"
)

func main() {
	utf8string := "test 1 2 3"
	byteSlice := []byte(utf8string)

	fmt.Println(string(byteSlice))
	reverse(byteSlice)
	fmt.Println(string(byteSlice))

}

//!+rev
// reverse reverses a slice of ints in place.
func reverse(b []byte) {
	if len(b) > 1 {
		b[0], b[len(b)-1] = b[len(b)-1], b[0]
		reverse(b[1 : len(b)-1])
	}
}

//!-rev
