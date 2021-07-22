package main

var a string

func f()  {
	println(a)
}
func hello()  {
	 a="hello"
	 go f()
	 println("return")
}
// 执行 go f() 语句创建Goroutine和 hello 函数是在同一个Goroutine中执行, 根据语句的书写 顺序可以确定Goroutine的创建发生在 hello 函数返回之前
//todo  但是新创建Goroutine对应 的 f() 的执行事件和 hello 函数返回的事件则是不可排序的，也就是并发的
// 调用 hello 可能 会在将来的某一时刻打印 "hello, world" ，也很可能是在 hello 函数执行完成后才打印。

func main() {
	hello()
}
// 1、只打印
// return

// 2、打印
// return
// hello