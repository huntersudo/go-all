package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex

	go func() {

		fmt.Println("你好")
		mu.Lock()
	}()
	mu.Unlock()
}
// 这两个事件不可排序也就是可以并发的。因为可能是并 发的事件，所以 main 函数中的 mu.Unlock() 很有可能先发生，
// 而这个时刻 mu 互斥对象还处于 未加锁的状态，从而会导致运行时异常

//修复
// 修复的方式是在 main 函数所在线程中执行两次 mu.Lock() ，当第二次加锁时会因为锁已经被占用 （不是递归锁）而阻塞， main 函数的阻塞状态驱动后台线程继续向前执行
// 当后台线程执行 到 mu.Unlock() 时解锁，此时打印工作已经完成了，解锁会导致 main 函数中的第二 个 mu.Lock() 阻塞状态取消，
// 此时后台线程和主线程再没有其它的同步事件参考，它们退出的事件 将是并发的：在 main 函数退出导致程序退出时，后台线程可能已经退出了，也可能没有退出。虽然 无法确定两个线程退出的时间，但是打印工作是可以正确完成的。
func main1() {
	var mu sync.Mutex

	mu.Lock()
	go func() {

		fmt.Println("你好")
		mu.Unlock()
	}()
	mu.Lock()
}