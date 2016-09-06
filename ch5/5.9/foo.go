package main

import (
	"fmt"
	"strings"
)

func main() {

	f := func(r string) string {
		return "done"
	}

	fmt.Printf("%s\n", expand("test test test $foo test", f))
}

func expand(s string, f func(string) string) string {
	return strings.Replace(s, "$foo", f("foo"), 1)
}
