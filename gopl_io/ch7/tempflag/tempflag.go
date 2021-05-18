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
/**
desktop-huntersudo@DESKTOP-C1 MINGW64 /d/workspace/ws1/src/github.com/go-all/gopl_io (master)
$ go build gopl.io/ch7/tempflag/

desktop-huntersudo@DESKTOP-C1 MINGW64 /d/workspace/ws1/src/github.com/go-all/gopl_io (master)
$ ls
CNAME  README.md  ch1  ch10  ch11  ch12  ch13  ch2  ch3  ch4  ch5  ch6  ch7  ch8  ch9  go.mod  tempflag.exe

desktop-huntersudo@DESKTOP-C1 MINGW64 /d/workspace/ws1/src/github.com/go-all/gopl_io (master)
$ ./tempflag.exe -temp -18C
-18°C

desktop-huntersudo@DESKTOP-C1 MINGW64 /d/workspace/ws1/src/github.com/go-all/gopl_io (master)
$ ./tempflag.exe -temp 211F
99.44444444444444°C

desktop-huntersudo@DESKTOP-C1 MINGW64 /d/workspace/ws1/src/github.com/go-all/gopl_io (master)
$ ./tempflag.exe -temp 212F
100°C


 */