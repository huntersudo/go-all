package main

func main() {
	// 示例1。
	ch1 := make(chan int, 1)
	ch1 <- 1
	ch1 <- 2 // 通道已满，因此这里会造成阻塞。
/**
  fatal error: all goroutines are asleep - deadlock!
  goroutine 1 [chan send]:
  main.main()
 */

	// 示例2。
	ch2 := make(chan int, 1)
	//elem, ok := <-ch2 // 通道已空，因此这里会造成阻塞。
	//_, _ = elem, ok
	ch2 <- 1

	// 示例3。
	var ch3 chan int
	//ch3 <- 1 // 通道的值为nil，因此这里会造成永久的阻塞！
	//<-ch3 // 通道的值为nil，因此这里会造成永久的阻塞！
	_ = ch3
}


/**
通道底层存储数据的是链表还是数组？
作者回复: 环形链表

作者回复: 再说一遍，Go语言里没有深层复制。数组是值类型，所以会被完全复制。

 */