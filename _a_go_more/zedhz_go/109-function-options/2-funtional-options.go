package main

import (
	"crypto/tls"
	"fmt"
	"time"
)

// 定义一个函数类型
type Option func(*Server)

//然后，我们可以使用函数式的方式定义一组如下的函数：

func Protocol(p string) Option {
	return func(s *Server) {
		s.Protocol = p
	}
}
func Timeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.Timeout = timeout
	}
}
func MaxConns(maxconns int) Option {
	return func(s *Server) {
		s.MaxConns = maxconns
	}
}
func TLS(tls *tls.Config) Option {
	return func(s *Server) {
		s.TLS = tls
	}
}
// 当我们调用其中的一个函数 MaxConns(30) 时，其返回值是一个 func(s* Server) { s.MaxConns = 30 } 的函数。
// 这个叫高阶函数
// todo 好了，现在我们再定一个 NewServer()的函数，其中，有一个可变参数 options  ，它可以传出多个上面的函数，
//   然后使用一个 for-loop 来设置我们的 Server 对象。


func NewServer(addr string, port int, options ...func(*Server)) (*Server, error) {
	srv := Server{
		Addr:     addr,
		Port:     port,
		Protocol: "tcp",
		Timeout:  30 * time.Second,
		MaxConns: 1000,
		TLS:      nil,
	}
	for _, option := range options {
		option(&srv)
	}
	//...
	return &srv, nil
}

func main()  {

	s1, _ := NewServer("localhost", 1024)
	s2, _ := NewServer("localhost", 2048, Protocol("udp"))
	s3, _ := NewServer("0.0.0.0", 8080, Timeout(300*time.Second), MaxConns(1000))
  fmt.Println(s1)
  fmt.Println(s2)
  fmt.Println(s3)

}

