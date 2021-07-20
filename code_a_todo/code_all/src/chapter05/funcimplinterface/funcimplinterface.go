package main

import (
	"fmt"
)

/**
todo:多看看
把函数作为接口来调用
*/
// 调用器接口
type Invoker interface {
	// 需要实现一个Call方法
	Call(interface{})
}

// 结构体类型
type Struct struct {
}

//定义结构体，该例子中的结构体无须任何成员，主要展示实现 Invoker 的方法。
// 实现Invoker的Call
func (s *Struct) Call(p interface{}) {
	fmt.Println("from struct", p)
}

/**
函数的声明不能直接实现接口，需要将函数定义为类型后，使用类型实现结构体，当类型方法被调用时，还需要调用函数本体。
FuncCaller 的 Call() 方法被调用与 func(interface{}) 无关，还需要手动调用函数本体。
*/
// 函数定义为类型
type FuncCaller func(interface{})

// 实现Invoker的Call
func (f FuncCaller) Call(p interface{}) {

	// 调用f函数本体
	f(p)
}

func main() {

	/**
	  将定义的 Struct 类型实例化，并传入接口中进行调用
	s 类型为 *Struct，已经实现了 Invoker 接口类型，因此赋值给 invoker 时是成功的。
	*/
	// 声明接口变量
	var invoker Invoker
	// 实例化结构体
	s := new(Struct)
	// 将实例化的结构体赋值到接口
	invoker = s
	// 使用接口调用实例化结构体的方法Struct.Call
	invoker.Call("hello")

	/**
	上面代码只是定义了函数类型，需要函数本身进行逻辑处理，FuncCaller 无须被实例化，只需要将函数转换为 FuncCaller 类型即可，函数来源可以是命名函数、匿名函数或闭包，参见下面代码：
	*/
	// 将匿名函数转为FuncCaller类型，再赋值给接口
	invoker = FuncCaller(func(v interface{}) {
		fmt.Println("from function", v)
	})

	// 使用接口调用FuncCaller.Call，内部会调用函数本体
	invoker.Call("hello")
}
