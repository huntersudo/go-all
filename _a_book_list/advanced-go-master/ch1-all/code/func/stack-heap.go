package main

import "fmt"

// 第一个函数直接返回了函数参数变量的地址——这似乎是不可以的，因为如果参数变量在栈上的话，函数 返回之后栈变量就失效了，返回的地址自然也应该失效了。
// 但是Go语言的编译器和运行时比我们聪明的 多，它会保证指针指向的变量在合适的地方。
func foo(x int) *int {
	return &x
}

func bar() *int {
	x := 6
	return &x;
}

// 内部虽然调用 new 函数创建了 *int 类 型的指针对象，但是依然不知道它具体保存在哪里。对于有C/C++编程经验的程序员需要强调的是：
// 不用关心Go语言中函数栈和堆的问题，编译器和运行时会帮我们搞定；同样不要假设变量在内存中的位置 是固定不变的，指针随时可能会变化，特别是在你不期望它变化的时候
func bar2() *int {
	var x = new(int)
	return x;
}

func main() {
	fmt.Println("Stack and Heap in Go")

	x := 5
	fmt.Printf("*(%#v) = %d\n", foo(x), *(foo(x)))
	fmt.Printf("*(%#v) = %d\n", bar(), *(bar()))
	fmt.Printf("*(%#v) = %d\n", bar2(), *(bar2()))
	//*((*int)(0xc00000a0b8)) = 5
	//*((*int)(0xc00000a0e8)) = 6
	//*((*int)(0xc00000a0f8)) = 0
}
// 我们无法知道函数参数或局部变量到底是保存在栈中还是 堆中，我们只需要知道它们能够正常工作就可以了
