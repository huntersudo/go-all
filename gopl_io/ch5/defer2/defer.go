// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 151.

// Defer2 demonstrates a deferred call to runtime.Stack during a panic.
package main

import (
	"fmt"
	"os"
	"runtime"
)

//!+
func main() {
	defer printStack()
	f(3)
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}

//!-

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // panics if x == 0
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}

/*
//!+printstack
goroutine 1 [running]:
server.printStack()
	server/gopl.io--/ch5/defer2/defer_close.go:20
server.f(0)
	server/gopl.io--/ch5/defer2/defer_close.go:27
server.f(1)
	server/gopl.io--/ch5/defer2/defer_close.go:29
server.f(2)
	server/gopl.io--/ch5/defer2/defer_close.go:29
server.f(3)
	server/gopl.io--/ch5/defer2/defer_close.go:29
server.server()
	server/gopl.io--/ch5/defer2/defer_close.go:15
//!-printstack
*/
