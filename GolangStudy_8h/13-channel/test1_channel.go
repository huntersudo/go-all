package main

import "fmt"

func main() {
	//定义一个channel
	c := make(chan int)

	go func() {
		defer fmt.Println("goroutine结束") // todo 在 num:= <-c 之后执行，

		fmt.Println("goroutine 正在运行...")

		c <- 666 //将666 发送给c
	}()

	num := <-c //从c中接受数据，并赋值给num 这个执行完，  todo c<- 666 才会进行下一步
	// todo chan 有同步的能力 ，拿不到值则会阻塞

	fmt.Println("num = ", num)
	fmt.Println("main goroutine 结束...")
}
