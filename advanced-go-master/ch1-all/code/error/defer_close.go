package main

import (
	"fmt"
	"io"
	"os"
)

//1.7.1 错误处理策略
func copyFile(destFileName, srcFileName string) (written int64, err error) {
	src, err := os.Open(srcFileName)
	if err != nil {
		return
	}
	defer src.Close()

	dest, err := os.Create(destFileName)
	if err != nil {
		return
	}
	defer dest.Close()
   // written, err = io.Copy(dst, src)
	return io.Copy(dest, src)
}

func main() {
	fmt.Println("Defer when Error")

	// copyFile("/tmp/log2.txt", "/tmp/log1.txt")
}
