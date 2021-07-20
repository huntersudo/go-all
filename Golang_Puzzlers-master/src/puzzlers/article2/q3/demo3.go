package main

import (
	"flag"
	"fmt"
	"os"
)

var name3 string

/**
我们索性不用全局的flag.CommandLine变量，转而自己创建一个私有的命令参数容器。我们在函数外再添加一个变量声明
 */
// 方式3。
var cmdLine = flag.NewFlagSet("question", flag.ExitOnError)

func init() {

	// 方式3。
	cmdLine.StringVar(&name3, "name3", "everyone", "The greeting object(方式3).")
}

func main() {

	// 方式3。
	cmdLine.Parse(os.Args[1:])
	fmt.Printf("Hello, %s!\n", name3)
}

/**
然后，我们把对flag.StringVar的调用替换为对cmdLine.StringVar调用，再把flag.Parse()替换为cmdLine.Parse(os.Args[1:])。

 */

// // go run xx  --help
/**
Usage of question:
  -name3 string
    	The greeting object(方式3). (default "everyone")

 */
