// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package treesort2

import (
	"math/rand"
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	data := make([]int, 50)
	for i := range data {
		data[i] = rand.Int() % 50
	}
	Sort(data)
	if !sort.IntsAreSorted(data) {
		t.Errorf("not sorted: %v", data)
	}
}

func TestStringSort(t *testing.T) {
	data := make([]int, 10)
	for i := 9; i >= 0; i-- {
		data[i] = i
	}
	Sort(data)

	expected := "0 1 2 3 4 5 6 7 8 9 "
	result := Sort(data)
	if result != expected {
		t.Errorf("result is deiffent from expexted '%s' - '%s'", result, expected)
	}

}
