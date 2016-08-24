package main

import (
	"fmt"
)

func main() {
	var pc [256]byte
	for i := range pc {
		pc[i] = pc[i>>1] + byte(i&1)
		// equals pc[i] = pc[i/2] + byte(i&1)
		// so the bitshift to the right has the same number of bits + the AND operation
		// of 1 on the current number to see if the rightmost bit is extra or not
	}

	for i := range pc {
		fmt.Printf("%d[%d]%.8b[%.8b]: %.4b + %.4b = %.4b (%d)\n", i, i/2, i, i/2, pc[i/2], i&1, pc[i], pc[i])
	}
}
