package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("%s\n", expand("test test test $foo test", func(r string) string { return "done" }))
}

func expand(s string, f func(string) string) string {
	return strings.Replace(s, "$foo", f("foo"), 1)
}
