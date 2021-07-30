package main

import "fmt"



func returAndDefer() int {
	defer deferFunc()
	return returnFunc()
}

func deferFunc() int {
	fmt.Println("deferFunc()")
	return 1
}

func returnFunc() int  {
	fmt.Println("returnFunc()")
	return 0
}

func main() {
	returAndDefer()   // todo return  先执行  defer在当前func声明周期之后才出栈执行
}
//returnFunc()
//deferFunc()