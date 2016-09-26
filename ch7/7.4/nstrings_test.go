package nstrings

import (
	"testing"

	"golang.org/x/net/html"
)

func TestReader(t *testing.T) {
	r := NewReader("<html><body>lalalala</body></html>")

	if r == nil {
		t.Error("reader is nil")
	}
	t.Log("parsing")

	n, err := html.Parse(r)
	if err != nil {
		t.Errorf("%s", err)
	}
	if n.FirstChild == nil {
		t.Error("first node not found")
	}

	if n.FirstChild.Data != "html" {
		t.Errorf("incorrect first child: %s", n.FirstChild.Data)
	}
}
