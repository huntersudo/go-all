package main

import (
	"flag"
	"fmt"
	"os"
)

var name1 string

func init() {

}

func main() {
	// 方式1。
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question(方式1)")
		flag.PrintDefaults()
	}
	flag.Parse()
	fmt.Printf("Hello, %s!\n", name1)
}

// go run xx  --help
/**

Usage of question(方式1):

 */