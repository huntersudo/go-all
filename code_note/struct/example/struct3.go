// 结构体 方法

package main

import (
	"fmt"
)

type User1 struct {
	Name string
	Age  int
	Sex  string
}

func (u User1) say() {
	fmt.Println(u.Name, u.Age, u.Sex) //  张三 12 男
}

func main() {

	u := User1{"张三", 12, "男"}
	//调用方法
	u.say()
}
