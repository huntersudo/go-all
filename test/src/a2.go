package main

import (
	"net/http"
)

//conn, err := net.Dial("tcp","127.0.0.1:8080")
//
//var conn net.Conn
//var err error
//conn, err = net.Dial("tcp", "127.0.0.1:8080")

conn, err := net.Dial("tcp", "127.0.0.1:8080")
conn2, err := net.Dial("tcp", "127.0.0.1:8080")


func main() {
	//使用 http.FileServer 文件服务器将当前目录作为根目录（/目录）的处理器，访问根目录，就会进入当前目录。
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":8080", nil)

	for a := 0; a < 10; a++ {
		// 循环代码
	}

}

type IntSlice []int

func (p IntSlice) Len() int {
	return len(p)
}
func (p IntSlice) Less(i, j int) bool {
	return p[i] < p[j]
}
func (p IntSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
