// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 101.

// Package treesort provides insertion sort using an unbalanced binary tree.
package treesort2

import "bytes"
import "fmt"

//!+
type tree struct {
	value       int
	left, right *tree
}

// Sort sorts values in place.
func Sort(values []int) string {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
	return root.String()
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

// String print sequence of tree
func (t *tree) String() string {
	var buffer bytes.Buffer
	traverseInOrder(&buffer, t)
	return buffer.String()
}

func traverseInOrder(buffer *bytes.Buffer, t *tree) {
	if t == nil {
		return
	}
	traverseInOrder(buffer, t.left)
	buffer.WriteString(fmt.Sprintf("%d ", t.value))
	traverseInOrder(buffer, t.right)
}

//!-
