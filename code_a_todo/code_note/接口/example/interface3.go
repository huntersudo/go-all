// 接口可以继承多个接口
// 注意：如果C接口继承A和B接口，如果要实现C接口，必须也要实现A和B接口中的方法

package main

import (
	"fmt"
)

type A3 interface {
	Say()
}

type B3 interface {
	Play()
}

type C3 interface {
	A3 // 继承A接口
	B3 // 继承B接口
	Eat()
}

type D struct {
}

// 下面是错误的，因为实现C接口，没有实现A和B接口中的方法
// func (d D) Eat() {
//	fmt.Println("Eat() ....")
//}

//func server() {

//	var d D
//	var c C = d

//	c.Eat()
//}

// 下面的是正确的
func (d D) Eat() {
	fmt.Println("D Eat() ....")
}

func (d D) Say() {
	fmt.Println("D Say() ....")
}

func (d D) Play() {
	fmt.Println(" D Play() ....")
}

func main() {

	var d D
	var c C3 = d

	c.Eat()
	c.Say()
	c.Play()

}
