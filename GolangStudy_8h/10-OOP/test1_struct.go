package main

import "fmt"

// todo 声明一种行的数据类型 myint， 是int的一个别名
type myint int

// Book1 todo 定义一个结构体
type Book1 struct {
	title string
	auth  string
}

func changeBook(book Book1) {
	//传递一个book的副本
	book.auth = "666"
}

func changeBook2(book *Book1) {
	//指针传递
	book.auth = "777"
}

func main() {

	var a myint = 10
	fmt.Println("a = ", a)
	fmt.Printf("type of a = %T\n", a)  //todo  type of a = main.myint

	var book1 Book1
	book1.title = "Golang"
	book1.auth = "zhang3"

	fmt.Printf("%v\n", book1) // {Golang zhang3}

	changeBook(book1)  // 值传递

	fmt.Printf("%v\n", book1) //{Golang zhang3}

	changeBook2(&book1)  // 引用传递

	fmt.Printf("%v\n", book1) // {Golang 777}
}
