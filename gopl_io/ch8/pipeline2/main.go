// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 229.

// Pipeline2 demonstrates a finite 3-stage pipeline.
package main

import "fmt"

//!+
func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// Squarer
	go func() {
		// range ，它依次从channel接收数据，当channel被关闭并且没有值可接收时跳出循环
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	// Printer (in server goroutine)
	for x := range squares {
		fmt.Println(x)
	}
}

//!-
