package main

import "fmt"

// 赢者为王   谁最快 取谁
func main(){
	ch :=make(chan string,32)
	go func() {
		ch <-searchByBing("golang")
	}()

	go func() {
		ch <-searchByBaidu("golang")
	}()

	go func() {
		ch <-searchByGoogle("golang")
	}()

	fmt.Println(<-ch)
 // 当任意一个搜索引擎最 先有结果之后，都会马上将结果发到管道中（因为管道带了足够的缓存，这个过程不会阻塞）。
 // 但是最 终我们只从管道取第一个结果，也就是最先返回的结果。
 // 通过适当开启一些冗余的线程，尝试用不同途径去解决同样的问题，最终以赢者为王的方式提升了程序 的相应性能。
}

