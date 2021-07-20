package main

import (
	"fmt"
	//"time"
)

func main() {
	num := 10
	sign := make(chan struct{}, num)

	/**
	Go语言里只有传值，没有传引用。
	如果go函数是无参数的匿名函数，那么在它里面的fmt.Println函数的参数只会在go函数被执行的时候才会求值。
	到那个时候，i的值可能已经是10（最后一个数）了，因为for语句那时候可能已经都执行完毕了。
	 */
	for i := 0; i < num; i++ {
		go func() {
			fmt.Println(i)
			// 空结构体类型
			sign <- struct{}{}
		}()
	}
	/** 结果每次都不一样
	10
	2
	10
	10
	10
	3
	10
	10
	10
	10
	 */

	// 办法1。
	//time.Sleep(time.Millisecond * 500)

	// 办法2。
	for j := 0; j < num; j++ {
		<-sign
	}
	/**
	我们先创建一个通道，它的长度应该与我们手动启用的 goroutine 的数量一致。在每个手动启用的 goroutine 即将运行完毕的时候，我们都要向该通道发送一个值。
	 */
}
