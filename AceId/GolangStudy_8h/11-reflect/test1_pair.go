package main

import "fmt"

func main() {

	var a string
	//todo pair<statictype:string, value:"aceld">
	// type 指针   value 指针
	a = "aceld"

	var allType interface{}
	allType = a
	// todo pair<type:string, value:"aceld">
	// todo 不管如何复制 ，type 指针 ，value指针会一直传递下去

	str, _ := allType.(string)  // 断言：找到具体的类型、
	fmt.Println(str)
}
