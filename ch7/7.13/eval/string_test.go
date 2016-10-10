package eval

import (
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	tests := []struct {
		expr string
		want string
	}{
		{"sqrt(A / pi)", "sqrt(A / pi)"},
		{"pow(x, 3) + pow(y, 3)", "pow(x, 3) + pow(y, 3)"},
	}

	for _, test := range tests {
		expr, err := Parse(test.expr)
		if err != nil {
			t.Error(err) // parse error
			continue
		}
		t.Logf("is: %s \nwant: %s", expr.String(), test.want)
		if strings.Compare(expr.String(), test.want) != 0 {
			t.Error(err) // parse error
		}

		previousExpr := expr.String()
		expr, err = Parse(previousExpr)
		if err != nil {
			t.Error(err) // parse error
			continue
		}
		if strings.Compare(expr.String(), previousExpr) != 0 {
			t.Error(err) // parse error
		}

	}
}
