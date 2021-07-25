
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 安全退出

// 
 // 标准库增加了一个 context 包，用来简化对于处理单个请求的多个Goroutine之间与请求域的数据、超时和退出等操作，
 // 以用 context 包来重 新实现前面的线程安全退出或超时的控制:

func work1(ctx context.Context,wg *sync.WaitGroup){
	defer wg.Done()
	for {
		select {
		default:
			fmt.Println("hello")
		case <-ctx.Done():
			return
		}
	}
}
func main() {
	//
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//cancel :=make(chan bool)   // todo context replace chan
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go work1(ctx,&wg)
	}

	time.Sleep(time.Second)
	cancel() // todo 当并发体超时或 main 主动停止工作者Goroutine时，每个工作者都可以安全退出。
	wg.Wait()
}