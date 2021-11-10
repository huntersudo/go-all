package main

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

type SumFunc func(int64, int64) int64

//todo 获取函数的名字 代码中使用了 Go 语言的反射机制来获取函数名；
func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func timedSumFunc(f func(int64, int64) int64) func(int64, int64) int64 {
//func timedSumFunc(f SumFunc) SumFunc {
	return func(start, end int64) int64 {

		defer func(t time.Time) {
			fmt.Printf("--- Time Elapsed (%s): %v ---\n",getFunctionName(f), time.Since(t))
		}(time.Now())

		return f(start, end)
	}
}

func Sum1(start, end int64) int64 {
	var sum int64
	sum = 0
	if start > end {
		start, end = end, start
	}
	for i := start; i <= end; i++ {
		sum += i
	}
	return sum
}

func Sum2(start, end int64) int64 {
	if start > end {
		start, end = end, start
	}
	return (end - start + 1) * (end + start) / 2
}

func main() {

	sum1 := timedSumFunc(Sum1)
	sum2 := timedSumFunc(Sum2)

	fmt.Printf("%d, %d\n", sum1(-10000, 10000000), sum2(-10000, 10000000))
}
// --- Time Elapsed (main.Sum1): 3.9912ms ---
//--- Time Elapsed (main.Sum2): 0s ---
//49999954995000, 49999954995000