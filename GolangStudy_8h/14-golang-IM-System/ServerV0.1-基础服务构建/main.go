package main

func main() {
	server := NewServer("127.0.0.1", 8888)
	server.Start()
}

 //测试
 //nc  127.0.0.1 8888