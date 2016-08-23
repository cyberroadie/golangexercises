package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"hash"
	"os"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	hashSize := flag.Int("hashSize", 256, "use 256/384/512")
	flag.Parse()

	hw1 := os.Args[len(os.Args)-2]
	hw2 := os.Args[len(os.Args)-1]

	if hw1 == "" || hw2 == "" {
		fmt.Fprintln(os.Stderr, "program needs 2 string inputs")
		os.Exit(1)
	}

	h1 := createHash(hw1, *hashSize)
	h2 := createHash(hw2, *hashSize)

	bd := byteArrayDiff(h1, h2)

	fmt.Printf("The hashes of %s and %s have %d different bits\n", hw1, hw2, bd)
	byte2Bits(h1)
	byte2Bits(h2)
}

func byte2Bits(s []byte) {
	for _, c := range s {
		fmt.Printf("%b", c)
	}
	fmt.Println("")
}

func byteArrayDiff(h1, h2 []byte) (bd int64) {
	for i, b1 := range h1 {
		if i == len(h2)-1 {
			return
		}
		bd += int64(bitDiff(b1, h2[i]))
	}
	return
}

func bitDiff(b1, b2 byte) int {
	return int(pc[b1&^b2])
}

func createHash(b string, size int) []byte {
	var hasher hash.Hash
	switch size {
	case 256:
		hasher = sha256.New()
	case 384:
		hasher = sha512.New384()
	case 512:
		hasher = sha512.New()
	}

	hasher.Write([]byte(b))
	return hasher.Sum(nil)
}
