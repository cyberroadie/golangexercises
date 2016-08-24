package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// initialize map
	var wordfreq = make(map[string]int)

	// read file
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		wordfreq[scanner.Text()]++
	}

	for k, v := range wordfreq {
		fmt.Printf("%s\t\t\t\t%d\n", k, v)
	}

}
