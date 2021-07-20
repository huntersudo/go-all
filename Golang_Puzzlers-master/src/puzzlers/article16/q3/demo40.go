package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var count uint32

	// trigger函数会不断地获取一个名叫count的变量的值，并判断该值是否与参数i的值相同
	trigger := func(i uint32, fn func()) {
		for {
			if n := atomic.LoadUint32(&count); n == i {
				fn()
				atomic.AddUint32(&count, 1)
				break
			}
			time.Sleep(time.Nanosecond)
		}
	}


	for i := uint32(0); i < 10; i++ {
		go func(i uint32) {
			fn := func() {
				fmt.Println(i)
			}
			trigger(i, fn)
		}(i)
		// 在go语句被执行时，我们传给go函数的参数i会先被求值，如此就得到了当次迭代的序号
	}
	// 由于当所有我手动启用的 goroutine 都运行完毕之后，count的值一定会是10，所以我就把10作为了第一个参数值。
	// 又由于我并不想打印这个10，所以我把一个什么都不做的函数作为了第二个参数值。
	trigger(10, func() {})
}
