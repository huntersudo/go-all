package main

import (
	"errors"
	"fmt"
	"log"
	"runtime"
)

// 剖析异常
// recover 函数调用的返回 值和 panic 函数的输入参数类型一致
//func panic(interface{})
//func recover() interface{}

func main() {

	if r := recover(); r != nil {
		log.Fatal(r)
	}
	panic(123)

	if r := recover(); r != nil {
		log.Fatal(r)
	}

}

// 上面程序中两个 recover 调用都不能捕获任何异常。
// 在第一个 recover 调用执行时，函数必然是 在正常的非异常执行流程中，这时候 recover 调用将返回 nil 。
// 发生异常时，第二 个 recover 调用将没有机会被执行到，todo 因为 panic 调用会导致函数马上执行已经注册 defer 的 函数后返回

func main1() {
	defer func() {
		// 无法捕获异常
		if r := MyRecover(); r != nil {
			fmt.Println(r)
		}
	}()
}
func MyRecover() interface{} {
	log.Println("trace...")
	return recover()
}

// todo 我们必须在 defer 函数中直接调用 recover 。
// 如 果 defer 中调用的是 recover 函数的包装函数的话，异常的捕获工作将失败

// todo 是在嵌套的 defer 函数中调用 recover 也将导致无法捕获异常：
// 2层嵌套的 defer 函数中直接调用 recover 和1层 defer 函数中调用包装的 MyRecover 函数 一样，都是经过了2个函数帧才到达真正的 recover 函数，
// todo 这个时候Goroutine的对应上一级栈帧中 已经没有异常信息。
func main2() {
	defer func() {

		defer func() {
			// 无法捕获异常
			if r := recover(); r != nil {
				fmt.Println(r)
			}
		}()
	}()
}

//todo 必须要和有异常的栈帧只隔一个栈帧， recover 函数才能正常捕获异常。换言之， recover 函数 捕获的是祖父一级调用函数栈帧的异常（刚好可以跨越一层 defer 函数）！

func foo() (err error) {
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case runtime.Error:
				// // 这是运行时错误类型异常
			case error:
				err = x
				// // 普通错误类型异常
			default:
				err = fmt.Errorf("Unknow Panic: %v", r)
			}
		}

	}()

}
