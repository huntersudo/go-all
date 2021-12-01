package main

var done1  =make(chan bool)
var msg1 string

func aGroutine1()  {
	msg1="你好"
	<-done1   // 阻塞
}

func main() {
	go aGroutine1()
	// 因为 main 线程中 done <- true 发送完成前，后台线程 <- done 接收已经开始，这保证 msg = "hello, world" 被执行了
	done1 <- true
	println(msg1)
}


