package main

import (
	"context"
	"fmt"
	"time"
)

/*
 Context是个接口
var (
	background = new(emptyCtx)
	todo       = new(emptyCtx)
)
 */
func main() {
	// Background，TODO() 会返回最简单的实现
	ctx := context.Background()
	ctx = context.TODO()

	before := time.Now()
	// withTimeout 指定超时时间
	preCtx, _ := context.WithTimeout(ctx, 500*time.Millisecond)

	go func() {
		// todo context的退出传播关系： 父context的退出会导致子context退出，反之不会
		childCtx, _ := context.WithTimeout(preCtx, 300*time.Millisecond)
		select {
		case <-childCtx.Done():
			after := time.Now()
			fmt.Println("child during:", after.Sub(before).Milliseconds())
		}
	}()

	select {
	case <-preCtx.Done():
		after := time.Now()
		fmt.Println("pre during:", after.Sub(before).Milliseconds())
	}

	time.Sleep(5 * time.Second)
}
