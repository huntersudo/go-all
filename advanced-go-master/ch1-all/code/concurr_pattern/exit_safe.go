
package main

import (
	"fmt"
	"sync"
	"time"
)

// 安全退出

//
 // 因为 main 线程并没有等待各个工作Goroutine退出工作完成的机制。我们可以结 合 sync.WaitGroup 来改进:
// 现在每个工作者并发体的创建、运行、暂停和退出都是在 main 函数的安全控制之下了
// 通过 close 来关闭 cancel 管道向多个Goroutine广播退出的指令
func work(wg *sync.WaitGroup, canncel chan bool){
	defer wg.Done()
	for {
		select {
		default:
			fmt.Println("hello")
		case <-canncel:
			return
		}
	}
}
func main() {

	cancel :=make(chan bool)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go work(&wg,cancel)
	}

	time.Sleep(time.Second)
	close(cancel)
	wg.Wait()
}