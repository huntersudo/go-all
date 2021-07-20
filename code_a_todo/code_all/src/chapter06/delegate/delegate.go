package main

import "fmt"

/**
方法和函数的统一调用
*/
// 声明一个结构体
type class struct {
}

// 给结构体添加Do方法
func (c *class) Do(v int) {

	fmt.Println("call method do:", v)
}

// 普通函数的Do
func funcDo(v int) {

	fmt.Println("call function do:", v)
}

func main() {

	// 声明一个函数回调
	//声明一个 delegate 的变量，类型为 func(int)，与 funcDo 和 class 的 Do() 方法的参数一致。
	var delegate func(int)

	// 创建结构体实例
	c := new(class)

	// 将回调设为c的Do方法,将 c.Do 作为值赋给 delegate 变量。
	delegate = c.Do

	// 调用
	delegate(100)

	// 将回调设为普通函数, //签名一致的函数变量就可以保存普通函数或是结构体方法
	delegate = funcDo

	// 调用
	delegate(100)
}
