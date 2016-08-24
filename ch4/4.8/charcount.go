// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 97.
//!+

// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	// unicode.IsDigit()
	// unicode.IsLetter()
	// unicode.IsSpace()

	countsDigits := make(map[rune]int)  // counts of Unicode digits
	countsLetters := make(map[rune]int) // counts of Unicode letters

	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		if unicode.IsDigit(r) {
			countsDigits[r]++
		} else if unicode.IsLetter(r) {
			countsLetters[r]++
		}

		utflen[n]++
	}
	fmt.Printf("letter rune\tcount\n")
	for c, n := range countsLetters {
		fmt.Printf("%q\t%d\n", c, n)
	}

	fmt.Printf("digit rune\tcount\n")
	for c, n := range countsDigits {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

//!-