package main

import (
	"fmt"
	"time"
)

//通道作为一等公民，作为参数和返回值，
// todo 通道是引用类型，实际都是同一个通道
//func worker(id int, c chan int) {
func worker(id int, c <-chan int) {  // todo 签名换成  单方向通道（只读）
	for n := range c {
		fmt.Printf("Worker %d received %c\n",
			id, n)
	}
}

func createWorker(id int) chan int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func chanDemo() {
	var channels [10]chan int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Millisecond)
}


func main(){

	chanDemo()
}

