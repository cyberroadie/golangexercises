// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 136.

// The toposort program prints the nodes of a DAG in topological order.
package main

import "fmt"
import "log"

//!+table
// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms":     {"data structures"},
	"calculus":       {"linear algebra"},
	"linear algebra": {"calculus"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

//!-table

//!+main
func main() {
	courses := topoSort(prereqs)

	for i := 0; i < len(courses); i++ {
		fmt.Printf("%d:\t%s\n", i+1, courses[i])
	}
}

func topoSort(m map[string][]string) map[int]string {
	order := make(map[int]string)
	seen := make(map[string]bool)
	var visitAll func(items []string)

	visitAll = func(items []string) {
		for _, item := range items {
			if checkForCycle(item) {
				log.Fatal("Resolve cycle before continuing")
			}

			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order[len(order)] = item
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	visitAll(keys)
	return order
}

func checkForCycle(course string) bool {
	pr1 := prereqs[course]
	if pr1 == nil {
		return false
	}

	for _, pc1 := range pr1 {
		pr2 := prereqs[pc1]
		if pr2 == nil {
			continue
		}
		for _, pc2 := range pr2 {
			if pc2 == course {
				fmt.Printf("\tCycle detected between '%s' and '%s'\n", pc1, course)
				return true
			}
		}

	}

	return false
}

//!-main
