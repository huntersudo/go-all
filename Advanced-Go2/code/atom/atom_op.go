package main

import (
	"fmt"
	"sync"
)

// 模拟原子操作

var total struct {
	sync.Mutex
	value int
}

func worker2(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 100; i++ {
		total.Lock()
		total.value += i
		total.Unlock()
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go worker2(&wg)
	go worker2(&wg) // 10100
	//go worker2(&wg) // 15150

	wg.Wait()

	fmt.Println(total.value)

}
