// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 136.

// The toposort program prints the nodes of a DAG in topological order.
package main

import "fmt"

//!+table
// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string]map[int]string{
	"algorithms": {1: "data structures"},
	"calculus":   {1: "linear algebra"},

	"compilers": {
		1: "data structures",
		2: "formal languages",
		3: "computer organization",
	},

	"data structures":       {1: "discrete math"},
	"databases":             {1: "data structures"},
	"discrete math":         {1: "intro to programming"},
	"formal languages":      {1: "discrete math"},
	"networks":              {1: "operating systems"},
	"operating systems":     {1: "data structures", 2: "computer organization"},
	"programming languages": {1: "data structures", 2: "computer organization"},
}

//!-table

//!+main
func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string]map[int]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items map[int]string)

	visitAll = func(items map[int]string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	keys := make(map[int]string)
	i := 0
	for key := range m {
		keys[i] = key
		i++
	}

	visitAll(keys)
	return order
}

//!-main
