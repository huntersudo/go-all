package main

/**
单向channel
顾名思义，单向channel指只能用于发送或接收数据，实际上也没有单向channel。 我们知道channel可以通过参数传递，
所谓单向channel只是对channel的一种使用限制，这跟C语言使用const修 饰函数参数为只读是一个道理。
func readChan(chanName <-chan int)： 通过形参限定函数内部只能从channel中读取数据
func writeChan(chanName chan<- int)： 通过形参限定函数内部只能向channel中写入数据

 */

func readChan(chanName <-chan int) {
	<- chanName
}

func writeChan(chanName chan<- int) {
	chanName <- 1
}
func main() {
	/**
	mychan是个正常的channel，而readChan()参数限制了传入的channel只能用来读，writeChan()参数限制了传 入的channel只能用来写
	 */
	var mychan = make(chan int, 10)
	writeChan(mychan)
	readChan(mychan)
}
