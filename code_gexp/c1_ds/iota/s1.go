package main

import "fmt"

type Priority int

//iota常用于const表达式中，我们还知道其值是从零开始，const声明块中每增加一行iota值自增1
const (
	// iota初始值为0，也即LOG_EMERG值为0，下面每个常量递增1。
	LOG_EMERG Priority = iota
	LOG_ALERT
	LOG_CRIT
	LOG_ERR
	LOG_WARNING
	LOG_NOTICE
	LOG_INFO
	LOG_DEBUG
)
const (
	// 1
	mutexLocked           = 1 << iota // mutex is locked
	// 1<<1 : 2
	mutexWoken
	// 1<<2: 4
	 mutexStarving                     // 4
	 // 行索引就是  3
	mutexWaiterShift      = iota      // 3
	starvationThresholdNs = 1e6       // 1000000
)

//todo  iota代表了const声明块的行索引（下标从0开始）
const (
	// 1  0
	bit0, mask0 = 1 << iota, 1<<iota - 1  // const声明第0行，即iota==0
	// 2  1
	bit1, mask1  // const声明第1行，即iota==1, 表达式继承上面的语句
	_, _          // const声明第2行，即iota==2 : 1 << 1, 1<<1 - 1
	// 8  7
	bit3, mask3   // const声明第3行，即iota==3: 1 << 3, 1<<3 - 1
)

func main() {
	//// 1 -> 100
	//fmt.Println(1<<2)
	//// 1 -> 1000
	//fmt.Println(1<<3)
	//// 1 -> 10000
	//fmt.Println(1<<4)
	//// 10 -> 100000 :32
	//fmt.Println(2<<4)
	//// 11 ->110000 :32+16
	//fmt.Println(3<<4)
	// 0
	fmt.Println(" ::", mutexWoken)
	// 2
	fmt.Println(" ::", LOG_CRIT)

	//1
	fmt.Println(" ::", mutexLocked)
	// 4
	fmt.Println(" ::", mutexStarving)
	// 3
	fmt.Println(" ::", mutexWaiterShift)

	// :: 1e+06
	// 1000000
	fmt.Println(" ::", starvationThresholdNs)
}
