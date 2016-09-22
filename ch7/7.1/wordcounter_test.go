package wordcounter

import (
	"testing"
)

func TestWrite(t *testing.T) {
	var wc WordCounter
	result, err := wc.Write("Lorem ipsum dolor sit amet, novum eripuit volumus nec eu.")
	if err != nil {
		t.Errorf("Error executing method Write %s", err)
	}
	expected := 10
	if result != expected {
		t.Errorf("wrong number of words counted %d instead of %d", result, expected)
	}
}
