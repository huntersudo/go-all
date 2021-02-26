// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 181.

// Tempflag prints the value of its -temp (temperature) flag.
package main

import (
	"flag"
	"fmt"

	"gopl.io/ch7/tempconv"
)

//!+
var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}

//!-

// ./tempflag -temp -18C
// ./tempflag -temp 273.15K
// ./tempflag -temp 212°F
/**
$ ./tempflag -temp 273.15K
invalid value "273.15K" for flag -temp: invalid temperature "273.15K"
Usage of C:\Users\desktop-huntersudo\AppData\Local\Temp\___go_build_gopl_io_ch7_tempflag.exe:
  -temp value
    	the temperature (default 20°C)

*/
