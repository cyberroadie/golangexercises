package wordcounter

import (
	"bufio"
	"strings"
)

// WordCounter counts words in a string
type WordCounter int

func (c *WordCounter) Write(input string) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		return -1, err
	}
	return count, nil
}
