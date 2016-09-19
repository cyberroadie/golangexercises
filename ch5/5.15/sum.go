package main

import (
	"fmt"
)

func main() {
	values := []int{1, 2, 3, 4, 5, 4, 3, 2, 1}

	fmt.Printf("max %d\n", max(values...))
	fmt.Printf("min %d\n", min(values...))
	fmt.Printf("max1 %d\n", max(values...))
	fmt.Printf("min1 %d\n", max(values...))

}

const maxInt = int(^uint(0) >> 1)
const minInt = -maxInt - 1

func max(vals ...int) (max int) {
	if len(vals) == 0 {
		panic("max() needs at least one value")
	}
	max = minInt
	for _, v := range vals {
		if v > max {
			max = v
		}
	}
	return
}

func min(vals ...int) (min int) {
	if len(vals) == 0 {
		panic("min() needs at least one value")
	}
	min = maxInt
	for _, v := range vals {
		if v < min {
			min = v
		}
	}
	return
}

func max1(max int, vals ...int) int {
	for _, v := range vals {
		if v > max {
			max = v
		}
	}
	return max
}

func min1(min int, vals ...int) int {
	for _, v := range vals {
		if v < min {
			min = v
		}
	}
	return min
}
