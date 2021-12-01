// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 6.
//!+

// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}

	// C:\Users\desktop-huntersudo\AppData\Local\Temp\___8go_build_gopl_io_ch1_echo2.exe
	// 执⾏命令本身的名字
	fmt.Println(os.Args[0])
	fmt.Println(s)
	fmt.Println(len(s))

	for index, value := range os.Args[1:] {
		fmt.Println(index, ":", value)
	}

}

//!-
