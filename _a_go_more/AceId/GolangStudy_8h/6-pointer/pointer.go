package main

import "fmt"

/*
func swap(a int ,b int) {
	var temp int
	temp = a
	a = b
	b = temp
}
*/

func swap(pa *int, pb *int) {
	var temp int
	temp = *pa //temp = main::a
	*pa = *pb  // main::a = main::b
	*pb = temp // main::b = temp
}
func changeValue(p *int){
	*p=10  // *指针 去找到指针对应的值
	// p=10   // 就会把指针 p 的值 改为 10
}

func main() {
	var a int = 10
	var b int = 20

	swap(&a, &b)

	fmt.Println("a = ", a, " b = ", b)


	var p *int
	p = &a  // 指针类型  =  指针类型
	fmt.Println(&a)
	fmt.Println(p)
	var pp **int //二级指针

	pp = &p  // 二级指针  =  一级指针的地址

	fmt.Println(&p)
	fmt.Println(pp)
}