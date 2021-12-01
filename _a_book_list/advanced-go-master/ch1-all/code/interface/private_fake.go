package main

import (
	"fmt"
	"testing"
)


//嵌入匿名 的 testing.TB 接口来伪造私有的 private 方法，因为接口方法是延迟绑定，编译 时 private 方法是否真的存在并不重要。
type TB struct {
	testing.TB
}

func (p *TB)Fatal(args ...interface{})  {
	fmt.Println("TB.Fatal disabled!")
}

func main() {
	var tb testing.TB =new (TB)
	tb.Fatal("Hello, playground")
}
// 在自己的 TB 结构体类型中重新实现了 Fatal 方法，然后通过将对象隐式转换 为 testing.TB 接口类型（因为内嵌了匿名的 testing.TB 对象，因此是满足 testing.TB 接口 的），
// 然后通过 testing.TB 接口来调用我们自己的 Fatal 方法。
// todo 这种通过嵌入匿名接口或嵌入匿名指针对象来实现继承的做法其实是一种纯虚继承，我们继承的只是接 口指定的规范，真正的实现在运行的时候才被注入