package palindrome

import "sort"

// Sequence contains a string as a rune slice
type Sequence struct {
	w []rune
}

// IsPalinDrome check if it's a palindrome
func IsPalinDrome(s sort.Interface) bool {
	j := s.Len()
	for i := 0; i < s.Len()/2; i++ {
		j = j - 1
		if !(!s.Less(i, j) && !s.Less(j, i)) {
			return false
		}
	}
	return true
}

func (s *Sequence) Len() int {
	return len(s.w)
}

func (s *Sequence) Less(i, j int) bool {
	return s.w[i] < s.w[j]
}

func (s *Sequence) Swap(i, j int) {
	s.w[i], s.w[j] = s.w[j], s.w[i]
}
