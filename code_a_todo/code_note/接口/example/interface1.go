// 接口的定义及使用

package main

import (
	"fmt"
)

type A1 interface {
	Say()
}

type B1 struct {
}

func (b B1) Say() {
	fmt.Println("B1 Say().....")
}

func main() {

	var b B1
	var a A1 = b
	a.Say()
	// output:
	// Say().....
}
