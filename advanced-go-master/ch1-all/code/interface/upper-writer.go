package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

/*
"Fprintf" function signature:
func Fprintln(w io.Writer, a ...interface{}) (n int, err error)

"io.Writer" is an interface:
type io.Write interface {
	Write(p []byte) (n int, err error)
}
*/

type upperWriter struct {
	io.Writer
}

func (p *upperWriter) Write(data []byte) (n int, err error) {
	return p.Writer.Write(bytes.ToUpper(data))
}

// fmt.Fprintf 却可以向任何自定义的输出流对象打印，可以打印到文件或标准输出、也可以 打印到网络、甚至可以打印到一个压缩文件
func main() {
	fmt.Println("Upper Writer")
	fmt.Fprintln(&upperWriter{os.Stdout}, "hello, world")
}
//Upper Writer
//HELLO, WORLD

