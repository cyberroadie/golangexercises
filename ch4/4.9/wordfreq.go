package main

import (
	"bufio"
	"fmt"
	"os"
	"text/tabwriter"
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

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)

	for k, v := range wordfreq {
		fmt.Fprintf(w, "%s\t%d\n", k, v)
	}
	w.Flush()
}
