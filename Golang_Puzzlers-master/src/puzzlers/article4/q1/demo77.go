package main

import (
	"flag"
	"fmt"
)

func main() {
	//var name string                                                   // [1]
	//flag.StringVar(&name, "name", "everyone", "The greeting object.") // [2]

	// 方式1。
	// flag.String函数返回的结果值的类型是*string而不是string。类型*string代表的是字符串的指针类型，而不是字符串类型。
	// 因此，这里的变量name代表的是一个指向字符串值的指针。
	var name = flag.String("name", "everyone", "The greeting object.")

	// 方式2。
	// 短变量声明
	//name := flag.String("name", "everyone", "The greeting object.")

	flag.Parse()
	//fmt.Printf("Hello, %v!\n", name)

	// 适用于方式1和方式2。
	// 通过操作符*把这个指针指向的字符串值取出来了
	fmt.Printf("Hello, %v!\n", *name)
}
