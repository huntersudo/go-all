package main

import (
	"fmt"
	"unsafe"
)

// 首先，我来介绍下 Slice，中文翻译叫“切片”，这个东西在 Go 语言中不是数组，而是一个结构体，其定义如下：

type slice struct {
	array unsafe.Pointer //指向存放数据的数组指针
	len   int            //长度有多大
	cap   int            //容量有多大
}
// 熟悉 C/C++ 的同学一定会知道在结构体里用数组指针的问题——数据会发生共享！下面我们来看看 Slice 的一些操作：
func  main(){
	foo := make([]int, 5)
	foo[3] = 42
	foo[4] = 100

	bar  := foo[1:4]
	bar[1] = 99
	fmt.Println(foo)
	fmt.Println(bar)
	// [0 0 99 42 100]  todo 可以发现foo[2] 也就是b[1] 从0 变为 99了
	//   [0 99 42]
}
/**
我来解释下这段代码：
首先，创建一个 foo 的 Slice，其中的长度和容量都是 5；
然后，开始对 foo 所指向的数组中的索引为 3 和 4 的元素进行赋值；
最后，对 foo 做切片后赋值给 bar，再修改 bar[1]。
todo 从这张图片中，我们可以看到，因为 foo 和 bar 的内存是共享的，所以，foo 和 bar 对数组内容的修改都会影响到对方。
 */




