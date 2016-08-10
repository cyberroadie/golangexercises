// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package popcount_test

import (
	"testing"

	"github.com/cyberroadie/golangexercises/ch2/2.3/popcount"
)

// -- Alternative implementations --

func BitCountLoop(x uint64) int {
	// Hacker's Delight, Figure 5-2.
	x = x - ((x >> 1) & 0x5555555555555555)
	x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
	x = (x + (x >> 4)) & 0x0f0f0f0f0f0f0f0f
	x = x + (x >> 8)
	x = x + (x >> 16)
	x = x + (x >> 32)
	return int(x & 0x7f)
}

func PopCountLoopByClearing(x uint64) int {
	n := 0
	for x != 0 {
		x = x & (x - 1) // clear rightmost non-zero bit
		n++
	}
	return n
}

func PopCountLoopByShifting(x uint64) int {
	n := 0
	for i := uint(0); i < 64; i++ {
		if x&(1<<i) != 0 {
			n++
		}
	}
	return n
}

// -- Benchmarks --

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountLoop(0x1234567890ABCDEF)
	}
}

func BenchmarkBitCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BitCountLoop(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountLoopByClearing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLoopByClearing(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountLoopByShifting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLoopByShifting(0x1234567890ABCDEF)
	}
}

// 2.67GHz Xeon
// $ go test -cpu=4 -bench=. .
// BenchmarkPopCount-4                  200000000         6.30 ns/op
// BenchmarkBitCount-4                  300000000         4.15 ns/op
// BenchmarkPopCountByClearing-4        30000000         45.2 ns/op
// BenchmarkPopCountByShifting-4        10000000        153 ns/op
//
// 2.5GHz Intel Core i5
// $ go test -cpu=4 -bench=. .
// testing: warning: no tests to run
// BenchmarkPopCount-4                  200000000         7.52 ns/op
// BenchmarkBitCount-4                  500000000         3.36 ns/op
// BenchmarkPopCountByClearing-4        50000000         34.3 ns/op
// BenchmarkPopCountByShifting-4        20000000        108 ns/op