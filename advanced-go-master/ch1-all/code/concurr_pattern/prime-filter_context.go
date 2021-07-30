package main

import (
	"context"
	"fmt"
)

// 1.6.6 素数筛
//  最初的 2, 3, 4, ... 自然数序列（不包含开头的0、1）
func generateNatural(ctx context.Context) chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			select {
			case <- ctx.Done():
				return
			case ch <- i:
			}
		}
	}()
	return ch
}

// 然后是为每个素数构造一个筛子：将输入序列中是素数倍数的数提出，并返回新的序列，是一个新的管 道。
// // 管道过滤器: 删除能被素数整除的数
func primeFilter(ctx context.Context, in <-chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			select {
			case <- ctx.Done():
				return
			case i := <-in:
				if i % prime != 0 {
					out <- i
				}
			}
		}
	}()

	return out
}


func main() {
	fmt.Println("Prime Filter")

	// todo 当 main 函数不再 使用管道时后台Goroutine有泄漏的风险。我们可以通过 context 包来避免
    //  通过 Context 控制后台Goroutine状态
	ctx, cancel := context.WithCancel(context.Background())

	ch := generateNatural(ctx) // // 自然数序列: 2, 3, 4, ...
	// Find out 100 Prime Number 筛子，筛100次，
	for i := 0; i < 100; i++ {
		prime := <-ch
		fmt.Printf("No.%v: %v\n", i+1, prime)
		ch = primeFilter(ctx, ch, prime)
	}

	cancel()  //todo 当main函数完成工作前，通过调用 cancel() 来通知后台Goroutine退出，这样就避免了 Goroutine的泄漏。
}
