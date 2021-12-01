package main

import (
	"sync"
)

//互斥锁的代价比普通整数的原子读写高很多，在性 能敏感的地方可以增加一个数字型的标志位，通过原子检测标志位状态降低互斥锁的使用次数来提高性能。

type singleton1 struct {}

var (
	instance1 *singleton1
	once sync.Once
)
//todo  基于 sync.Once 重新实现单件模式：
func Instance1() *singleton {
	once.Do(func() {
		instance1 =&singleton1{}
	})

	return instance
}
