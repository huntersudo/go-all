package main

import (
	"bytes"
	"container/list"
	"fmt"
)

func main() {
	l := list.New()

	// 尾部添加
	l.PushBack("canon")
	print(l, "PushBack")

	// 头部添加
	l.PushFront(67)
	print(l, "PushFront")

	// 尾部添加后保存元素句柄
	element := l.PushBack("fist")
	// 在“fist”之后添加”high”
	l.InsertAfter("high", element)
	print(l, "InsertAfter")

	// 在“fist”之前添加”noon”
	l.InsertBefore("noon", element)
	print(l, "InsertBefore")

	// 使用
	l.Remove(element)
	print(l, "Remove")
}

func print(l *list.List, msg string) {
	fmt.Printf("==%s==", msg)
	fmt.Println()
	var buffer bytes.Buffer
	for i := l.Front(); i != nil; i = i.Next() {
		//fmt.Println(i.Value)
		//// 将interface{}类型格式化为字符串
		buffer.WriteString(fmt.Sprintf("%v", i.Value))
		buffer.WriteString(",")

	}
	fmt.Println(buffer.String())
	fmt.Println()
}
