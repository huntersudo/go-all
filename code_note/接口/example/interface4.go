// 接口是引用类型
// 空接口返回nil

package main

import (
	"fmt"
)

type A4 interface {
}

func main() {
	var a A4

	fmt.Println(a) // output: <nil>
}
