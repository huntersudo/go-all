// 结构体的嵌套之一

package main

import (
	"fmt"
)

type Book1 struct {
	BookName  string
	BookPrice int
}

type Stu1 struct {
	StuName    string
	StuAge     int
	StuFavBook Book1
}

func main() {

	s1 := Stu1{
		StuName: "pengw",
		StuAge:  20,
		StuFavBook: Book1{
			BookName:  "Go语言",
			BookPrice: 20,
		},
	}

	fmt.Println(s1, s1.StuFavBook.BookName) // {pengw 20 {Go语言 20}} Go语言
}
