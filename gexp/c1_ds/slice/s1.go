package main

import (
	"fmt"
)

/**
slice跟据数组array创建，与数组共享存储空间，slice起始位置是array[5]，
长度为1，容量为5，
slice[0]和array[5]地址相同。
 */
func main() {
	var array [10]int
	var slice = array[5:6]
	// 1
	fmt.Println("lenth of slice: ", len(slice))
	// 5
	fmt.Println("capacity of slice: ", cap(slice))
	// true
	fmt.Println(&slice[0] == &array[5])
}
