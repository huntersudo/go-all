package main

import (
	"fmt"
)

// 1.6.6 素数筛
//  最初的 2, 3, 4, ... 自然数序列（不包含开头的0、1）

func generateNatural2() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}




func primeFilter2(in <-chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i:=<-in;i%prime !=0{
				out <- i
			}
		}
	}()

	return out
}


func main() {
	fmt.Println("Prime Filter")
	

	ch := generateNatural2() // todo 如何存储自然数序列的？ 无限?
	// Find out 100 Prime Number 筛子，筛100次，
	for i := 0; i < 100; i++ {
		prime := <-ch
		fmt.Printf("No.%v: %v\n", i+1, prime)
		ch = primeFilter2( ch, prime)
	}

}
