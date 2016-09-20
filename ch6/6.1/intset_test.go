// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package intset

import "testing"

func TestOne(t *testing.T) {
	//!+main
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	t.Log(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	t.Log(y.String()) // "{9 42}"

	x.UnionWith(&y)
	t.Log(x.String()) // "{1 9 42 144}"

	t.Log(x.Has(9), x.Has(123)) // "true false"
	//!-main

	// Output:
	// {1 9 144}
	// {9 42}
	// {1 9 42 144}
	// true false
}

func TestTwo(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	//!+note
	t.Log(&x)         // "{1 9 42 144}"
	t.Log(x.String()) // "{1 9 42 144}"
	t.Log(x)          // "{[4398046511618 0 65536]}"
	//!-note

	// Output:
	// {1 9 42 144}
	// {1 9 42 144}
	// {[4398046511618 0 65536]}
}

func TestThree(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	// Len
	if x.Len() != 4 {
		t.Errorf("Error testing Len(): %v != 4", &x)
	}

	// Remove
	x.Remove(144)
	if x.Has(144) {
		t.Error("Error testing Remove(): 144 not removed")
	}

	// Copy
	y := x.Copy()
	if y == nil || y.Len() != 3 || !y.Has(1) || !y.Has(9) || !y.Has(42) {
		t.Errorf("Error testing Copy(): %v != %v", y, x)
	}

	// Clear
	x.Clear()
	if x.Len() > 0 {
		t.Error("Error testing Clear(): InSet not cleared")
	}

}
