package main

import "fmt"

//todo  接口完整性检查

// 另外，我们可以看到，Go 语言的编译器并没有严格检查一个对象是否实现了某接口所有的接口方法，如下面这个示例：

type Shape interface {
	Sides() int
	Area() int
}
type Square struct {
	len int
}
func (s* Square) Sides() int {
	return 4
}
func main() {
	s := Square{len: 5}
	fmt.Printf("%d\n",s.Sides())
    /**
	Cannot use '(*Square)(nil)' (type *Square) as the type Shape Type
	does not implement 'Shape' as some methods are missing: Area() int
     */
	//var _ Shape = (*Square)(nil)
}
// todo 可以看到，Square 并没有实现 Shape 接口的所有方法，程序虽然可以跑通，但是这样的编程方式并不严谨，如果我们需要强制实现接口的所有方法，那该怎么办呢
//  在 Go 语言编程圈里，有一个比较标准的做法：
//   声明一个 _ 变量（没人用）会把一个 nil 的空指针从 Square 转成 Shape，这样，如果没有实现完相关的接口方法，编译器就会报错
//   这样就做到了强验证的方法。
