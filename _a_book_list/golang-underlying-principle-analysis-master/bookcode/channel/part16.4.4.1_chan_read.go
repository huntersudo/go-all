package main

import (
	"fmt"
	"time"
)

func main() {

	a := make(chan int,2)

	go func() {
		for i := 0; i < 2; i++ {  // 0,1   2会被阻塞，放到sendq队列上
			a<-i
		}
	}()
	time.Sleep(5*time.Second)
	for i := 0; i < 2; i++ {  // 0,1   2会被阻塞，放到sendq队列上
		fmt.Printf("seq:%v, value:%v\n",i,<-a)
	}

}
