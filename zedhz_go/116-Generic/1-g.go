package main

import "fmt"
// 没有泛型的话，这个函数需要写出 int 版，float版，string 版，以及我们的自定义类型（struct）的版本
func print[T any] (arr []T) {
	for _, v := range arr {
		fmt.Print(v)
		fmt.Print(" ")
	}
	fmt.Println("")
}

func main() {
	strs := []string{"Hello", "World",  "Generics"}
	decs := []float64{3.14, 1.14, 1.618, 2.718 }
	nums := []int{2,4,6,8}

	print(strs)
	print(decs)
	print(nums)
}


