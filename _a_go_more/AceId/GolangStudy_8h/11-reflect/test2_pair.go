package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//todo： tty: pair<type:*os.File, value:"/dev/tty"文件描述符>  ，不管tty赋值给谁，这个pair保持不变的 ！！！
	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)

	if err != nil {
		fmt.Println("open file error", err)
		return
	}

	//todo： r: pair<type:  , value:>
	var r io.Reader
	//todo： r: pair<type:*os.File, value:"/dev/tty"文件描述符>
	r = tty

	//todo w: pair<type:  , value:>
	var w io.Writer
	//todo  w: pair<type:*os.File, value:"/dev/tty"文件描述符>   这里的值是不变的
	w = r.(io.Writer)

	w.Write([]byte("HELLO THIS is A TEST!!!\n"))  // 调用的还是 os.File的的write方法
}
