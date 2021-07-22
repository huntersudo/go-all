package main

import (
	"sync"
	"sync/atomic"
	"time"
)

//互斥锁的代价比普通整数的原子读写高很多，在性 能敏感的地方可以增加一个数字型的标志位，通过原子检测标志位状态降低互斥锁的使用次数来提高性能。

type singleton struct {}

var (
	instance *singleton
	initialized uint32
	mu sync.Mutex
)
// todo 参考sync.Once的源码  类似的
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



// 简化的生产消费模型、后台线程生澄最新的配置信息、前台多个工作者线程获取最新的配置，所有线程共享配置信息资源
func produceConsume(){

	// example 2
	var config atomic.Value  // 保存当前配置信息

	// 初始化配置信息
	config.Store(loadConfig())

	// 启动一个后台线程 ，加载更新后的配置信息

	go func(){
		for{
			time.Sleep(time.Second)
			config.Store(loadConfig())
		}
	}()

	// 用于处理请求的工作线程始终采用最新的配置信息
	for i := 0; i < 10; i++ {
		go func() {
			for r:=range requests(){
				c:=config.Load()
				// ...
			}
		}()
	}



}

