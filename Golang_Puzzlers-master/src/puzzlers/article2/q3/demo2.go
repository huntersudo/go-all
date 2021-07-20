package main

import (
	"flag"
	"fmt"
	"os"
)

var name2 string

func init() {
	// 方式2。
	flag.CommandLine = flag.NewFlagSet("", flag.ExitOnError)

	flag.CommandLine.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
		flag.PrintDefaults()
	}
	flag.StringVar(&name2, "name2", "everyone", "The greeting object.")
}

func main() {
	flag.Parse()
	fmt.Printf("Hello, %s!\n", name2)
}
// go run xx  --help
/**
Usage of question:
  -name2 string
    	The greeting object. (default "everyone")
 */