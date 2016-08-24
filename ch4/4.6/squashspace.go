package main

import (
	"fmt"
	"unicode"
)

func main() {
	lala := "this is a    unicode string"
	var byteSlice []byte
	for i, elem := range lala {
		if unicode.IsSpace(elem) {
			continue
		}
		if i > 0 && unicode.IsSpace(rune(lala[i-1])) {
			byteSlice = append(byteSlice, ' ')
		}
		byteSlice = append(byteSlice, byte(elem))
	}
	fmt.Println(string(byteSlice))
}
