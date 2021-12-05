package main

import (
	"fmt"
	"time"
)
// todo
//  chan  读写
//  chan <- T 只写T类型
//  <- chan T 只读T类型
func main(){
	// 通道初始化 make
	var c  = make(chan int)

	go func() {
		data,ok:= <-c
		fmt.Println("goroutine one: ",data,ok)
	}()
	go func() {
		data,ok:= <-c
		fmt.Println("goroutine two: ",data,ok)
	}()
	close(c) //todo 关闭通道，==向所有正在读取的协程中都写入了数据
	time.Sleep(1*time.Second)
   //goroutine one:  0 false
	//goroutine two:  0 false
}
