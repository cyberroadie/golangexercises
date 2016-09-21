// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 165.

// Package intset provides a set of integers based on a bit vector.
package intset

import (
	"bytes"
	"fmt"
)

//!+intset

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// IntersectWith sets s to the intersection of s and t
func (s *IntSet) IntersectWith(t *IntSet) {
	for word := range s.words {
		if word > len(t.words) {
			s.words[word] = 0
			continue
		}
		s.words[word] = s.words[word] & t.words[word]
	}
}

// DifferenceWith sets s to the difference of s and t
func (s *IntSet) DifferenceWith(t *IntSet) {
	for tword, tbits := range t.words {
		if tword > len(s.words) {
			s.words = append(s.words, tbits)
			continue
		}
		s.words[tword] = s.words[tword] &^ tbits
	}
}

// SymetricDifferenceWith sets s to the symetric difference of s and t
func (s *IntSet) SymetricDifferenceWith(t *IntSet) {
	for tword, tbits := range t.words {
		if tword > len(s.words) {
			s.words = append(s.words, tbits)
			continue
		}
		s.words[tword] = s.words[tword] ^ tbits
	}
}

//!-intset

// AddAll adds all given integers
func (s *IntSet) AddAll(vars ...int) {
	for _, v := range vars {
		s.Add(v)
	}
}

// Len returns number of elements
func (s *IntSet) Len() (count int) {
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				count++
			}
		}
	}
	return
}

// Remove removes x from the set
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	s.words[word] = s.words[word] &^ (1 << bit)
}

// Clear remove all elements from the set
func (s *IntSet) Clear() {
	for i := range s.words {
		s.words[i] = 0
	}
}

// Copy returns a copy of the set
func (s *IntSet) Copy() *IntSet {
	var t IntSet
	for _, v := range s.words {
		t.words = append(t.words, v)
	}
	return &t
}

// Elems returns an iterable slice
func (s *IntSet) Elems() (elems []int) {
	for word, bits := range s.words {
		if bits == 0 {
			continue
		}
		for i := 0; i < 64; i++ {
			if bits&(1<<uint(i)) != 0 {
				elems = append(elems, 64*word+i)
			}
		}
	}
	return elems
}

//!+string

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

//!-string
