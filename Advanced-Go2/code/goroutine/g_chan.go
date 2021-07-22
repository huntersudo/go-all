package main

var done  =make(chan bool)
var msg string

func aGroutine()  {
	msg="你好"
	done <- true
	// 若在关闭Channel后继续从中接收数据，接收者就会收到该Channel返回的零值
	//close(done)   效果一样

}

func main() {
	go aGroutine()
	<-done
	println(msg)
}


