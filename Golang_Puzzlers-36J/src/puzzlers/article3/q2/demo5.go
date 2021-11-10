package main

import (
	"flag"
	//idea  enable go mod in
	aa "puzzlers/src/puzzlers/article3/q2/lib"
	//"puzzlers/article3/q2/lib"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Parse()
	aa.Hello(name)
}
