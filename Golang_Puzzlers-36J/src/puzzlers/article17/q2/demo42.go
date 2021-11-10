package main

import "fmt"

func main() {
	//示例1。
	value1 := [...]int8{0, 1, 2, 3, 4, 5, 6}
	switch 1 + 3 {   // int
	case value1[0], value1[1]:  // 这条语句无法编译通过。  int8 类型不同，无法编译
		fmt.Println("0 or 1")
	case value1[2], value1[3]:  // 这条语句无法编译通过。
		fmt.Println("2 or 3")
	case value1[4], value1[5], value1[6]:  // 这条语句无法编译通过。
		fmt.Println("4 or 5 or 6")
	}

	// 示例2。
	value2 := [...]int8{0, 1, 2, 3, 4, 5, 6}
	switch value2[4] {   // int8
		// 如果case表达式中子表达式的结果值是无类型的常量，那么它的类型会被自动地转换为switch表达式的结果类型
	case 0, 1:   // 无类型
		fmt.Println("0 or 1")
	case 2, 3:
		fmt.Println("2 or 3")
	case 4, 5, 6:
		fmt.Println("4 or 5 or 6")
	}
	// 4 or 5 or 6
}
