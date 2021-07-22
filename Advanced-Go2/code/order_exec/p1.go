package main

import "sync"

func main() {
	go print("hello")
}
// todo main函数退出时，不会等待任何后台进程，因为goroutine的执行和main函数的返回事件 是并发的


// 明确顺序
func main1() {
	done:=make(chan int)
	go func() {
		print("hello")
		done <- 1
	}()

	<- done
}

// 当 <-done 执行时，必然要求 done <- 1 也已经执行
// 根据同一个Gorouine依然满足顺序一致性 规则，我们可以判断当 done <- 1 执行时， println("你好, 世界") 语句必然已经执行完成了
// 这个时候，可以保证正常打印结果

// 方案2 sync.Mutex 互斥量
func main2() {
	var mu sync.Mutex
	mu.Lock()
	go func() {
		print("hello")
		mu.Unlock()
	}()

	mu.Lock()
}
// 确定后台线程的 mu.Unlock() 必然在 println("你好, 世界") 完成后发生（同一个线程满足 顺序一致性）
// main 函数的第二个 mu.Lock() 必然在后台线程的 mu.Unlock() 之后发生

