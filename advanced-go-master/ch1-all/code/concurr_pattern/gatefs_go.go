package main

import (
	"golang.org/x/tools/godoc/vfs"
	"os"

	//"golang.org/x/tools/godoc/vfs/gatefs"
)

var limit =make(chan int,3)

func main1()  {
	for _,w := range work{
		go func(){
	      limit<-1
	      w()
	      <- limit
	   }()
	}

	select {}
}
//并发数控制的原理在前面一节已经讲过，就是通过 带缓存管道的发送和接收规则来实现最大并发阻塞：
// 不过 gatefs 对此做一个抽象类型 gate ，增加了 enter 和 leave 方法分别对应并发代码的 进入和离开。当超出并发数目限制的时候， enter 方法会阻塞直到并发数降下来为止。

type gate chan bool

func (g gate) enter() { g <- true }
func (g gate) leave() { <-g }

type gatefs struct {
	fs vfs.FileSystem
	gate
}

func (fs gatefs) lstat(p string) (os.FileInfo ,error){
	fs.enter()  // 底层chan大小受控
	defer fs.leave()
	return fs.fs.Lstat(p)
}

func main()  {
	// 然后 gatefs.New 基于现有 的虚拟文件系统构造一个并发受控的虚拟文件系统。
	fs := gatefs.New(vfs.OS("/path"), make(chan bool, 8))
}