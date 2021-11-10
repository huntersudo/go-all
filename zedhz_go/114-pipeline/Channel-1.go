package main

import "fmt"

// channel 转发函数
//Rob Pike 在 Go Concurrency Patterns: Pipelines and cancellation 这篇博客中介绍了一种编程模式，下面我们来学习下。
// https://blog.golang.org/pipelines

// 我们需要一个 echo()函数，它会把一个整型数组放到一个 Channel 中，并返回这个 Channel。
func echo(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}


// 平方函数
func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

//过滤奇数函数
func odd(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			if n%2 != 0 {
				out <- n
			}
		}
		close(out)
	}()
	return out
}

//求和函数
func sum(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		var sum = 0
		for n := range in {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}

func main1() {

	var nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for n := range sum(sq(odd(echo(nums)))) {
		fmt.Println(n)  // 165
	}

	var nums1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for n := range pipeline(nums1, echo, odd, sq, sum){
		fmt.Println(n)  // 165
	}

}
// 上⾯的代码类似于我们执⾏了Unix/Linux命令： echo $nums | sq | sum
// todo 同样，如果你不想有那么多的函数嵌套，你可以使⽤⼀个代理函数来完成。


type EchoFunc func ([]int) (<- chan int)
type PipeFunc func (<- chan int) (<- chan int)

func pipeline(nums []int, echo EchoFunc, pipeFns ... PipeFunc) <- chan int {
	ch  := echo(nums)
	for i := range pipeFns {
		ch = pipeFns[i](ch)
	}
	return ch
}
