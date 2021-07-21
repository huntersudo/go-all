package main

import (
	"sync"
	"sync/atomic"
)

//互斥锁的代价比普通整数的原子读写高很多，在性 能敏感的地方可以增加一个数字型的标志位，通过原子检测标志位状态降低互斥锁的使用次数来提高性能。

type singleton struct {}

var (
	instance *singleton
	initialized uint32
	mu sync.Mutex
)

func Instance() *singleton  {
	if atomic.LoadUint32(&initialized)==1{
		return instance
	}
	mu.Lock()
	defer mu.Unlock()

	if instance == nil {
		defer atomic.StoreUint32(&initialized,1) // defer
		instance =&singleton{}
	}
	return instance
}
