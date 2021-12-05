package main

import "sync/atomic"


var flag  int64  = 0
var count  int64  = 0
func add() {
	//TODO P289 原子操作  构建 自旋锁  ---？
	for {
		if atomic.CompareAndSwapInt64(&flag, 0, 1) {
			count++
			atomic.StoreInt64(&flag, 0)
			return
		}
	}
}

func main() {
	go add()
	go add()
}
