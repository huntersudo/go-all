package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 模拟原子操作   sync/atomic
var total2 uint64

func worker3(wg *sync.WaitGroup) {
	defer wg.Done()
	var i uint64
	for i = 0; i <= 100; i++ {
		// todo atomic.AddUint64 函数调用保证了 total 的读取、更新和保存是一个原子操作，因此在多线程 中访问也是安全的。
		atomic.AddUint64(&total2,i)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go worker3(&wg)
	go worker3(&wg) // 10100
	//go worker3(&wg) // 15150

	wg.Wait()

	fmt.Println(total2)

}
