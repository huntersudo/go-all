// 自定义类型实现多个接口

package main

import (
	"fmt"
)

// A接口
type A2 interface {
	Say()
}

// B接口
type B2 interface {
	Play()
}

type C2 struct {
}

// C2 同时实现2个接口
func (c C2) Say() {
	fmt.Println("C2 Say() ....")
}

func (c C2) Play() {
	fmt.Println("C2 Play() ....")
}

func main() {

	var c C2

	var a A2 = c
	var b B2 = c

	a.Say()
	b.Play()
}
