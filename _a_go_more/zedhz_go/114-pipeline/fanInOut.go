package main

import (
	"fmt"
	"math"
	"sync"
)

// todo 动用 Go 语言的 Go Routine 和 Channel 还有一个好处，就是可以写出 1 对多，或多对 1 的 Pipeline，
// 也就是 Fan In/ Fan Out。下面，我们来看一个 Fan in 的示例。
// 假设我们要通过并发的方式对一个很长的数组中的质数进行求和运算，我们想先把数组分段求和，然后再把它们集中起来。
func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func main() {
	nums := makeRange(1, 10000)
	in := echo(nums)

	//const nProcess = 1 // 5736396
	const nProcess = 5 // 5736396
	var chans [nProcess]<-chan int
	for i := range chans {
		chans[i] = sum(prime(in))
	}

	for n := range sum(merge(chans[:])) {
		fmt.Println(n)  // 5736396
	}
}
// 我来简单解释下这段代码。首先，我们制造了从 1 到 10000 的数组；
// 然后，把这堆数组全部 echo到一个 Channel 里—— in；
// todo 此时，生成 5 个 Channel，接着都调用 sum(prime(in)) ，于是，每个 Sum 的 Go Routine 都会开始计算和；
// todo 这里怎么分组的： in 输入是一个channel, 并发安全的可以被5个go channel 访问
//  每次调用prime，都会启动一个goroutine，然后nProcess个goroutine都会从channel中“抢”数字，
//  但每次只有一个goroutine能成功抢到，至于是哪个goroutine，那就是随机看运气，这样所有数字，基本平均的分成了nProcess组。
//最后，再把所有的结果再求和拼起来，得到最终的结果。
//todo  merge就是需要待所有的channel都处理完成了。但是计算的事都并⾏完了，所以， 在计算上是并发的，在merge上并不是。但真正耗时的是计算⽽不是merge，所以， 并发是有⽤的。
func is_prime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value) / 2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

func prime(in <-chan int) <-chan int {
	out := make(chan int)
	go func ()  {
		for n := range in {
			if is_prime(n) {
				out <- n
			}
		}
		close(out)
	}()
	return out
}


func merge(cs []<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)
    fmt.Println("cs  len:",len(cs))
	wg.Add(len(cs))
	for _, c := range cs {
		go func(c <-chan int) {
			for n := range c {
				out <- n
			}
			wg.Done()
		}(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}