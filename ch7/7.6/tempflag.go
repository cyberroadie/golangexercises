// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 181.

// Tempflag prints the value of its -cel (celcius temperature) and -far (fahrenheit) flag.
package main

import (
	"flag"
	"fmt"

	"./tempconv"
)

//!+
var tempC = tempconv.CelsiusFlag("cel", 20.0, "the temperature")
var tempF = tempconv.FahrenheitFlag("far", 20)

func main() {
	flag.Parse()
	fmt.Println(*tempC)
	fmt.Println(*tempF)
}

//!-
