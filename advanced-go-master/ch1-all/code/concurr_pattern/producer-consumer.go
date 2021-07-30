package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// 生产者: 生成 factor 整数倍的序列  todo: X chan<- type , send to channel
func producer(factor int, out chan<- int) {
	for i := 0; ; i++ {
		time.Sleep(500 * time.Millisecond)
		out <- i * factor
	}
}

// 消费者 todo  X <-chan type   ,receive from  channel
func consumer(in <-chan int) {
	for v := range in {
		time.Sleep(250 * time.Millisecond)
		fmt.Println(v)
	}
}

func main() {
	fmt.Println("Produce and Consumer Pattern")

	ch := make(chan int, 64)

	go producer(3, ch) //  生成 3 的倍数的序列
	go producer(5, ch) // 生成 5 的倍数的序列
	go consumer(ch)   // 消费 生成的队列

	//// Ctrl+C 退出
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT)
	fmt.Printf("quit (%v)\n", <-sig)
}