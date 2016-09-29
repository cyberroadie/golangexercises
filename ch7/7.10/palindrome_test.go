package palindrome

import "testing"

func TestIsPalinDrome(t *testing.T) {
	seq := Sequence{[]rune("lalaalal")}

	if !IsPalinDrome(&seq) {
		t.Errorf("%s should be a palindrome", string(seq.w))
	}
}

func TestNotPalinDrome(t *testing.T) {
	seq := Sequence{[]rune("palindrome")}

	if IsPalinDrome(&seq) {
		t.Errorf("%s should not be a palindrome", string(seq.w))
	}
}
