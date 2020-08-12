// 继承使用细节之二
// 匿名结构体字段简化写法

package main

import (
	_ "fmt"
)

type Stu2 struct {
	Name string
}

type Puple2 struct {
	Stu2
}

func main() {

	var p Puple2

	// 可简化成 p.Name = "张三"
	// 如果Puple结构体内没有字段Name，就看查看此结构体内是否有匿名结构体，
	// 如果有则查看匿名结构体内是否含有字段Name，有则使用，否则报错
	p.Stu2.Name = "张三"
}
