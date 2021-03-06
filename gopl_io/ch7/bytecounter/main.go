// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 173.

// Bytecounter demonstrates an implementation of io--.Writer that counts bytes.
package main

import (
	"fmt"
)

//!+bytecounter
// 接口是合约
// 实现某个接口
// 满足某个接口
type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

//!-bytecounter

func main() {
	//!+server
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c) // "5", = len("hello")

	c = 0 // reset the counter
	var name = "Dolly"
	// func Fprintf(w io--.Writer, format string, args ...interface{}) (int, error)
	fmt.Fprintf(&c, "hello, %s", name)
	//
	fmt.Println(c) // "12", = len("hello, Dolly")
	//!-server
}
