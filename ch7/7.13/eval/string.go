package eval

import "fmt"

// // An Expr is an arithmetic expression.
// type Expr interface {
// 	// Eval returns the value of this Expr in the environment env.
// 	Eval(env Env) float64
// 	// Check reports errors in this Expr and adds its Vars to the set.
// 	Check(vars map[Var]bool) error

// 	// String pretty print tree
// 	String() string
// }

// String pretty-prints a syntax tree
func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	return fmt.Sprintf("%f", l)
}

func (u unary) String() string {
	switch u.op {
	case '+':
		fmt.Print(" - ")
		return u.x.String()
	case '-':
		fmt.Print(" - ")
		return u.x.String()
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

func (b binary) String() string {
	return fmt.Sprintf("%s %c %s", b.x.String(), b.op, b.y.String())
}

func (c call) String() string {
	switch c.fn {
	case "pow":
		return fmt.Sprintf("pow(%s, %s)", c.args[0].String(), c.args[1].String())
	case "sin":
		return fmt.Sprintf("sin(%s)", c.args[0].String())
	case "sqrt":
		return fmt.Sprintf("sqrt(%s)", c.args[0].String())
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}
