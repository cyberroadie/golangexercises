package main

import (
	"fmt"
	"os"
)

func main() {
	word1 := os.Args[1]
	word2 := os.Args[2]

	if word1 == "" || word2 == "" {
		fmt.Fprint(os.Stderr, "need two strings as input")
		os.Exit(1)
	}

	if isAnagram(word1, word2) {
		fmt.Fprintf(os.Stdout, "%s and %s are anagrams\n", word1, word2)
	} else {
		fmt.Fprintf(os.Stdout, "%s and %s are not anagrams\n", word1, word2)
	}
}

func isAnagram(w1, w2 string) bool {
	if len(w1) != len(w2) {
		return false
	}

	c1 := rune(w1[0])
	for i, c2 := range w2 {
		if c1 == c2 {
			if len(w2) == 1 {
				return true
			}
			if i == 0 {
				return isAnagram(w1[1:], w2[1:])
			} else if i == len(w2)-1 {
				return isAnagram(w1[1:], w2[0:i])
			} else if i == 1 {
				return isAnagram(w1[1:], string(w2[0])+w2[2:])
			}
			return isAnagram(w1[1:], w2[0:i-1]+w2[i+1:])
		}
	}

	return false
}
