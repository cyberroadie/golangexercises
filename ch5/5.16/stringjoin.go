package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%s\n", join("hello ", "world", "!"))
}

func join(vals ...string) (result string) {
	for _, s := range vals {
		result = result + s
	}
	return
}
