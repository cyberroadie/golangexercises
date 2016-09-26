package nio

import (
	"strings"
	"testing"
)

func TestLimitReader(t *testing.T) {
	r := strings.NewReader("01234567890")
	lr := LimitReader(r, 4)
	var p = make([]byte, 1024)
	n, err := lr.Read(p)
	if err != nil {
		t.Errorf("%s", err)
	}

	if n != 4 {
		t.Errorf("to many bytes read: %d", n)
	}

	for i := 4; i < len(p); i++ {
		if p[i] != 0 {
			t.Errorf("to many bytes written: %v", p)
		}
	}
}
