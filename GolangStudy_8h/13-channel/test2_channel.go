package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3) //带有缓冲的channel

	fmt.Println("len(c) = ", len(c), ", cap(c)", cap(c))  // len(c) =  0 , cap(c) 3

	go func() {
		defer fmt.Println("子go程结束")

		for i := 0; i < 4; i++ {
			c <- i
			fmt.Println("子go程正在运行, 发送的元素=", i, " len(c)=", len(c), ", cap(c)=", cap(c))
		}
		// 子go程正在运行, 发送的元素= 0  len(c)= 1 , cap(c)= 3
		//子go程正在运行, 发送的元素= 1  len(c)= 2 , cap(c)= 3
		//子go程正在运行, 发送的元素= 2  len(c)= 3 , cap(c)= 3
	}()

	time.Sleep(2 * time.Second)

	//for i := 0; i < 3; i++ {   上面会阻塞
	for i := 0; i < 4; i++ {
		num := <-c //从c中接收数据，并赋值给num
		fmt.Println("num = ", num)
	}

	fmt.Println("main 结束")
}
