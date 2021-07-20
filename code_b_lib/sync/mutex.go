// sync包中的互斥锁 Mutex
// 开启两个线程模拟窗口1和窗口2对10张票进行售卖
package main

import (
	"fmt"
	"sync"
	"time"
)

// 定义全局变量
var tickets = 10
var w sync.WaitGroup
var m sync.Mutex

func main() {

	w.Add(2)
	go sellTicket("窗口1")
	go sellTicket("窗口2")
	w.Wait()
}

func sellTicket(name string) {
	defer w.Done()
	for {
		m.Lock() // 给线程加锁
		if tickets > 0 {
			fmt.Println(name, "售出第", tickets, "张票")
			time.Sleep(1) // for test
			tickets--
		} else {
			m.Unlock() // 如果票售完，释放锁，以免造成死锁
			break
		}
		m.Unlock() // 线程执行完，释放锁
	}
	fmt.Println(name, "票全部售完")
}
// 窗口1 售出第 10 张票
//窗口1 售出第 9 张票
//窗口2 售出第 8 张票
//窗口1 售出第 7 张票
//窗口2 售出第 6 张票
//窗口1 售出第 5 张票
//窗口2 售出第 4 张票
//窗口1 售出第 3 张票
//窗口2 售出第 2 张票
//窗口1 售出第 1 张票
//窗口2 票全部售完
//窗口1 票全部售完