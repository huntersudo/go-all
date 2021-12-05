package main

import (
	"fmt"
	"time"
)

// TODO case是随机选取的
func main() {
	c := make(chan int)

	tick := time.Tick(time.Second) // 定时执行，每1s向tick通道写入数据

	// todo 很多时候，不希望select执行完一个就退出，配合for，实现循环往复
	//   time.After 新一轮的for+select会重置time.After
	for{
		select {  // 如果没有任何chan准备好，会陷入阻塞
		case <-c:
			fmt.Println("random 01")
		case <-tick:
			fmt.Println("tick")
			// select 与 超时控制
		case <-time.After(800 * time.Millisecond):
			fmt.Println("timeout")
		}
	}
}


