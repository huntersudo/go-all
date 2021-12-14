package main

import "fmt"

func main(){
	a := 1
	// 参数预计算 P126
	defer func(b int) {
		fmt.Println("defer b",b)
	}(a+1)

	a = 99
}
// defer b 2
