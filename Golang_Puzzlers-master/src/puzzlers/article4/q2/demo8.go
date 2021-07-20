package main

import (
	"flag"
	"fmt"
)

func main() {
	// 我们不显式地指定变量name的类型，使得它可以被赋予任何类型的值
	// 在你改变getTheFlag函数的结果类型之后，Go 语言的编译器会在你再次构建该程序的时候，自动地更新变量name的类型
	// 通过这种类型推断，你可以体验到动态类型编程语言所带来的一部分优势，即程序灵活性的明显提升
	var name = getTheFlag()
	flag.Parse()
	fmt.Printf("Hello, %v!\n", *name)
}

func getTheFlag() *string {
	return flag.String("name", "everyone", "The greeting object.")
}

//上面函数的实现也可以是这样的。
//func getTheFlag() *int {
//	return flag.Int("num", 1, "The number of greeting object.")
//}
