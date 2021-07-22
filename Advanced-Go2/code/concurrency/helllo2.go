package main

import "fmt"

// 使用 sync.Mutex 互斥锁同步是比较低级的做法。我们现在改用无缓存的管道来实现同步：
func main() {
	done:=make(chan int)
	go func() {
		fmt.Println("你好")
		<- done //
	}()

	done <- 1  // 这里如果想完成，则  <-done 必须开始，

}
// 范，对于从无缓冲Channel进行的接收，发生在对该Channel进行的发送完成 之前
// 。因此，后台线程 <-done 接收操作完成之后， main 线程的 done <- 1 发送操作才可能完 成（从而退出main、退出程序），而此时打印工作已经完成了


//todo 改进：
// 上面的代码虽然可以正确同步，但是对管道的缓存大小太敏感：如果管道有缓存的话，就无法保证main 退出之前后台线程能正常打印了
// 更好的做法是将管道的发送和接收方向调换一下，这样可以避免同步 事件受管道缓存大小的影响：
func main2() {
	done:=make(chan int,1)
	go func() {
		fmt.Println("你好")
		done <- 1
	}()

	<- done   // 想读，但是得等有人写了才可以读
}

// 打印线程扩充

func main4() {
	done := make(chan int, 10) // 带 10 个缓存

	// 开N个后台打印线程

	for i := 0; i < cap(done); i++ {
		go func(){
			fmt.Println("你好, 世界")
			done <- 1
		}()
	}

	// todo 等待N个后台线程完成
	for i := 0; i <cap(done) ; i++ {
		<-done
	}
}

//改进 sync.WaitGroup



